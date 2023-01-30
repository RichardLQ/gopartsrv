package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var cfg map[string]interface{}

const key = "DrvlKOAbcxbxqMKb"

func init() {
	viper.SetDefault("sys", map[string]string{"env": "pro"})
	viper.SetConfigName("config")
	viper.AddConfigPath("../config")
	viper.AddConfigPath("../public/config")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("./public/config")
	viper.AddConfigPath(".")
	viper.SetConfigType("toml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	viper.WatchConfig()
	cfg = make(map[string]interface{})
	cfg = viper.AllSettings()
}

//获取配置
func GetConfig() map[string]interface{} {
	return cfg
}

//获取db配置
func GetDBConfig() map[string]interface{} {
	config := cfg["database"].(map[string]interface{})
	env := (cfg["env"]).(string)
	dbCfg := config[env].(map[string]interface{})
	return dbCfg
}

//获取sqlite配置
func GetSqliteConfig() map[string]interface{} {
	config := cfg["sqlite"].(map[string]interface{})
	env := (cfg["env"]).(string)
	dbCfg := config[env].(map[string]interface{})
	return dbCfg
}

//获取邮箱配置
func GetEmailConfig() map[string]interface{} {
	config := cfg["email"].(map[string]interface{})
	env := (cfg["env"]).(string)
	dbCfg := config[env].(map[string]interface{})
	return dbCfg
}

//获取redis配置
func GetRedisConfig() map[string]interface{} {
	config := cfg["redis"].(map[string]interface{})
	env := (cfg["env"]).(string)
	redisCfg := config[env].(map[string]interface{})
	return redisCfg
}

//加密key
func GetKey() string {
	return key
}

//获取环境
func GetEnvConfig() string {
	return (cfg["env"]).(string)
}

//返回类型
func GetIsType() int64 {
	return 2
}
