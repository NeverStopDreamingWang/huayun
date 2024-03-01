package role

import (
	"github.com/NeverStopDreamingWang/goi/model"
)

import (
	"github.com/NeverStopDreamingWang/goi/migrate"
)

func init() {
	// sqlite 数据库
	SQLite3Migrations := model.SQLite3MakeMigrations{
		DATABASES: []string{"default"},
		MODELS: []model.SQLite3Model{
			PermissionModel{},     // 权限表
			RoleModel{},           // 角色表
			RolePermissionModel{}, // 角色-权限表
		},
	}
	migrate.SQLite3Migrate(SQLite3Migrations)
}

// 权限表
type PermissionModel struct {
	Id     *int64  `field_name:"id" field_type:"INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT" json:"id"`
	Name   *string `field_name:"name" field_type:"TEXT NOT NULL UNIQUE" json:"name"`
	Method *string `field_name:"method" field_type:"TEXT NOT NULL" json:"method"`
	Uri    *string `field_name:"uri" field_type:"TEXT NOT NULL UNIQUE" json:"uri"`
}

func (PermissionModel) ModelSet() *model.SQLite3Settings {
	modelSettings := &model.SQLite3Settings{
		MigrationsHandler: model.MigrationsHandler{ // 迁移时处理函数
			BeforeFunc: nil, // 迁移之前处理函数
			AfterFunc:  nil, // 迁移之后处理函数
		},

		TABLE_NAME: "permission", // 设置表名

		// 自定义配置
		MySettings: model.MySettings{},
	}
	return modelSettings
}

// 角色表
type RoleModel struct {
	Id              *int64  `field_name:"id" field_type:"INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT" json:"id"`
	Name            *string `field_name:"name" field_type:"TEXT NOT NULL UNIQUE" json:"name"`
	Status          *uint8  `field_name:"status" field_type:"INTEGER NOT NULL DEFAULT 1" json:"status"`
	Remark          *string `field_name:"remark" field_type:"TEXT" json:"remark"`
	Create_Datetime *string `field_name:"create_datetime" field_type:"DATETIME NOT NULL" json:"create_datetime"`
	Update_Datetime *string `field_name:"update_datetime" field_type:"DATETIME" json:"update_datetime"`
}

func (RoleModel) ModelSet() *model.SQLite3Settings {
	modelSettings := &model.SQLite3Settings{
		MigrationsHandler: model.MigrationsHandler{ // 迁移时处理函数
			BeforeFunc: nil, // 迁移之前处理函数
			AfterFunc:  nil, // 迁移之后处理函数
		},

		TABLE_NAME: "role", // 设置表名

		// 自定义配置
		MySettings: model.MySettings{},
	}
	return modelSettings
}

// 角色-权限表
type RolePermissionModel struct {
	Role_Id         *int64  `field_name:"role_id" field_type:"INTEGER NOT NULL" json:"role_id"`
	Permission_Id   *int64  `field_name:"permission_id" field_type:"TEXT NOT NULL" json:"permission_id"`
	Create_Datetime *string `field_name:"create_datetime" field_type:"DATETIME NOT NULL" json:"create_datetime"`
}

func (RolePermissionModel) ModelSet() *model.SQLite3Settings {
	modelSettings := &model.SQLite3Settings{
		MigrationsHandler: model.MigrationsHandler{ // 迁移时处理函数
			BeforeFunc: nil, // 迁移之前处理函数
			AfterFunc:  nil, // 迁移之后处理函数
		},

		TABLE_NAME: "role_permission", // 设置表名

		// 自定义配置
		MySettings: model.MySettings{},
	}
	return modelSettings
}
