package common

import "pkt_statistics/library/storage"

const (
	// ItemRecordStore 道具记录存储
	ItemRecordStore storage.AREA = 1
	// TexasHandOverRecordStore 德州每手记录存储
	TexasHandOverRecordStore storage.AREA = 2
	// InsuranceRecordStore 保险记录存储
	InsuranceRecordStore storage.AREA = 3
	// FishingHandOverRecordStore 钓鱼每手记录存储
	FishingHandOverRecordStore storage.AREA = 4
	// FishingUserHandOverRecordStore 钓鱼用户每手记录存储
	FishingUserHandOverRecordStore storage.AREA = 5

	// UpdateNotification 更新通知
	UpdateNotification storage.AREA = 15
)
