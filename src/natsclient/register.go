package natsclient

import (
	"pkt_statistics/common"
	"pkt_statistics/library/logger"
	"pkt_statistics/src/handlers"
)

// 注册每个Subject的消息处理函数
func (nats *Nats) register() {
	logger.Info("start listen message")
	for {
		msg := <-nats.msgs

		switch msg.Subject {

		// 道具消息处理
		case common.ItemSubject.String():
			handlers.ItemRecordHandler(msg)

		// 德州消息处理
		case common.PkcHandOverSubject.String():
			handlers.TexasHandOverRecordHandler(msg)

			// 钓鱼消息处理
			// case common.FishHandOverSubject.String():
			// 	handlers.FishingHandOverRecordHandler(msg)
		}
	}
}
