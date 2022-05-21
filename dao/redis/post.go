package redis

import "bluebell/models"

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
