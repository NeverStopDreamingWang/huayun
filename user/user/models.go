package user

import (
	"github.com/NeverStopDreamingWang/goi/migrate"
	"github.com/NeverStopDreamingWang/goi/model"
	"time"
)

func init() {
	// sqlite 数据库
	SQLite3Migrations := model.SQLite3MakeMigrations{
		DATABASES: []string{"default"},
		MODELS: []model.SQLite3Model{
			UserModel{},
			UserRoleModel{},
		},
	}
	migrate.SQLite3Migrate(SQLite3Migrations)
}

// 用户表
type UserModel struct {
	Id                *int64     `field_name:"id" field_type:"INTEGER NOT NULL" json:"id"`
	Username          *string    `field_name:"username" field_type:"TEXT NOT NULL UNIQUE" json:"username"`
	Password          *string    `field_name:"password" field_type:"TEXT NOT NULL" json:"password"`
	Status            *uint8     `field_name:"status" field_type:"INTEGER NOT NULL DEFAULT 1" json:"status"`
	Security_Entrance *string    `field_name:"security_entrance" field_type:"TEXT NOT NULL" json:"security_entrance"`
	Remark            *string    `field_name:"remark" field_type:"TEXT" json:"remark"`
	Create_Datetime   *time.Time `field_name:"create_datetime" field_type:"DATETIME NOT NULL" json:"create_datetime"`
	Update_Datetime   *time.Time `field_name:"update_datetime" field_type:"DATETIME" json:"update_datetime"`
}

func (UserModel) ModelSet() *model.SQLite3Settings {
	modelSettings := &model.SQLite3Settings{
		MigrationsHandler: model.MigrationsHandler{ // 迁移时处理函数
			BeforeFunc: nil, // 迁移之前处理函数
			AfterFunc:  nil, // 迁移之后处理函数
		},

		TABLE_NAME: "user", // 设置表名

		// 自定义配置
		MySettings: model.MySettings{},
	}
	return modelSettings
}

// 用户-角色表
type UserRoleModel struct {
	User_Id         *int64     `field_name:"user_id" field_type:"INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT" json:"user_id"`
	Role_Id         *int64     `field_name:"role_id" field_type:"INTEGER NOT NULL" json:"role_id"`
	Create_Datetime *time.Time `field_name:"create_datetime" field_type:"DATETIME NOT NULL" json:"create_datetime"`
}

func (UserRoleModel) ModelSet() *model.SQLite3Settings {
	modelSettings := &model.SQLite3Settings{
		MigrationsHandler: model.MigrationsHandler{ // 迁移时处理函数
			BeforeFunc: nil, // 迁移之前处理函数
			AfterFunc:  nil, // 迁移之后处理函数
		},

		TABLE_NAME: "user_role", // 设置表名

		// 自定义配置
		MySettings: model.MySettings{},
	}
	return modelSettings
}

// 用户-网站表
type UserWebsiteModel struct {
	User_Id            *int64     `field_name:"user_id" field_type:"INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT" json:"user_id"`
	Website_Id         *int64     `field_name:"website_id" field_type:"INTEGER NOT NULL" json:"website_id"`
	Website_Table_Name *string    `field_name:"website_table_name" field_type:"TEXT NOT NULL" json:"website_table_name"`
	Status             *uint8     `field_name:"status" field_type:"INTEGER NOT NULL DEFAULT 1" json:"status"`
	Create_User_Id     *int64     `field_name:"create_user_id" field_type:"INTEGER" json:"create_user_id"`
	Create_Datetime    *time.Time `field_name:"create_datetime" field_type:"DATETIME NOT NULL" json:"create_datetime"`
}

func (UserWebsiteModel) ModelSet() *model.SQLite3Settings {
	modelSettings := &model.SQLite3Settings{
		MigrationsHandler: model.MigrationsHandler{ // 迁移时处理函数
			BeforeFunc: nil, // 迁移之前处理函数
			AfterFunc:  nil, // 迁移之后处理函数
		},

		TABLE_NAME: "user_website", // 设置表名

		// 自定义配置
		MySettings: model.MySettings{},
	}
	return modelSettings
}

// 用户-数据库表
type UserDatabaseModel struct {
	User_Id             *int64     `field_name:"user_id" field_type:"INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT" json:"user_id"`
	Database_Id         *int64     `field_name:"database_id" field_type:"INTEGER NOT NULL" json:"database_id"`
	Database_Table_Name *string    `field_name:"database_table_name" field_type:"TEXT NOT NULL" json:"database_table_name"`
	Status              *uint8     `field_name:"status" field_type:"INTEGER NOT NULL DEFAULT 1" json:"status"`
	Create_User_Id      *int64     `field_name:"create_user_id" field_type:"INTEGER" json:"create_user_id"`
	Create_Datetime     *time.Time `field_name:"create_datetime" field_type:"DATETIME NOT NULL" json:"create_datetime"`
}

func (UserDatabaseModel) ModelSet() *model.SQLite3Settings {
	modelSettings := &model.SQLite3Settings{
		MigrationsHandler: model.MigrationsHandler{ // 迁移时处理函数
			BeforeFunc: nil, // 迁移之前处理函数
			AfterFunc:  nil, // 迁移之后处理函数
		},

		TABLE_NAME: "user_database", // 设置表名

		// 自定义配置
		MySettings: model.MySettings{},
	}
	return modelSettings
}

// 用户-文件表
type UserFileModel struct {
	User_Id         *int64     `field_name:"user_id" field_type:"INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT" json:"user_id"`
	File_Id         *int64     `field_name:"file_id" field_type:"INTEGER NOT NULL" json:"file_id"`
	Status          *uint8     `field_name:"status" field_type:"INTEGER NOT NULL DEFAULT 1" json:"status"`
	Create_User_Id  *int64     `field_name:"create_user_id" field_type:"INTEGER" json:"create_user_id"`
	Create_Datetime *time.Time `field_name:"create_datetime" field_type:"DATETIME NOT NULL" json:"create_datetime"`
}

func (UserFileModel) ModelSet() *model.SQLite3Settings {
	modelSettings := &model.SQLite3Settings{
		MigrationsHandler: model.MigrationsHandler{ // 迁移时处理函数
			BeforeFunc: nil, // 迁移之前处理函数
			AfterFunc:  nil, // 迁移之后处理函数
		},

		TABLE_NAME: "user_file", // 设置表名

		// 自定义配置
		MySettings: model.MySettings{},
	}
	return modelSettings
}

// 用户-计划任务表
type UserCrontabModel struct {
	User_Id         *int64     `field_name:"user_id" field_type:"INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT" json:"user_id"`
	Crontab_Id      *int64     `field_name:"crontab_id" field_type:"INTEGER NOT NULL" json:"crontab_id"`
	Status          *uint8     `field_name:"status" field_type:"INTEGER NOT NULL DEFAULT 1" json:"status"`
	Create_User_Id  *int64     `field_name:"create_user_id" field_type:"INTEGER" json:"create_user_id"`
	Create_Datetime *time.Time `field_name:"create_datetime" field_type:"DATETIME NOT NULL" json:"create_datetime"`
}

func (UserCrontabModel) ModelSet() *model.SQLite3Settings {
	modelSettings := &model.SQLite3Settings{
		MigrationsHandler: model.MigrationsHandler{ // 迁移时处理函数
			BeforeFunc: nil, // 迁移之前处理函数
			AfterFunc:  nil, // 迁移之后处理函数
		},

		TABLE_NAME: "user_crontab", // 设置表名

		// 自定义配置
		MySettings: model.MySettings{},
	}
	return modelSettings
}
