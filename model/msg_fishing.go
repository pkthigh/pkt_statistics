package model

// FishHandOverRecord 牌局结束时
type FishHandOverRecord struct {
	Mid         int         // 牌局ID
	HandId      int         // 手数ID
	BeginTime   int64       // 牌局开始时间
	EndTime     int64       // 牌局结束时间
	Conf        *FishDetail // 配置
	PublicCards []uint32    // 公共牌
	FisherCards []uint32    // 渔夫手牌
	SharkCards  []uint32    // 鲨鱼手牌
	CardType    int         // 赢家牌型

	TotalBet int // 总下注
	TotalPay int // 总赔付

	Bet11 int // 鲨鱼赢下注额
	Bet12 int // 渔夫赢下注额
	Bet13 int // 平局下注额
	Bet14 int // 渔夫对2或者高牌赢下注额
	Bet21 int // 任一人手牌-对子下注额
	Bet22 int // 任一人手牌-同花/连牌/同花连牌下注额
	Bet23 int // 任一人手牌-对A下注额
	Bet31 int // 获胜牌型-高牌/一对下注额
	Bet32 int // 获胜牌型-两对下注额
	Bet33 int // 获胜牌型-三条/顺子/同花下注额
	Bet34 int // 获胜牌型-葫芦下注额
	Bet35 int // 获胜牌型-金刚/同花顺/皇同下注额

	Pay11 int // 鲨鱼赢赔付额
	Pay12 int // 渔夫赢赔付额
	Pay13 int // 平局赔付额
	Pay14 int // 渔夫对2或者高牌赢赔付额
	Pay21 int // 任一人手牌-对子赔付额
	Pay22 int // 任一人手牌-同花/连牌/同花连牌赔付额
	Pay23 int // 任一人手牌-对A赔付额
	Pay31 int // 获胜牌型-高牌/一对赔付额
	Pay32 int // 获胜牌型-两对赔付额
	Pay33 int // 获胜牌型-三条/顺子/同花赔付额
	Pay34 int // 获胜牌型-葫芦赔付额
	Pay35 int // 获胜牌型-金刚/同花顺/皇同赔付额

	Win11 int // 鲨鱼赢(0不中 1中)
	Win12 int // 渔夫赢(0不中 1中)
	Win13 int // 平局(0不中 1中)
	Win14 int // 渔夫对2或者高牌赢(0不中 1中)
	Win21 int // 任一人手牌-对子(0不中 1中)
	Win22 int // 任一人手牌-同花/连牌/同花连牌(0不中 1中)
	Win23 int // 任一人手牌-对A(0不中 1中)
	Win31 int // 获胜牌型-高牌/一对(0不中 1中)
	Win32 int // 获胜牌型-两对(0不中 1中)
	Win33 int // 获胜牌型-三条/顺子/同花(0不中 1中)
	Win34 int // 获胜牌型-葫芦(0不中 1中)
	Win35 int // 获胜牌型-金刚/同花顺/皇同(0不中 1中)

	UserRecordList map[int]*FishUserRecord // 用户战绩列表
}

// FishUserRecord 用户战绩列表
type FishUserRecord struct {
	Uid          int
	ClubId       int // 所属俱乐部ID
	UnionId      int // 所属联盟ID
	SuperUnionId int // 所属超级联盟ID
	TotalBet     int // 总下注
	TotalPay     int // 总赔付

	Bet11 int // 鲨鱼赢下注额
	Bet12 int // 渔夫赢下注额
	Bet13 int // 平局下注额
	Bet14 int // 渔夫对2或者高牌赢下注额
	Bet21 int // 任一人手牌-对子下注额
	Bet22 int // 任一人手牌-同花/连牌/同花连牌下注额
	Bet23 int // 任一人手牌-对A下注额
	Bet31 int // 获胜牌型-高牌/一对下注额
	Bet32 int // 获胜牌型-两对下注额
	Bet33 int // 获胜牌型-三条/顺子/同花下注额
	Bet34 int // 获胜牌型-葫芦下注额
	Bet35 int // 获胜牌型-金刚/同花顺/皇同下注额

	Pay11 int // 鲨鱼赢赔付额
	Pay12 int // 渔夫赢赔付额
	Pay13 int // 平局赔付额
	Pay14 int // 渔夫对2或者高牌赢赔付额
	Pay21 int // 任一人手牌-对子赔付额
	Pay22 int // 任一人手牌-同花/连牌/同花连牌赔付额
	Pay23 int // 任一人手牌-对A赔付额
	Pay31 int // 获胜牌型-高牌/一对赔付额
	Pay32 int // 获胜牌型-两对赔付额
	Pay33 int // 获胜牌型-三条/顺子/同花赔付额
	Pay34 int // 获胜牌型-葫芦赔付额
	Pay35 int // 获胜牌型-金刚/同花顺/皇同赔付额
}

// FishDetail 配置详情
type FishDetail struct {
	Area *FishArea `json:"area" toml:"area"`
	Prop *FishProp `json:"prop" toml:"prop"`
}

// FishProp 钓鱼分成比例配置
type FishProp struct {
	SuperUnionProp int `json:"super_union_prop" toml:"superUnionProp"` // 超级联盟分成比例
	UnionProp      int `json:"union_prop" toml:"unionProp"`            // 联盟分成比例
	ClubProp       int `json:"club_prop" toml:"clubProp"`              // 俱乐部分成比例
}

// FishArea 钓鱼每个区域配置
type FishArea struct {
	SharkWin          AreaInfo `json:"shark_win" toml:"sharkWin"`
	FisherManWin      AreaInfo `json:"fisherman_win" toml:"fisherManWin"`
	Draw              AreaInfo `json:"draw" toml:"draw"`
	Fisher22OrHighWin AreaInfo `json:"fisher_22_or_high_win" toml:"fisher22OrHighWin"`

	HandOnePair         AreaInfo `json:"hand_one_pair" toml:"handOnePair"`
	HandFlushOrStraight AreaInfo `json:"hand_flush_or_straight" toml:"handFlushOrStraight"`
	HandAA              AreaInfo `json:"hand_aa" toml:"handAA"`

	HighOrOnePairWin          AreaInfo `json:"high_or_one_pair_win" toml:"highOrOnePairWin"`
	TwoPairWin                AreaInfo `json:"two_pair_win" toml:"twoPairWin"`
	TripsOrFlushOrStraightWin AreaInfo `json:"trips_or_flush_or_straight_win" toml:"tripsOrFlushOrStraightWin"`
	FullHouseWin              AreaInfo `json:"full_house_win" toml:"fullHouseWin"`
	QuadsOrFlushStraightWin   AreaInfo `json:"quads_or_flush_straight_win" toml:"quadsOrFlushStraightWin"`
}

// AreaInfo 每个区域配置信息
type AreaInfo struct {
	Odds   int   `json:"odds" toml:"odds"`      // 赔率, 乘以100
	MaxBet int64 `json:"max_bet" toml:"maxBet"` // 限红
}
