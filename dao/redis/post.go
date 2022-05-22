package redis

import (
	"bluebell/models"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

func GetPostIDsInorder(p *models.ParamPostList) ([]string, error) {
	//从redis中获取id
	var key string
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	} else {
		key = getRedisKey(KeyPostTimeZSet)
	}
	start := (p.Page - 1) * p.Size
	end := start - 1 + p.Size
	//查询
	return Rdb.ZRevRange(key, start, end).Result()
}

//根据ids查询每篇帖子的投赞成票的数据
func GetPostVoteData(ids []string) (data []int64, err error) {
	//data=make([]64,0,len(ids))
	// for _,id:=range ids{
	// 	key:=getRedisKey(KeyPostVotedZSetPF+id)
	// 	v1:=Rdb.ZCount(key,"1","1").Val()
	// 	data=append(data, v1)
	// }
	// return
	pipeline := Rdb.Pipeline()
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedZSetPF + id)
		pipeline.ZCount(key, "1", "1")
	}
	cmders, err := pipeline.Exec()
	if err != nil {
		return nil, err
	}
	data = make([]int64, 0, len(cmders))
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return
}

// GetCommunityPostIDsInOrder 按社区查询ids
func GetCommunityPostIDsInOrder(p *models.ParamCommunityPostList) ([]string, error) {
	//传入page,size,社区号，排序方式
	//使用zinterstore 把分区的帖子set与帖子分数的zset生成一个新的zset
	//使用新的zset来进行逻辑处理
	var orderkey string
	if p.Order == models.OrderScore {
		orderkey = getRedisKey(KeyPostScoreZSet)
	} else {
		orderkey = getRedisKey(KeyPostTimeZSet)
	}
	//社区key,存储帖子id
	ckey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(int(p.CommunityID)))

	// 利用缓存key减少zinterstore执行的次数
	//orderkey存储以帖子id为成员的按分数或者时间排序的Zet
	key := orderkey + strconv.Itoa(int(p.CommunityID))
	if Rdb.Exists(key).Val() < 1 {
		// 不存在，需要计算
		pipeline := Rdb.Pipeline()
		pipeline.ZInterStore(key, redis.ZStore{
			Aggregate: "MAX",
		}, ckey, orderkey) // zinterstore 计算
		pipeline.Expire(key, 60*time.Second) // 设置超时时间
		_, err := pipeline.Exec()
		if err != nil {
			return nil, err
		}
	}
	start := (p.Page - 1) * p.Size
	end := start - 1 + p.Size
	//查询
	return Rdb.ZRevRange(key, start, end).Result()
}
