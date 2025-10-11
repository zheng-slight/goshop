package models

//收货地址相关

type Address struct {
	Id             int    `json:"id"`
	Uid            int    `json:"uid"`
	Phone          string `json:"phone"`
	Name           string `json:"name"`
	Address        string `json:"address"`
	DefaultAddress int    `json:"default_address"`  //默认地址:0 否, 1 是
	AddTime        int    `json:"add_time"`
}

func (Address) TableName() string {
	return "address"
}
