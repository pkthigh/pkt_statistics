package handlers

import (
	"encoding/json"
	"errors"
	"pkt_statistics/common"
	"pkt_statistics/library/logger"
	"pkt_statistics/model"
	"pkt_statistics/src"
	"strconv"
	"time"

	"github.com/nats-io/stan.go"
)

// ActivityStatusInProgress 活动状态
const ActivityStatusInProgress = 2 //进行中
//GetTimeStr ...
func GetTimeStr(timestamp int64, timeFlag string) (timeStr string, err error) {
	if timestamp <= 0 {
		err = errors.New("时间戳不合法")
		return
	}
	if timeFlag == "HMS" {
		//截掉年月日
		timeStr = time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")[11:]
	} else if timeFlag == "YMD" {
		//截掉时分秒
		timeStr = time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")[0:10]
	}
	return
}

//GetTimeStamp ...
//获取时间戳:传年-月-日 时-分-秒
func GetTimeStamp(timeStr string) (timestamp int64) {
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc)
	timestamp = theTime.Unix()
	return
}

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

	date := TodayDate() //今日日期
	texas := src.Store.Rds(common.TexasHandOverRecordStore)
	insur := src.Store.Rds(common.InsuranceRecordStore)
	notice := src.Store.Rds(common.UpdateNotification)

	// 记录手数
	//1，当前进行中的活动信息
	var activity model.AcActivity
	if err := src.Store.DBs(common.ActivityDsnReadOlny).Table("ac_activity").
		Where("`status` = ?", ActivityStatusInProgress).Find(&activity).Error; err != nil {
		logger.ErrorF("query In-Progress Activity from db error: %v", err)
		return
	}
	//这里由于在设置活动任务时间的时候传的是：[当前年-月-日 用户设置的时-分-秒] 转换的时间戳
	//计算当日实际的任务开始时间与结束时间的时间戳
	taskStTimeStr, err := GetTimeStr(int64(activity.TaskBeginTime), "HMS") //实际的任务开始时间 时-分-秒
	if err != nil {
		logger.Error(err)
		return
	}
	taskStTime := GetTimeStamp(TodayDate() + " " + taskStTimeStr) //今日任务开始时间时间戳

	taskEtTimeStr, err := GetTimeStr(int64(activity.TaskEndTime), "HMS") //实际的任务结束时间 时-分-秒
	if err != nil {
		logger.Error(err)
		return
	}
	taskEtTime := GetTimeStamp(TodayDate() + " " + taskEtTimeStr) //今日任务结束时间时间戳

	//获取活动设置的手数
	// activityHandNum := activity.HandNum

	//查询这一手牌的时间是否是在当前进行中的活动任务时间段内进行的，是则统计进手数
	//2，比较玩家这一手牌的时候是否是在活动任务时间区间
	var players []string
	if record.Process.HandTime >= taskStTime && record.Process.HandTime <= taskEtTime {
		for key := range record.Process.Players {
			userid := strconv.Itoa(int(key))
			players = append(players, userid)
			// 手数统计
			if err := texas.HIncrBy(userid, date, 1).Err(); err != nil {
				logger.ErrorF("TexasHandOverRecordHandler redis HIncrBy %v error: %v", userid, err)
			}
			logger.InfoF("[德州手数]Subject: %v, Date: %v UserID: %v", msg.Subject, date, userid)

			// //如果玩家手数今日已经满足，则存入redis一次抽奖资格
			// //获取玩家今日所有手数
			// playHandNum, err := texas.HGet(userid, date).Int64()
			// if err != nil {
			// 	logger.ErrorF("TexasHandOverRecordHandler从redis获取玩家当日手数失败,userid:%v,date:%v,err:%v", userid, date, err)
			// }
			// if playHandNum >= int64(activityHandNum) {
			// 	//先查询今日抽奖资格，如果已经存在，说明今日的抽奖资格已经存入redis，不应该再继续累加抽奖资格
			// 	//qualification:资格
			// 	result := texas.HGet(userid+"qualification"+strconv.Itoa(int(activity.ID)), date)
			// 	if result.Err() == redis.Nil {
			// 		//存入redis一次抽奖资格
			// 		if err := texas.HIncrBy(userid+"qualification"+strconv.Itoa(int(activity.ID)), date, 1).Err(); err != nil {
			// 			logger.ErrorF("TexasHandOverRecordHandler redis HIncrBy %v error: %v", userid, err)
			// 		}
			// 	}
			// }

		}
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
		logger.InfoF("[德州手数更新通知成功]UpdateNotification failed, Date: %v UserID: %v", date, player)
	}

}
