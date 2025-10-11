package models

//角色-权限 关联表

type RoleAccess struct {
	AccessId int
	RoleId   int
}

func (RoleAccess) TableName() string {
	return "role_access"
}
