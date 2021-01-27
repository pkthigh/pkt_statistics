package storage

import (
	"fmt"
	"pkt_statistics/library/config"
	"pkt_statistics/library/logger"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/redis.v5"
)

// Storage 存储模块
type Storage struct {
	dbs map[MYSQL]*gorm.DB
	rds map[AREA]*redis.Client
	mgo map[COLL]*mongo.Collection
}

// NewStorage new storage
func NewStorage() (*Storage, error) {
	storage := &Storage{
		dbs: make(map[MYSQL]*gorm.DB),
		rds: make(map[AREA]*redis.Client),
		mgo: make(map[COLL]*mongo.Collection),
	}

	// 载入Mysql
	for name, addr := range config.GetStorageConf().SQL.DBs {
		db, err := gorm.Open("mysql", addr)
		if err != nil {
			logger.ErrorF("connect %v error: %v", name, err)
		}
		logger.Info(db)

		db.LogMode(true)
		db.Callback().Create().Replace("gorm:update_time_stamp", updateTimestampForCreateCallback)
		db.Callback().Update().Replace("gorm:update_time_stamp", updateTimestampForUpdateCallback)
		db.Callback().Delete().Replace("gorm:delete", updateTimestampForDeleteCallback)

		storage.dbs[MYSQL(name)] = db
	}

	// 载入Redis
	if config.GetStorageConf().Rds.Addr != "" {
		fmt.Printf("redis address: %v\n", config.GetStorageConf().Rds.Addr)
		for i := 1; i <= 15; i++ {
			cli := redis.NewClient(&redis.Options{
				Addr:     config.GetStorageConf().Rds.Addr,
				Password: config.GetStorageConf().Rds.Password,
				DB:       i,
			})
			fmt.Println("redis db: ", i, cli.Ping().String())
			storage.rds[AREA(i)] = cli
		}
	}

	/*
		// 载入Mongo
		ctx := context.Background()
		cli, err := mongo.Connect(ctx, options.Client().ApplyURI(config.GetStorageConf().Mgo.URI))
		if err != nil {
			return nil, err
		}
		if mgo := cli.Database(config.GetStorageConf().Mgo.DataBase); mgo != nil {
			for _, coll := range colls {
				storage.mgo[coll] = mgo.Collection(coll.String())
			}
		} else {
			return nil, fmt.Errorf("connection mongo database fail")
		}
	*/

	return storage, nil
}

// DBs mysql dbs client
func (storage *Storage) DBs(mysql MYSQL) *gorm.DB {
	return storage.dbs[mysql]
}

// Rds redis area client
func (storage *Storage) Rds(area AREA) *redis.Client {
	return storage.rds[area]
}

// Mgo mongo coll client
func (storage *Storage) Mgo(coll COLL) *mongo.Collection {
	return storage.mgo[coll]
}

// --- ---

// updateTimestampForCreateCallback will set `CreatedAt`, `UpdatedAt` when creating
func updateTimestampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedAt"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("UpdatedAt"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

// updateTimestampForCreateCallback will set `UpdatedAt` when updating
func updateTimestampForUpdateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if modifyTimeField, ok := scope.FieldByName("UpdatedAt"); ok {
			modifyTimeField.Set(nowTime)
		}
	}
}

// updateTimestampForCreateCallback will set `DeletedAt` when deleting
func updateTimestampForDeleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedAt")

		if !scope.Search.Unscoped && hasDeletedOnField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(time.Now().Unix()),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
