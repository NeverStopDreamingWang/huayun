package huayun

import (
	"github.com/NeverStopDreamingWang/huayun/huayun"
	// 注册 app
	_ "github.com/NeverStopDreamingWang/huayun/crontab"          // 定时任务
	_ "github.com/NeverStopDreamingWang/huayun/database/mongodb" // 数据库-MongoDB
	_ "github.com/NeverStopDreamingWang/huayun/database/mysql"   // MySQL
	_ "github.com/NeverStopDreamingWang/huayun/database/redis"   // SQLite3
	_ "github.com/NeverStopDreamingWang/huayun/database/sqlite3" // SQLite3
	_ "github.com/NeverStopDreamingWang/huayun/docker"           // Docker
	_ "github.com/NeverStopDreamingWang/huayun/file"             // 文件
	_ "github.com/NeverStopDreamingWang/huayun/firewall"         // 防火墙
	_ "github.com/NeverStopDreamingWang/huayun/home"             // 首页
	_ "github.com/NeverStopDreamingWang/huayun/log"              // 日志
	_ "github.com/NeverStopDreamingWang/huayun/settings"         // 设置
	_ "github.com/NeverStopDreamingWang/huayun/site/java"        // 网站-java
	_ "github.com/NeverStopDreamingWang/huayun/site/php"         // 网站-php
	_ "github.com/NeverStopDreamingWang/huayun/site/python"      // 网站-python
	_ "github.com/NeverStopDreamingWang/huayun/software"         // 软件
	_ "github.com/NeverStopDreamingWang/huayun/terminal"         // 终端
	_ "github.com/NeverStopDreamingWang/huayun/user/role"        // 用户-角色
	_ "github.com/NeverStopDreamingWang/huayun/user/user"        // 用户-用户
)

func Start() {
	// 启动服务
	huayun.Server.RunServer()
}

func Stop() {
	err := huayun.Server.StopServer()
	panic(err)
}
