package redis

import (
	"bluebell/models"

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
