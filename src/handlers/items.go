package handlers

import (
	"encoding/json"
	"pkt_statistics/common"
	"pkt_statistics/library/logger"
	"pkt_statistics/model"
	"pkt_statistics/src"
	"strconv"
	"time"

	"github.com/nats-io/stan.go"
)

// TodayDate 今日日期
func TodayDate() string {
	return time.Now().Format("2006-01-02")
}

// ItemRecordHandler 道具消息处理
func ItemRecordHandler(msg *stan.Msg) {

	var record model.ItemRecord
	if err := json.Unmarshal(msg.Data, &record); err != nil {
		logger.ErrorF("ItemRecordHandler JSON Unmarshal error: %v", err)
		return
	}

	// 忽略时间币类型消息
	if record.CoinType == common.COIN_TYPE_TC {
		return
	}
	store := src.Store.Rds(common.ItemRecordStore)

	date := TodayDate()
	userid := strconv.Itoa(record.Uid)
	// 统计单用户消耗的道具数量
	if err := store.HIncrByFloat(userid, date, record.Cost).Err(); err != nil {
		logger.ErrorF("ItemRecordHandler redis IncrByFloat %v error: %v", userid, err)
	}
	logger.InfoF("[道具]Subject: %v, Date: %v UserID: %v Cost: %v", msg.Subject, date, userid, record.Cost)
}
