package src

import (
	"pkt_statistics/library/logger"
	"pkt_statistics/library/storage"
)

// Store 存储模块
var Store *storage.Storage

func init() {
	// 初始化存储模块
	var err error
	Store, err = storage.NewStorage()
	if err != nil {
		logger.FatalF("store init fail: %v", err)
	}
	logger.Info("store init successful")
}
