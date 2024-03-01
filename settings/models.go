package settings

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"github.com/NeverStopDreamingWang/goi"
	"github.com/NeverStopDreamingWang/goi/crypto"
	"github.com/NeverStopDreamingWang/goi/db"
	"github.com/NeverStopDreamingWang/goi/migrate"
	"github.com/NeverStopDreamingWang/goi/model"
	"github.com/NeverStopDreamingWang/huayun/huayun"
	"strconv"
)

func init() {
	// sqlite 数据库
	SQLite3Migrations := model.SQLite3MakeMigrations{
		DATABASES: []string{"default"},
		MODELS: []model.SQLite3Model{
			SettingsModel{},
		},
	}
	migrate.SQLite3Migrate(SQLite3Migrations)

	err := ReadSettings()
	if err != nil {
		panic(err)
	}

	// var (
	// 	name      = "test1"
	// 	value     = "test1_value1111"
	// 	valueType = "string"
	// 	desc      = "测试"
	// )
	//
	// NewSettings := &SettingsModel{
	// 	Name:            &name,
	// 	Value:           &value,
	// 	Type:            &valueType,
	// 	Desc:            &desc,
	// 	Update_Datetime: nil,
	// }
	// result, err := NewSettings.WriteSettings()
	// if err != nil {
	// 	goi.Log.Error(result, err)
	// } else {
	// 	goi.Log.Info(result)
	// }
}

// 配置表
type SettingsModel struct {
	Name            *string `field_name:"name" field_type:"TEXT NOT NULL UNIQUE" json:"name"`
	Value           *string `field_name:"value" field_type:"TEXT NOT NULL" json:"value"`
	Type            *string `field_name:"type" field_type:"TEXT NOT NULL" json:"type"`
	Desc            *string `field_name:"desc" field_type:"TEXT" json:"desc"`
	Update_Datetime *string `field_name:"update_datetime" field_type:"DATETIME" json:"update_datetime"`
}

func (settingsModel SettingsModel) ModelSet() *model.SQLite3Settings {
	modelSettings := &model.SQLite3Settings{
		MigrationsHandler: model.MigrationsHandler{ // 迁移时处理函数
			BeforeFunc: nil,          // 迁移之前处理函数
			AfterFunc:  InItSettings, // 迁移之后处理函数
		},

		TABLE_NAME: "settings", // 设置表名

		// 自定义配置
		MySettings: model.MySettings{},
	}
	return modelSettings
}

func InItSettings() error {
	SQLite3DB, err := db.SQLite3Connect("default")
	defer SQLite3DB.Close()
	if err != nil {
		return err
	}

	// 生成 AES 密钥
	secretKeyBytes := make([]byte, 32)
	err = crypto.GenerateAES(secretKeyBytes)
	if err != nil {
		return err
	}
	// 将随机字节转换为Base64编码的字符串
	secretKey := base64.StdEncoding.EncodeToString(secretKeyBytes)

	// 生成 RSA 密钥
	var privateKeyBytes, publicKeyBytes []byte
	err = crypto.GenerateRSA(2048, &privateKeyBytes, &publicKeyBytes)
	if err != nil {
		return err
	}
	privateKey := string(privateKeyBytes)
	publicKey := string(publicKeyBytes)

	configData := [][]string{
		// 面板信息
		{"PanelName", "华运面板", "string", "面板名称"},
		{"Language", "zh", "string", "面板语言"},
		{"Theme", "auto", "string", "面板主题"},
		{"Version", "v0.1.0", "string", "面板版本"},
		// 面板服务
		{"NetWork", "tcp", "string", "网络协议"},
		{"BindAddress", "0.0.0.0", "string", "绑定地址"},
		{"Port", "8080", "int", "面板端口"},
		{"Domain", "", "string", "绑定域名"},
		{"SSL", "false", "bool", "SSL状态"},
		{"SSLType", "self", "string", "SSL类型"},
		{"CertPath", "", "string", "SSL证书"},
		{"KeyPath", "", "string", "SSL密钥"},
		{"TimeZone", "Asia/Shanghai", "string", "面板时区"},
		{"EvictPolicy", "2", "int", "缓存淘汰策略"},
		{"ExpirationPolicy", "1", "int", "过期策略"},
		{"MaxSize", "52428800", "int", "缓存大小"},
		{"Debug", "true", "bool", "开发模式"},
		// 其它设置
		{"SecretKey", secretKey, "string", "AES密钥"},
		{"PrivateKey", privateKey, "string", "RSA私钥"},
		{"PublicKey", publicKey, "string", "RSA公钥"},
		{"SystemIP", "", "string", "系统IP"},
		{"SessionTimeout", "86400", "int", "Session过期时间"},
		{"FileRecycleBin", "true", "bool", "文件回收站"},
	}

	for _, item := range configData {
		user := &SettingsModel{
			Name:            &item[0],
			Value:           &item[1],
			Type:            &item[2],
			Desc:            &item[3],
			Update_Datetime: nil,
		}
		SQLite3DB.SetModel(SettingsModel{})
		_, err = SQLite3DB.Insert(user)
		if err != nil {
			return err
		}
	}
	return nil
}

