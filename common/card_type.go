package common

const (
	COIN_TYPE_TC  = 0 // 时间币
	COIN_TYPE_PKC = 1 // 微级币
)

// CardType 牌型
type CardType int

const (
	CTYPE_NONE          = 0  //无牌型
	CTYPE_HIGHCARD      = 1  //高牌
	CTYPE_ONEPAIR       = 2  //一对
	CTYPE_TWOPAIR       = 3  //两对
	CTYPE_THREEKIND     = 4  //三条
	CTYPE_STRAIGHT      = 5  //顺子
	CTYPE_FLUSH         = 6  //同花
	CTYPE_FULLHOUSE     = 7  //葫芦
	CTYPE_KINGKONG      = 8  //四条(金刚)
	CTYPE_FLUSHSTRAIGHT = 9  //同花顺
	CTYPE_ROYALFLUSH    = 10 //皇家同花顺
)

func (ct CardType) String() string {
	switch ct {
	case CTYPE_NONE:
		return "CTYPE_NONE" //无牌型
	case CTYPE_HIGHCARD:
		return "CTYPE_HIGHCARD" //高牌
	case CTYPE_ONEPAIR:
		return "CTYPE_ONEPAIR" //一对
	case CTYPE_TWOPAIR:
		return "CTYPE_TWOPAIR" //两对
	case CTYPE_THREEKIND:
		return "CTYPE_THREEKIND" //三条
	case CTYPE_STRAIGHT:
		return "CTYPE_STRAIGHT" //顺子
	case CTYPE_FLUSH:
		return "CTYPE_FLUSH" //同花
	case CTYPE_FULLHOUSE:
		return "CTYPE_FULLHOUSE" //葫芦
	case CTYPE_KINGKONG:
		return "CTYPE_KINGKONG" //四条(金刚)
	case CTYPE_FLUSHSTRAIGHT:
		return "CTYPE_FLUSHSTRAIGHT" //同花顺
	case CTYPE_ROYALFLUSH:
		return "CTYPE_ROYALFLUSH" //皇家同花顺
	default:
		return "CTYPE_UNKNOWN" //未知
	}
}
