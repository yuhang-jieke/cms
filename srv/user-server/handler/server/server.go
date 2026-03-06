package server

import (
	__ "cms/srv/proto"
	"cms/srv/user-server/basic/config"
	"cms/srv/user-server/model"
	"context"

	"github.com/pkg/errors"
)

type Server struct {
	__.UnimplementedEcommerceServiceServer
}

// SayHello implements helloworld.GreeterServer
func (s *Server) AddCmsCategory(_ context.Context, in *__.AddCmsCategoryReq) (*__.AddCmsCategoryResp, error) {
	content := model.CmsContent{
		Title:      in.Title,
		CateGoryId: int(in.CateGoryId),
		Content:    in.Content,
	}
	err := content.ContentAdd(config.DB)
	if err != nil {
		return nil, errors.New("添加失败")
	}
	return &__.AddCmsCategoryResp{
		Message: "添加成功",
	}, nil
}
func (s *Server) GetCmsCategoryById(_ context.Context, in *__.GetCmsCategoryByIdReq) (*__.GetCmsCategoryByIdResp, error) {
	var content model.CmsContent
	err := content.GetById(config.DB, in.Id)
	if err != nil {
		return nil, errors.New("查询失败")
	}

	return &__.GetCmsCategoryByIdResp{
		CmsCategory: &__.CmsCategory{
			Title:      content.Title,
			CateGoryId: int64(content.CateGoryId),
			Content:    content.Content,
		},
	}, nil
}
func (s *Server) SearchCmsCategory(_ context.Context, in *__.SearchCmsCategoryReq) (*__.SearchCmsCategoryResp, error) {
	var content model.CmsContent
	var list []model.CmsContent

	list, err := content.Search(config.DB, in)
	if err != nil {
		return nil, errors.New("搜索失败")
	}
	var contentList []*__.CmsCategory
	for _, cmsContent := range list {
		contentList = append(contentList, &__.CmsCategory{
			Title:      cmsContent.Title,
			CateGoryId: int64(cmsContent.CateGoryId),
			Content:    cmsContent.Content,
		})
	}
	return &__.SearchCmsCategoryResp{
		CmsCategory: contentList,
	}, nil
}