// 从数据库中加载配置
func (settings SettingsModel) WriteSettings() (sql.Result, error) {
	// 连接数据库
	sqlite3DB, err := db.SQLite3Connect("default")
	if err != nil {
		panic(fmt.Sprintf("连接 SQLite3 [default] 数据库 错误: %v", err))
	}

	sqlite3DB.SetModel(SettingsModel{})

	settingsTemp := SettingsModel{}
	err = sqlite3DB.Where("name=?", settings.Name).First(&settingsTemp)

	if settingsTemp.Name == nil {
		return sqlite3DB.Insert(settings)
	}
	return sqlite3DB.Where("name=?", settings.Name).Update(settings)
}

// 从数据库中加载配置
func ReadSettings(nameList ...string) error {
	// 连接数据库
	sqlite3DB, err := db.SQLite3Connect("default")
	if err != nil {
		return err
	}

	settingsList := make([]SettingsModel, 0)
	sqlite3DB.SetModel(SettingsModel{})
	if len(nameList) != 0 {
		// err = sqlite3DB.Where("name in (?)", "'"+strings.Join(names, "','")+"'").Select(&settingsList)
		settingsList = make([]SettingsModel, len(nameList))
		for i, name := range nameList {
			data := SettingsModel{}
			err = sqlite3DB.Where("name=?", name).First(&data)
			if err != nil {
				return err
			}
			settingsList[i] = data
		}
	} else {
		err = sqlite3DB.Select(&settingsList)
		if err != nil {
			return err
		}
	}

	// 从数据库中加载配置
	for _, item := range settingsList {
		switch *item.Name {
		case "NetWork":
			huayun.Server.Settings.NET_WORK = *item.Value
		case "BindAddress":
			huayun.Server.Settings.BIND_ADDRESS = *item.Value
		case "Port":
			Port, _ := strconv.Atoi(*item.Value)
			huayun.Server.Settings.PORT = uint16(Port)
		case "Domain":
			huayun.Server.Settings.Domain = *item.Value
		case "SSL":
			huayun.Server.Settings.SSL = goi.MetaSSL{
				STATUS:    *item.Value == "true", // SSL 开关
				TYPE:      "",
				CERT_PATH: "",
				KEY_PATH:  "",
			}
		case "SSLType":
			huayun.Server.Settings.SSL.TYPE = *item.Value
		case "CertPath":
			huayun.Server.Settings.SSL.CERT_PATH = *item.Value
		case "KeyPath":
			huayun.Server.Settings.SSL.KEY_PATH = *item.Value
		case "TimeZone":
			huayun.Server.Settings.TIME_ZONE = *item.Value
		case "EvictPolicy":
			EvictPolicy, _ := strconv.Atoi(*item.Value)
			huayun.Server.Cache.EVICT_POLICY = goi.EvictPolicy(EvictPolicy)
		case "ExpirationPolicy":
			ExpirationPolicy, _ := strconv.Atoi(*item.Value)
			huayun.Server.Cache.EXPIRATION_POLICY = goi.ExpirationPolicy(ExpirationPolicy)
		case "MaxSize":
			MaxSize, _ := strconv.Atoi(*item.Value)
			huayun.Server.Cache.MAX_SIZE = int64(MaxSize)
		case "Debug":
			huayun.Server.Log.DEBUG = *item.Value == "true"
		case "SecretKey":
			huayun.Server.Settings.SECRET_KEY = *item.Value
		case "PrivateKey":
			huayun.Server.Settings.PRIVATE_KEY = *item.Value
		case "PublicKey":
			huayun.Server.Settings.PUBLIC_KEY = *item.Value
		default:
			huayun.Server.Settings.Set(*item.Name, *item.Value)
		}
	}
	return nil
}
