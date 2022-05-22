package logic

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/models"
	"bluebell/pkg/snowflake"

	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {
	//1.为帖子生成postid
	p.ID = int64(snowflake.GenID())
	//2.保存到数据库
	err = mysql.CreatePost(p)
	if err != nil {
		return err
	}
	//创建帖子生成时间
	err = redis.CreatePost(p.ID, p.CommunityID)
	return
	// 3. 返回
}

func GetPostById(pid int64) (data *models.ApiPostDetail, err error) {
	post, err := mysql.GetPostById(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostById(pid) failed",
			zap.Int64("pid", pid),
			zap.Error(err))
		return
	}
	user, err := mysql.GetUserById(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserById(post.AuthorID) failed",
			zap.Int64("author_id", post.AuthorID),
			zap.Error(err))
		return
	}
	// 根据社区id查询社区详细信息
	community, err := mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetUserById(post.AuthorID) failed",
			zap.Int64("community_id", post.CommunityID),
			zap.Error(err))
		return
	}
	data = &models.ApiPostDetail{
		AuthorName:      user.Username,
		Post:            post,
		CommunityDetail: community,
	}
	return
}

// GetPostList 分页获取社区列表
func GetPostList(page int64, size int64) (data []*models.ApiPostDetail, err error) {
	posts, err := mysql.GetPostList(page, size)
	if err != nil {
		return nil, err
	}
	data = make([]*models.ApiPostDetail, 0, len(posts))
	//遍历获得的每一条帖子
	for _, post := range posts {
		// 根据作者id查询作者信息
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorID) failed",
				zap.Int64("author_id", post.AuthorID),
				zap.Error(err))
			continue
		}
		// 根据社区id查询社区详细信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorID) failed",
				zap.Int64("community_id", post.CommunityID),
				zap.Error(err))
			continue
		}
		postDetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: community,
		}
		//加入 data切片中
		data = append(data, postDetail)
	}
	return
}
func GetPostList2(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	//2.从redis中查询id列表
	ids, err := redis.GetPostIDsInorder(p)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostIDsInOrder(p) return 0 data")
		return
	}
	//3.从mysql中根据查询到的id查询帖子详细信息
	posts, err := mysql.GetPostList2(ids)
	if err != nil {
		return
	}
	//提前分配好空间提高效率
	data = make([]*models.ApiPostDetail, 0, len(posts))
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return
	}
	//遍历获得的每一条帖子
	for idx, post := range posts {
		// 根据作者id查询作者信息
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorID) failed",
				zap.Int64("author_id", post.AuthorID),
				zap.Error(err))
			continue
		}
		// 根据社区id查询社区详细信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorID) failed",
				zap.Int64("community_id", post.CommunityID),
				zap.Error(err))
			continue
		}
		postDetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			VoteNum:         voteData[idx],
			Post:            post,
			CommunityDetail: community,
		}
		//加入 data切片中
		data = append(data, postDetail)
	}
	return
}

func GetCommunityPostList(p *models.ParamCommunityPostList) (data []*models.ApiPostDetail, err error) {
	ids, err := redis.GetCommunityPostIDsInOrder(p)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostIDsInOrder(p) return 0 data")
		return
	}
	//3.从mysql中根据查询到的id查询帖子详细信息
	posts, err := mysql.GetPostList2(ids)
	if err != nil {
		return
	}
	//提前分配好空间提高效率
	data = make([]*models.ApiPostDetail, 0, len(posts))
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return
	}
	//遍历获得的每一条帖子
	for idx, post := range posts {
		// 根据作者id查询作者信息
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorID) failed",
				zap.Int64("author_id", post.AuthorID),
				zap.Error(err))
			continue
		}
		// 根据社区id查询社区详细信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorID) failed",
				zap.Int64("community_id", post.CommunityID),
				zap.Error(err))
			continue
		}
		postDetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			VoteNum:         voteData[idx],
			Post:            post,
			CommunityDetail: community,
		}
		//加入 data切片中
		data = append(data, postDetail)
	}
	return
}
