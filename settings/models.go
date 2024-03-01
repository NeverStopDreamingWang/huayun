package settings

import (
	"encoding/base64"
	"github.com/NeverStopDreamingWang/goi/crypto"
	"github.com/NeverStopDreamingWang/goi/db"
	"github.com/NeverStopDreamingWang/goi/migrate"
	"github.com/NeverStopDreamingWang/goi/model"
	"time"
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
}

// 配置表
type SettingsModel struct {
	Name            *string    `field_name:"name" field_type:"TEXT NOT NULL UNIQUE" json:"name"`
	Value           *string    `field_name:"value" field_type:"TEXT NOT NULL" json:"value"`
	Type            *string    `field_name:"type" field_type:"TEXT NOT NULL" json:"type"`
	Desc            *string    `field_name:"desc" field_type:"TEXT" json:"desc"`
	Update_Datetime *time.Time `field_name:"update_datetime" field_type:"DATETIME" json:"update_datetime"`
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
		// 面板设置
		{"PanelName", "华运面板", "string", "面板名称"},
		{"Language", "zh", "string", "面板语言"},
		{"Theme", "auto", "string", "面板主题"},
		{"Version", "v0.1.0", "string", "面板版本"},
		// 面板服务
		{"BindAddress", "0.0.0.0", "string", "绑定地址"},
		{"BindDomain", "", "string", "绑定域名"},
		{"Port", "8080", "int", "面板端口"},
		{"Ipv6", "disable", "string", "Ipv6状态"},
		{"DefaultNetwork", "all", "string", "默认网络"},
		{"SSL", "disable", "string", "SSL状态"},
		{"SSLType", "self", "string", "SSL类型"},
		{"TimeZone", "Asia/Shanghai", "string", "面板时区"},
		// 其它设置
		{"SecretKey", secretKey, "string", "AES密钥"},
		{"PrivateKey", privateKey, "string", "RSA私钥"},
		{"PublicKey", publicKey, "string", "RSA公钥"},
		{"SystemIP", "", "string", "系统IP"},
		{"SessionTimeout", "86400", "int", "Session过期时间"},
		{"FileRecycleBin", "enable", "string", "文件回收站"},
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
