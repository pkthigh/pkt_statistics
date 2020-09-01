package handlers

import (
	"encoding/json"
	"pkt_statistics/common"
	"pkt_statistics/library/logger"
	"pkt_statistics/model"
	"pkt_statistics/src"
	"strconv"

	"github.com/nats-io/stan.go"
)

// TexasHandOverRecordHandler 德州每手消息处理
func TexasHandOverRecordHandler(msg *stan.Msg) {
	var record model.HandOverRecord
	if err := json.Unmarshal(msg.Data, &record); err != nil {
		logger.ErrorF("TexasHandOverRecordHandler JSON Unmarshal error: %v", err)
		return
	}
	if record.CoinType != common.COIN_TYPE_PKC {
		return
	}

	// 忽略非联盟局消息
	if record.UnionId == 0 {
		return
	}

	date := TodayDate()
	texas := src.Store.Rds(common.TexasHandOverRecordStore)
	insur := src.Store.Rds(common.InsuranceRecordStore)
	notice := src.Store.Rds(common.UpdateNotification)

	// 记录手数
	var players []string
	for key := range record.Process.Players {
		userid := strconv.Itoa(int(key))
		players = append(players, userid)

		// 手数统计
		if err := texas.HIncrBy(userid, date, 1).Err(); err != nil {
			logger.ErrorF("TexasHandOverRecordHandler redis HIncrBy %v error: %v", userid, err)
		}
		logger.InfoF("[德州手数]Subject: %v, Date: %v UserID: %v", msg.Subject, date, userid)
	}

	// 记录保险
	for _, ins := range record.InsDetails {
		userid := strconv.Itoa(ins.UserId)
		if err := insur.HIncrBy(userid, date, ins.Buy).Err(); err != nil {
			logger.ErrorF("InsuranceRecordStore redis IncrBy %v error: %v", userid, err)
		}
		logger.InfoF("[德州保险]Subject: %v, Date: %v UserID: %v", msg.Subject, date, userid)
	}

	for _, player := range players {
		// 德州更新通知
		if err := notice.Publish(common.PkcHandOverSubject.String(), player).Err(); err != nil {
			logger.ErrorF("[德州手数更新通知失败]UpdateNotification failed, Date: %v UserID: %v Err: %v", date, player, err)
		}
		// logger.InfoF("[德州手数更新通知成功]UpdateNotification failed, Date: %v UserID: %v", date, player)
	}

}
