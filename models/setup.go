package models

var (
	ROLE_USER = Role{
		Id:          1,
		Name:        "USER",
		Description: "普通用户",
	}

	ROLE_ADMIN = Role{
		Id:          2,
		Name:        "ADMIN",
		Description: "管理员",
	}

	ROLE_SUPER = Role{
		Id:          3,
		Name:        "SUPER_ADMIN",
		Description: "超级管理员",
	}

	ROLES = []Role{ROLE_USER, ROLE_ADMIN, ROLE_SUPER}
)
