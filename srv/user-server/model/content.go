package model

import (
	__ "cms/srv/proto"

	"gorm.io/gorm"
)

type CmsContent struct {
	gorm.Model
	Title      string `gorm:"type:varchar(30);comment:标题"`
	CateGoryId int    `gorm:"type:int(11);comment:分类Id"`
	Content    string `gorm:"type:varchar(30);comment:内容Id"`
}

func (c *CmsContent) ContentAdd(db *gorm.DB) error {
	return db.Create(&c).Error
}

func (c *CmsContent) GetById(db *gorm.DB, id int64) error {
	return db.Model(&CmsContent{}).Where("id=?", id).First(&c).Error
}

func (c *CmsContent) Search(db *gorm.DB, in *__.SearchCmsCategoryReq) ([]CmsContent, error) {
	var list []CmsContent
	if in.Page <= 0 || in.Page > 3 {
		in.Page = 1
	}
	if in.Size <= 0 || in.Size > 3 {
		in.Size = 1
	}
	offset := (in.Page - 1) * in.Size
	err := db.Model(&CmsContent{}).Offset(int(offset)).Limit(int(in.Size)).Find(&list).Error
	return list, err
}
