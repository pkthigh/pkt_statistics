package common

import "pkt_statistics/library/storage"

const (
	// DataDsn 流水数据库
	DataDsn storage.MYSQL = "data_dsn"
	// GamelogDsn 游戏日志数据库
	GamelogDsn storage.MYSQL = "gamelog_dsn"
	// ActivityDsnReadOlny 活动从库
	ActivityDsnReadOlny storage.MYSQL = "activity_dsn_read_only"
)
