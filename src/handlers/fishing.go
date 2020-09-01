package handlers

import (
	"encoding/json"
	"pkt_statistics/common"
	"pkt_statistics/library/logger"
	"pkt_statistics/model"
	"pkt_statistics/src"

	"github.com/nats-io/stan.go"
)

// 人数统计: 	 区域        日期      UID  Count
var buyers map[string]map[string]map[int]int

// FishingHandOverRecordHandler 钓鱼每手消息处理
func FishingHandOverRecordHandler(msg *stan.Msg) {
	var record model.FishHandOverRecord
	if err := json.Unmarshal(msg.Data, &record); err != nil {
		logger.ErrorF("FishingHandOverRecordHandler JSON Unmarshal error: %v", err)
		return
	}
	date := TodayDate()
	match := src.Store.Rds(common.FishingHandOverRecordStore)

	// 今日完成总手数
	match.HIncrBy("Count", date, 1)
	// 今日总下注统计
	match.HIncrBy("TotalBet", date, int64(record.TotalBet))
	// 今日总赔付统计
	match.HIncrBy("TotalPay", date, int64(record.TotalPay))
	// 今日下注区域统计
	match.HIncrBy("Bet11", date, int64(record.Bet11))
	match.HIncrBy("Bet12", date, int64(record.Bet12))
	match.HIncrBy("Bet13", date, int64(record.Bet13))
	match.HIncrBy("Bet14", date, int64(record.Bet14))
	match.HIncrBy("Bet21", date, int64(record.Bet21))
	match.HIncrBy("Bet22", date, int64(record.Bet22))
	match.HIncrBy("Bet23", date, int64(record.Bet23))
	match.HIncrBy("Bet31", date, int64(record.Bet31))
	match.HIncrBy("Bet32", date, int64(record.Bet32))
	match.HIncrBy("Bet33", date, int64(record.Bet33))
	match.HIncrBy("Bet34", date, int64(record.Bet34))
	match.HIncrBy("Bet35", date, int64(record.Bet35))
	// 今日赔付区域统计
	match.HIncrBy("Pay11", date, int64(record.Pay11))
	match.HIncrBy("Pay12", date, int64(record.Pay12))
	match.HIncrBy("Pay13", date, int64(record.Pay13))
	match.HIncrBy("Pay14", date, int64(record.Pay14))
	match.HIncrBy("Pay21", date, int64(record.Pay21))
	match.HIncrBy("Pay22", date, int64(record.Pay22))
	match.HIncrBy("Pay23", date, int64(record.Pay23))
	match.HIncrBy("Pay31", date, int64(record.Pay31))
	match.HIncrBy("Pay32", date, int64(record.Pay32))
	match.HIncrBy("Pay33", date, int64(record.Pay33))
	match.HIncrBy("Pay34", date, int64(record.Pay34))
	match.HIncrBy("Pay35", date, int64(record.Pay35))
	// 今日胜利区域统计
	match.HIncrBy("Win11", date, int64(record.Win11))
	match.HIncrBy("Win12", date, int64(record.Win12))
	match.HIncrBy("Win13", date, int64(record.Win13))
	match.HIncrBy("Win14", date, int64(record.Win14))
	match.HIncrBy("Win21", date, int64(record.Win21))
	match.HIncrBy("Win22", date, int64(record.Win22))
	match.HIncrBy("Win23", date, int64(record.Win23))
	match.HIncrBy("Win31", date, int64(record.Win31))
	match.HIncrBy("Win32", date, int64(record.Win32))
	match.HIncrBy("Win33", date, int64(record.Win33))
	match.HIncrBy("Win34", date, int64(record.Win34))
	match.HIncrBy("Win35", date, int64(record.Win35))
	// 今日区域未命中金额统计
	if record.Bet11 > 0 && record.Win11 == 0 {
		match.HIncrBy("Miss11", date, int64(record.Bet11))
		match.HIncrBy("TotalMiss", date, int64(record.Bet11))
	}
	if record.Bet12 > 0 && record.Win12 == 0 {
		match.HIncrBy("Miss12", date, int64(record.Bet12))
		match.HIncrBy("TotalMiss", date, int64(record.Bet11))
	}
	if record.Bet13 > 0 && record.Win13 == 0 {
		match.HIncrBy("Miss13", date, int64(record.Bet13))
		match.HIncrBy("TotalMiss", date, int64(record.Bet11))
	}
	if record.Bet14 > 0 && record.Win14 == 0 {
		match.HIncrBy("Miss14", date, int64(record.Bet14))
		match.HIncrBy("TotalMiss", date, int64(record.Bet11))
	}
	if record.Bet21 > 0 && record.Win21 == 0 {
		match.HIncrBy("Miss21", date, int64(record.Bet21))
		match.HIncrBy("TotalMiss", date, int64(record.Bet11))
	}
	if record.Bet22 > 0 && record.Win22 == 0 {
		match.HIncrBy("Miss22", date, int64(record.Bet22))
		match.HIncrBy("TotalMiss", date, int64(record.Bet11))
	}
	if record.Bet23 > 0 && record.Win23 == 0 {
		match.HIncrBy("Miss23", date, int64(record.Bet23))
		match.HIncrBy("TotalMiss", date, int64(record.Bet11))
	}
	if record.Bet31 > 0 && record.Win31 == 0 {
		match.HIncrBy("Miss31", date, int64(record.Bet31))
		match.HIncrBy("TotalMiss", date, int64(record.Bet11))
	}
	if record.Bet32 > 0 && record.Win32 == 0 {
		match.HIncrBy("Miss32", date, int64(record.Bet32))
		match.HIncrBy("TotalMiss", date, int64(record.Bet11))
	}
	if record.Bet33 > 0 && record.Win33 == 0 {
		match.HIncrBy("Miss33", date, int64(record.Bet33))
		match.HIncrBy("TotalMiss", date, int64(record.Bet11))
	}
	if record.Bet34 > 0 && record.Win34 == 0 {
		match.HIncrBy("Miss34", date, int64(record.Bet34))
		match.HIncrBy("TotalMiss", date, int64(record.Bet11))
	}
	if record.Bet35 > 0 && record.Win35 == 0 {
		match.HIncrBy("Miss35", date, int64(record.Bet35))
		match.HIncrBy("TotalMiss", date, int64(record.Bet11))
	}

	// 今日区域下注人数统计
	if _, ok := buyers["Buyers11"][date]; !ok {
		buyers["Buyers11"][date] = make(map[int]int)
	}
	if _, ok := buyers["Buyers12"][date]; !ok {
		buyers["Buyers12"][date] = make(map[int]int)
	}
	if _, ok := buyers["Buyers13"][date]; !ok {
		buyers["Buyers13"][date] = make(map[int]int)
	}
	if _, ok := buyers["Buyers14"][date]; !ok {
		buyers["Buyers14"][date] = make(map[int]int)
	}
	if _, ok := buyers["Buyers21"][date]; !ok {
		buyers["Buyers21"][date] = make(map[int]int)
	}
	if _, ok := buyers["Buyers22"][date]; !ok {
		buyers["Buyers22"][date] = make(map[int]int)
	}
	if _, ok := buyers["Buyers23"][date]; !ok {
		buyers["Buyers23"][date] = make(map[int]int)
	}
	if _, ok := buyers["Buyers31"][date]; !ok {
		buyers["Buyers31"][date] = make(map[int]int)
	}
	if _, ok := buyers["Buyers32"][date]; !ok {
		buyers["Buyers32"][date] = make(map[int]int)
	}
	if _, ok := buyers["Buyers33"][date]; !ok {
		buyers["Buyers33"][date] = make(map[int]int)
	}
	if _, ok := buyers["Buyers34"][date]; !ok {
		buyers["Buyers34"][date] = make(map[int]int)
	}
	if _, ok := buyers["Buyers35"][date]; !ok {
		buyers["Buyers35"][date] = make(map[int]int)
	}
	if _, ok := buyers["TotalBuyers"][date]; !ok {
		buyers["TotalBuyers"][date] = make(map[int]int)
	}

	for user, record := range record.UserRecordList {
		if record.Bet11 > 0 {
			buyers["Buyers11"][date][user]++
		}
		if record.Bet12 > 0 {
			buyers["Buyers12"][date][user]++
		}
		if record.Bet13 > 0 {
			buyers["Buyers13"][date][user]++
		}
		if record.Bet14 > 0 {
			buyers["Buyers14"][date][user]++
		}
		if record.Bet21 > 0 {
			buyers["Buyers21"][date][user]++
		}
		if record.Bet22 > 0 {
			buyers["Buyers22"][date][user]++
		}
		if record.Bet23 > 0 {
			buyers["Buyers23"][date][user]++
		}
		if record.Bet31 > 0 {
			buyers["Buyers31"][date][user]++
		}
		if record.Bet32 > 0 {
			buyers["Buyers32"][date][user]++
		}
		if record.Bet33 > 0 {
			buyers["Buyers33"][date][user]++
		}
		if record.Bet34 > 0 {
			buyers["Buyers34"][date][user]++
		}
		if record.Bet35 > 0 {
			buyers["Buyers35"][date][user]++
		}
		buyers["TotalBuyers"][date][user]++
	}

	match.HSet("Buyers11", date, len(buyers["Buyers11"][date]))
	match.HSet("Buyers12", date, len(buyers["Buyers12"][date]))
	match.HSet("Buyers13", date, len(buyers["Buyers13"][date]))
	match.HSet("Buyers14", date, len(buyers["Buyers14"][date]))
	match.HSet("Buyers21", date, len(buyers["Buyers21"][date]))
	match.HSet("Buyers22", date, len(buyers["Buyers22"][date]))
	match.HSet("Buyers23", date, len(buyers["Buyers23"][date]))
	match.HSet("Buyers31", date, len(buyers["Buyers31"][date]))
	match.HSet("Buyers32", date, len(buyers["Buyers32"][date]))
	match.HSet("Buyers33", date, len(buyers["Buyers33"][date]))
	match.HSet("Buyers34", date, len(buyers["Buyers34"][date]))
	match.HSet("Buyers35", date, len(buyers["Buyers35"][date]))
	match.HSet("TotalBuyers", date, len(buyers["TotalBuyers"][date]))

	// 备份Buyers统计数据入Redis
	if str, err := json.Marshal(buyers); err != nil {
		logger.ErrorF("json marshal buyers error: %v", err)
	} else {
		match.Set("Buyers", str, 0)
	}

	logger.InfoF("[钓鱼]Subject: %v, Date: %v MID: %v HID %v\n", msg.Subject, date, record.Mid, record.HandId)
}
