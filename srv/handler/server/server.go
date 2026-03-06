package server

import (
	"cms/srv/api-getaway/basic/config"
	"cms/srv/handler/request"
	__ "cms/srv/proto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ContentAdd(c *gin.Context) {
	var form request.ContentAdd
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数不正确",
		})
		return
	}
	_, err := config.ContentClient.AddCmsCategory(c, &__.AddCmsCategoryReq{
		Title:      form.Title,
		CateGoryId: int64(form.CateGoryId),
		Content:    form.Content,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数不正确",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "添加成功",
	})
	return
}
func GetById(c *gin.Context) {
	var form request.GetById
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数不正确",
		})
		return
	}
	r, err := config.ContentClient.GetCmsCategoryById(c, &__.GetCmsCategoryByIdReq{
		Id: int64(form.Id),
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "查询失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": r,
	})
	return
}
func Search(c *gin.Context) {
	var form request.Search
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数不正确",
		})
		return
	}
	r, err := config.ContentClient.SearchCmsCategory(c, &__.SearchCmsCategoryReq{
		Page: int64(form.Page),
		Size: int64(form.Size),
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "搜索失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "搜索成功",
		"data": r,
	})
	return
}
