package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"gopartsrv/public/config"
	"time"
)


var (
	SqliteTeamMap map[string]*gorm.DB
	sqliteCfg     map[string]interface{}
)

func init() {
	SqliteTeamMap = make(map[string]*gorm.DB)
	//getSqliteConfig()
}

//获取端口配置
func getConnsConfig(instance string) (string, int, int, error) {
	instanceCfg := sqliteCfg[instance].(map[string]interface{})
	idleConn := instanceCfg["max_idle_conn"].(int64)
	openConn := instanceCfg["max_open_conn"].(int64)
	return instanceCfg["address"].(string), int(idleConn), int(openConn), nil
}

func newSqlite(instance string) error {
	address, ic, oc, err := getConnsConfig(instance)
	SqliteTeamMap[instance],err = gorm.Open("sqlite3", address)
	//SqliteTeamMap[instance],err = gorm.Open("sqlite3", "../../../gowork/public/config/piclist.db")
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}
	SqliteTeamMap[instance].DB().SetMaxIdleConns(ic)
	SqliteTeamMap[instance].DB().SetMaxOpenConns(oc)
	SqliteTeamMap[instance].DB().Ping()
	SqliteTeamMap[instance].DB().SetConnMaxLifetime(10 * time.Minute)
	return nil
}

//获取Sqlite配置
func getSqliteConfig() {
	sqliteCfg = config.GetSqliteConfig()
	for k := range sqliteCfg {
		err := newSqlite(k)
		if err != nil {
			panic(err)
		}
	}
}

