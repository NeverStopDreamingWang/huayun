package log

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
			LogModel{},
		},
	}
	migrate.SQLite3Migrate(SQLite3Migrations)
}

// 日志表
type LogModel struct {
	Id              *int64     `field_name:"id" field_type:"INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT" json:"id"`
	User_Id         *int64     `field_name:"user_id" field_type:"INTEGER NOT NULL" json:"user_id"`
	Uri             *string    `field_name:"uri" field_type:"TEXT NOT NULL" json:"uri"`
	Type            *string    `field_name:"type" field_type:"TEXT NOT NULL" json:"type"`
	Status          *uint8     `field_name:"status" field_type:"INTEGER NOT NULL" json:"status"`
	Response        *string    `field_name:"response" field_type:"TEXT NOT NULL" json:"response"`
	Elapsed_Time    *int64     `field_name:"elapsed_time" field_type:"INTEGER NOT NULL" json:"elapsed_time"`
	Create_Datetime *time.Time `field_name:"create_datetime" field_type:"DATETIME NOT NULL" json:"create_datetime"`
}

func (LogModel) ModelSet() *model.SQLite3Settings {
	modelSettings := &model.SQLite3Settings{
		MigrationsHandler: model.MigrationsHandler{ // 迁移时处理函数
			BeforeFunc: nil, // 迁移之前处理函数
			AfterFunc:  nil, // 迁移之后处理函数
		},

		TABLE_NAME: "log", // 设置表名

		// 自定义配置
		MySettings: model.MySettings{},
	}
	return modelSettings
}
