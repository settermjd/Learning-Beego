package models

type Article struct {
	Id     int    `form:"-"`
	Name   string `form:"name,text,name:" valid:"MinSize(5);MaxSize(20)"`
	Client string `form:"client,text,client:"`
	Url    string `form:"url,text,url:"`
}

func (a *Article) TableName() string {
	return "articles"
}
