package model

// ItemRecord 道具
type ItemRecord struct {
	Uid      int     // 用户ID
	Mid      int     // 比赛ID
	PropId   int     // 道具ID
	PropType string  // 道具类型
	Cost     float64 // 消耗金币
	Utime    int     // 使用时间
	CoinType uint32  // 币种 0: 时间币 1: 扑克币
}
