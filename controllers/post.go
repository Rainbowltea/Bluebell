package controllers

import (
	"bluebell/logic"
	"bluebell/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreatePostHandler(c *gin.Context) {
	//参数校验
	//业务执行：创建帖子
	//返回响应
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("c.ShouldBindJSON(p) error", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//从c中获取用户id
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = userID
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
func PostDetailHandler(c *gin.Context) {
	//1.参数验证，
	//2.事务处理
	//返回

	//取出帖子数据
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//根据id取出数据
	data, err := logic.GetPostById(pid)
	if err != nil {
		zap.L().Error("logic.GetPostById(pid) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, data)
}
func PostListHandler(c *gin.Context) {
	//参数校验
	//事务处理
	//返回值

	//获取分页数和每页内容
	page, size := getPageInfo(c)
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, data)
}

//根据前端传来的参数动态（分数||创建的时间先后）获取帖子列表
//1.获取参数
//2.从redis中获取id
//3.redis中的id从mysql中获取帖子详细信息
func PostListHandler2(c *gin.Context) {
	//参数校验
	//事务处理
	//返回值
	p := &models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderTime, // magic string
	}
	if err := c.ShouldBind(p); err != nil {
		zap.L().Error("PostListHandler2 with invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	data, err := logic.GetPostList2(p)
	if err != nil {
		zap.L().Error("logic.GetPostList2() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, data)
}
