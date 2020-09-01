package model

import "pkt_statistics/common"

// HandOverRecord 每一手的数据发送给MQ
type HandOverRecord struct {
	Mid        int // 比赛ID
	GameKind   int
	MatchType  int
	UnionId    int
	ClubId     int
	Ante       int
	Sb         int
	Straddle   bool
	Process    *HandProc // 牌局一手流程
	CoinType   uint32
	InsDetails []*InsuranceDetail // 保险详情
}

// HandProc 每手牌详细流程
type HandProc struct {
	Hid        uint64              // 手牌id, 与gameNum相同
	HandTime   int64               // 当前手牌开始时间
	Players    map[uint32]*Player  // 当前手牌参与玩家, Key: userId
	Boardcards []uint32            // 公共牌(已发的)
	Pots       map[uint32]*PotInfo // 当前手牌底池信息, Key: potId
	//TotalPotNum uint64   // 总底池大小
	Insurance int64  // 保险盈利
	PotFee    uint64 // 底池抽水金额
	LuckyCard *LuckyCardInfo
	Over      bool         // 本手牌是否结束标志
	Rounds    []*RoundProc // 每一轮流程
	LookCards []uint32     //所有的公共牌(包括未发出来的,主要用于发发看)
}

// InsuranceDetail 保险详情(2020/05/28 修改)
type InsuranceDetail struct {
	UnionId           int
	ClubId            int
	UserId            int     // 买保险的用户
	AllInUserIds      []int   // All-in 用户
	Outs              int     // 购买 Outs 数量
	Buy               int64   // 购买金额
	Pay               int64   // 赔付金额
	Round             string  // turn 转牌、river 河牌
	PotId             int     // 0 主池、1 边池一、2 边池二
	PotChips          int64   // 底池筹码
	OverOuts          []int   // 购买的反超 outs
	DeuceOuts         []int   // 购买的平分 outs
	TotalOverOutsNum  int     // 总的反超 outs 数量
	TotalDeuceOutsNum int     // 总的平分 outs 数量
	Odds              float64 // 赔率
}

// Player player
type Player struct {
	Sid          int             // 座位id
	Uid          uint32          // 用户id
	Payload      string          `json:"-"` // 用户信息
	HoleCards    []uint32        // 手牌
	CardType     common.CardType // 手牌类型
	StartChips   uint64          // 牌局开始时的筹码
	Result       int64           // 盈利
	Insurance    int64           // 保险盈利
	Position     string          // 位置
	ShowSets     []bool          // 是否亮牌
	Ip           string
	Gps          *Gps
	ShowAllHands bool // 是否强制亮牌
	Vpip         int
}

// Gps gps
type Gps struct {
	Longitude float64 // 经度: 例 113
	Latitude  float64 // 纬度: 例 22
}

// PotInfo 每个底池的详细信息
type PotInfo struct {
	Pid      uint32                // pot id
	PotChips uint64                // 底池大小
	Players  map[uint32]uint64     // 每个玩家投入底池的具体额度
	Winners  map[uint32]*PotWinner // 本底池获胜者 Key: userId
}

// PotWinner winner
type PotWinner struct {
	Sid int
	Uid uint32
	Num uint64
}

// LuckyCardInfo LuckyCardInfo
type LuckyCardInfo struct {
	PotId     int   // 从哪个底池扣筹码加入到幸运牌奖池
	PotChips  int   // 从底池扣筹码加入到幸运牌奖池
	LuckyCard []int // Lucky Card
	//LuckyPot  int64 // Lucky Pot
}

// RoundProc 每轮玩家行动流程
type RoundProc struct {
	Gstate      int    // 阶段, 翻前，翻后，转牌，河牌，showdown
	RoundTime   int64  // 阶段开始时间
	RoundPlayer uint32 // 本轮参与人数
	RoundPot    uint64 // 本轮总底池
	//RoundInsurance uint64 // 本轮保险
	//RoundItem      uint32 // 本轮道具使用次数
	Actions []*ActionProc // 本轮玩家动作流程
}

// ActionProc 本轮玩家动作流程
type ActionProc struct {
	Uid        uint32 // 玩家id
	Schips     uint64 // 动作前筹码
	Chips      uint64 // 操作行为的筹码数
	Echips     uint64 // 操作后的筹码数
	Epot       uint64 // 操作后的总底池(不包含底池费)
	Opt        int    // 操作类型
	Shove      bool   // 是否all-in
	ValidRaise bool   // 是否有效加注
}
