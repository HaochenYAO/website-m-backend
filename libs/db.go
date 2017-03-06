package libs

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" // go_mysql库
	"github.com/go-xorm/xorm"
	"github.com/revel/revel"
	"github.com/robfig/config"
)

// DbEngine 数据库query engine
var DbEngine *xorm.Engine

// InitDB 初始化数据库engine
func InitDB() {
	c, err := config.ReadDefault(revel.BasePath + "/conf/database.conf")
	if err != nil {
		panic(err)
	}
	driver, _ := c.String(revel.RunMode, "db.driver")
	dbname, _ := c.String(revel.RunMode, "db.dbname")
	user, _ := c.String(revel.RunMode, "db.user")
	password, _ := c.String(revel.RunMode, "db.password")
	host, _ := c.String(revel.RunMode, "db.host")

	params := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", user, password, host, dbname)

	DbEngine, err = xorm.NewEngine(driver, params)
	// defer DbEngine.Close()
	if err != nil {
		panic(err)
	}
	// DbEngine.ShowSQL = revel.DevMode
}
