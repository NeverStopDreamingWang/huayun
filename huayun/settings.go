package huayun

import (
	"github.com/NeverStopDreamingWang/goi"
	"os"
	"path"
)

// Http 服务
var Server *goi.Engine

func init() {
	// 创建 http 服务
	Server = goi.NewHttpServer()

	// 项目路径
	Server.Settings.BASE_DIR, _ = os.Getwd()

	// 设置网络协议
	Server.Settings.NET_WORK = "tcp" // 默认 "tcp" 常用网络协议 "tcp"、"tcp4"、"tcp6"、"udp"、"udp4"、"udp6
	// 运行地址
	Server.Settings.BIND_ADDRESS = "0.0.0.0"
	// 端口
	Server.Settings.PORT = 8080

	// 数据库配置
	Server.Settings.DATABASES["default"] = goi.MetaDataBase{
		ENGINE:   "sqlite3",
		NAME:     path.Join(Server.Settings.BASE_DIR, "data/huayun.db"),
		USER:     "",
		PASSWORD: "",
		HOST:     "",
		PORT:     0,
	}

	// 设置 SSL
	Server.Settings.SSL = goi.MetaSSL{
		STATUS:    false, // SSL 开关
		CERT_PATH: path.Join(Server.Settings.BASE_DIR, "ssl/example.crt"),
		KEY_PATH:  path.Join(Server.Settings.BASE_DIR, "ssl/example.key"),
	}

	// 设置时区
	Server.Settings.TIME_ZONE = "Asia/Shanghai" // 默认 Asia/Shanghai
	// Server.Settings.TIME_ZONE = "America/New_York"

	// 设置最大缓存大小
	Server.Cache.EVICT_POLICY = goi.ALLKEYS_LRU    // 缓存淘汰策略
	Server.Cache.EXPIRATION_POLICY = goi.SCHEDULED // 过期策略
	Server.Cache.MAX_SIZE = 1024 * 1024 * 50       // 单位为字节，0 为不限制使用

	// 日志设置
	// 日志设置
	Server.Log.DEBUG = true
	// 日志列表
	defaultLog := newDefaultLog()
	accessLog := newAccessLog()
	errorLog := newErrorLog()
	Server.Log.LOGGERS = []*goi.MetaLogger{
		defaultLog, // 默认日志
		accessLog,  // 访问日志
		errorLog,   // 错误日志
	}

	// 设置自定义配置
	// redis配置
	// Server.Settings.Set("REDIS_HOST", "127.0.0.1")
	// Server.Settings.Set("REDIS_PORT", 6379)
	// Server.Settings.Set("REDIS_PASSWORD", "123")
	// Server.Settings.Set("REDIS_DB", 0)
}
