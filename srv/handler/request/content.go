package request

type ContentAdd struct {
	Title      string `form:"title"   binding:"required"`
	CateGoryId int    `form:"cate_gory_id"  binding:"required"`
	Content    string `form:"content"  binding:"required"`
}
type GetById struct {
	Id int `form:"id"  binding:"required"`
}
type Search struct {
	Page int `form:"page"  binding:"required"`
	Size int `form:"size"  binding:"required"`
}
