package redis

import (
	"errors"
	"github.com/go-redis/redis"
	"math"
	"time"
)

const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePerVote     = 234 // 每一票值多少分
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
	ErrVoteRepeated   = errors.New("不允许重复投票")
)

func VoteForPost(userId, postId string, direction float64) error {
	// 判断投票
	// 需要去redis拿到帖子发布时间
	postTime := rdb.ZScore(GetRedisZKey(KeyPostTimeZSet), postId).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrVoteTimeExpire
	}
	// 更新帖子分数
	// 先查当前用户给当前帖子的投票记录
	ov := rdb.ZScore(GetRedisZKey(KeyPostVotedZSetPf+postId), userId).Val()

	// 更新：如果这一次投票的值和之前保存的值一致，就提示不允许重复投票
	if direction == ov {
		return ErrVoteRepeated
	}
	var op float64
	if direction > ov {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(ov - direction) // 计算两次投票的差值
	pipeline := rdb.TxPipeline()
	pipeline.ZIncrBy(GetRedisZKey(KeyPostScoreZSet), op*diff*scorePerVote, postId)

	// 记录用户为帖子投票的数据
	if direction == 0 {
		pipeline.ZRem(GetRedisZKey(KeyPostVotedZSetPf+postId), userId)
	} else {
		pipeline.ZAdd(GetRedisZKey(KeyPostVotedZSetPf+postId), redis.Z{
			Score:  direction, // 赞成票还是反对票
			Member: userId,
		})
	}
	_, err := pipeline.Exec()
	return err
}

func CreatePost(postId int64) (err error) {
	pipeline := rdb.TxPipeline()

	//将创建的帖子时间和分数写入redis，要求同时成功，所以加入事务
	pipeline.ZAdd(GetRedisZKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postId,
	})

	pipeline.ZAdd(GetRedisZKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postId,
	})

	_, err = pipeline.Exec()
	return
}
