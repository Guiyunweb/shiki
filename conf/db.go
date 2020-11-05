package conf

import (
	"log"
	"os"
)

// 数据库配置
type DbConf struct {
	DbType     string
	DefaultDsn string
}

func LoadDatabase() *DbConf {
	DefaultDsn := Conf.GetString("DB_DEFAULT_DSN")
	if DefaultDsn == "" {
		log.Println("数据库配置不能为空")
		os.Exit(-1)
	}
	DbType := Conf.GetString("DB_TYPE")
	if DbType == "" {
		log.Println("数据库配置不能为空")
		os.Exit(-1)
	}
	return &DbConf{
		DbType,
		DefaultDsn,
	}
}
