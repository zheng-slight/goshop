package models

//轮播图

type Focus struct {
	Id        int
	Title     string
	FocusType int  // 1 网站, 2 app, 3 小程序
	FocusImg  string
	Link      string
	Sort      int
	Status    int
	AddTime   int
}

func (Focus) TableName() string {
	return "focus"
}
