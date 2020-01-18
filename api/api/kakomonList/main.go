package main

import (
	_ "api/kakomonList/routers"

	"github.com/astaxie/beego"

)

func init() {
	user := beego.AppConfig.String("mysqluser")
    pass := beego.AppConfig.String("mysqlpass")
    host := beego.AppConfig.String("mysqlurls")
    db := beego.AppConfig.String("mysqldb")
    datasource := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8", user, pass, host, db)
    orm.RegisterDataBase("default", "mysql", datasource, 30)

    err := orm.RunSyncdb("default", false, true)
    if err != nil {
        panic(err)
    }


	// param 1: driverName
	// param 2: database type
	// This mapping driverName and database type
	// mysql / sqlite3 / postgres registered by default already
    orm.RegisterDriver("mysql", orm.DRMySQL)

    // param 1:        Database alias. ORM will use it to switch database.
    // param 2:        driverName
    // param 3:        connection string
    orm.RegisterDataBase("default", "mysql", "root:root@/orm_test?charset=utf8")
}

func main() {
    o := orm.NewOrm()
    o.Using("default") // Using default, you can use other database

    profile := new(Profile)
    profile.Age = 30

    user := new(User)
    user.Profile = profile
    user.Name = "slene"

    fmt.Println(o.Insert(profile))
    fmt.Println(o.Insert(user))

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
