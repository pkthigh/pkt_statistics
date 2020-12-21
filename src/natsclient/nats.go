package natsclient

import (
	"os"
	"os/signal"
	"pkt_statistics/common"
	"pkt_statistics/library/config"
	"pkt_statistics/library/logger"
	"syscall"
	"time"

	stan "github.com/nats-io/stan.go"
)

var client *Nats

// Nats nats-io client
type Nats struct {
	date string
	conn stan.Conn
	msgs chan *stan.Msg
}

func init() {
	conf := config.GetNatsConf()
	conn, err := stan.Connect(
		conf.Cluster,
		conf.Client,
		stan.NatsURL(conf.URLs),
		stan.ConnectWait(15*time.Second),
		stan.Pings(3, 5),
		stan.SetConnectionLostHandler(
			func(c stan.Conn, err error) {
				logger.FatalF("nats connect interrupt: %v", err)
			},
		),
	)
	if err != nil {
		logger.FatalF("nats connect error: %v", err)
	}
	client = &Nats{date: time.Now().Format("2006-01-02"), conn: conn, msgs: make(chan *stan.Msg, 1024)}

	// 订阅消息
	client.Sub(common.ItemSubject)         // 道具
	client.Sub(common.PkcHandOverSubject)  // 德州
	client.Sub(common.FishHandOverSubject) // 钓鱼

	logger.Info("nats init successful")
}

// Sub 订阅
func (nats *Nats) Sub(subject common.Subject) {
	_, err := nats.conn.Subscribe(subject.String(), func(msg *stan.Msg) {
		nats.msgs <- msg
	})
	if err != nil {
		logger.ErrorF("nats sub subject %v fail: %v", subject, err)
	}
}

// Close 连接关闭
func (nats *Nats) Close() error {
	if err := nats.conn.Close(); err != nil {
		logger.ErrorF("nats conn close error: %v", err)
		return err
	}
	return nil
}

// Run 运行
func Run() {

	// 处理业务数据
	go client.register()

	// 侦听关闭信号
	signals := make(chan os.Signal)
	signal.Notify(signals, os.Interrupt, os.Kill, syscall.SIGUSR1, syscall.SIGUSR2)
	<-signals

	if err := client.Close(); err != nil {
		logger.ErrorF("close client subs error: %v", err)
	}
}

/*
// ReloadLastCloseDateData 重新载入上次服务关闭日至今的数据
func ReloadLastCloseDateData(val string) {
	var store *redis.Client

	logger.InfoF("get last close time: %v", val)

	lastime, err := time.ParseInLocation("2006-01-02", val, time.Local)
	if err != nil {
		logger.FatalF("last close time parse error: %v", err)
	}

	type Stat struct {
		Date string
		UID  int
		Num  int
	}

	// 德州道具
	var itemfrees []Stat
	src.Store.DBs(common.DataDsn).Raw("SELECT FROM_UNIXTIME(`created_at`, '%Y-%m-%d') AS 'date', `user_id` AS `uid`, `num` AS 'num' FROM pk_data.pkc_check_props_flow WHERE(`created_at` >= ? AND `created_at` < ?)", lastime.Unix(), time.Now().Unix()).Scan(&itemfrees)

	var items map[int]map[string]int = make(map[int]map[string]int)
	for _, free := range itemfrees {
		if _, ok := items[free.UID]; !ok {
			items[free.UID] = make(map[string]int)
			items[free.UID][free.Date] += free.Num
		}
	}

	store = src.Store.Rds(common.ItemRecordStore)
	for uid, data := range items {
		userid := strconv.Itoa(uid)
		for date, num := range data {
			store.HSet(userid, date, num)
		}
	}

	// 德州保险
	var insufrees []Stat
	src.Store.DBs(common.DataDsn).Raw("SELECT FROM_UNIXTIME(`created_at`, '%Y-%m-%d') AS 'date', `user_id` AS `uid`, `buy` AS 'num' FROM pk_data.insurance_flow_log_v2 WHERE(`created_at` >= ? AND `created_at` < ?)", lastime.Unix(), time.Now().Unix()).Scan(&insufrees)

	var insus map[int]map[string]int = make(map[int]map[string]int)
	for _, free := range insufrees {
		if _, ok := insus[free.UID]; !ok {
			insus[free.UID] = make(map[string]int)
			insus[free.UID][free.Date] += free.Num
		}
	}

	store = src.Store.Rds(common.InsuranceRecordStore)
	for uid, data := range items {
		userid := strconv.Itoa(uid)
		for date, num := range data {
			store.HSet(userid, date, num)
		}
	}

	src.Store.Rds(common.UpdateNotification).Del("service_close_date")
}
*/
