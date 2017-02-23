package libs

import (
  "fmt"

  "github.com/revel/revel"
  _ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/robfig/config"
)

var (
	Engine *xorm.Engine
)


func InitDB() {
	c, err := config.ReadDefault(revel.BasePath + "/conf/database.conf")

	driver, _ := c.String(revel.RunMode, "db.driver")
	dbname, _ := c.String(revel.RunMode, "db.dbname")
	user, _ := c.String(revel.RunMode, "db.user")
	password, _ := c.String(revel.RunMode, "db.password")
	host, _ := c.String(revel.RunMode, "db.host")

	params := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", user, password, host, dbname)

	Engine, err = xorm.NewEngine(driver, params)
	// defer Engine.Close()
	if err != nil {
		panic(err)
	}
	// Engine.ShowSQL = revel.DevMode
}
