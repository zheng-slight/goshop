package models

//商品图片

type GoodsImage struct {
	Id      int `json:"id"`  // 反射操作,让其与表单name一致
	GoodsId int `json:"goods_id"`		//商品id
	ImgUrl  string `json:"img_url"`	//图片保存路径:一般只会保存类似于/static/upload/20230313/xxx.png这种格式
	ColorId int `json:"color_id"`		//颜色id
	Sort    int `json:"srt"`
	AddTime int `json:"add_time"`
	Status  int `json:"status"`
}

func (GoodsImage) TableName() string {
	return "goods_image"
}
