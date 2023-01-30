package db

import (
	"encoding/hex"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gopartsrv/public/config"
	"gopartsrv/utils/encrypt"
	"time"
)

var (
	DBTeamMap map[string]*gorm.DB
	dbCfg     map[string]interface{}
)

func init() {
	DBTeamMap = make(map[string]*gorm.DB)
	getDbConfig()
}

//获取端口配置
func getConnConfig(instance string) (string, int, int, error) {
	temp := "%s:%s@tcp(%s:%s)/%s?charset=%s"
	instanceCfg := dbCfg[instance].(map[string]interface{})
	username := instanceCfg["username"]
	byt, err := hex.DecodeString(instanceCfg["password"].(string))
	password := encrypt.AesDecryptCFB(byt, []byte(config.GetKey()))
	if err != nil {
		return "", 0, 0, err
	}
	host := instanceCfg["host"]
	port := instanceCfg["port"]
	database := instanceCfg["database"]
	charset := instanceCfg["charset"]
	dsn := fmt.Sprintf(temp, username, password, host, port, database, charset)
	idleConn := instanceCfg["max_idle_conn"].(int64)
	openConn := instanceCfg["max_open_conn"].(int64)
	return dsn, int(idleConn), int(openConn), nil
}

//连接mysql
func newMysql(instance string) error {
	dsn, ic, oc, err := getConnConfig(instance)
	DBTeamMap[instance], err = gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}
	//DBTeamMap[instance].LogMode(true)
	DBTeamMap[instance].DB().SetMaxIdleConns(ic)
	DBTeamMap[instance].DB().SetMaxOpenConns(oc)
	DBTeamMap[instance].DB().Ping()
	DBTeamMap[instance].DB().SetConnMaxLifetime(10 * time.Minute)
	return nil
}

//获取db配置
func getDbConfig() {
	dbCfg = config.GetDBConfig()
	for k := range dbCfg {
		err := newMysql(k)
		if err != nil {
			panic(err)
		}
	}
}
