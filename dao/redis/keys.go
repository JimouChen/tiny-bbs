package redis

// redis zset 的key使用命名空间的形式，方便查询和拆分
const (
	KeyPrefix          = "bbs:" // 公共前缀
	KeyPostTimeZSet    = "post:time"
	KeyPostScoreZSet   = "post:score"
	KeyPostVotedZSetPf = "post:voted:" //记录用户及投票类型，参数是post id
)

func GetRedisZKey(key string) string {
	return KeyPrefix + key
}
