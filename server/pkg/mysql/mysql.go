package mysql

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error

	var MYSQLDATABASE = os.Getenv("MYSQLDATABASE")
	var MYSQLPORT = os.Getenv("MYSQLPORT")
	var MYSQLHOST = os.Getenv("MYSQLHOST")
	var MYSQLPASSWORD = os.Getenv("MYSQLPASSWORD")
	var MYSQLUSER = os.Getenv("MYSQLUSER")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", MYSQLUSER, MYSQLPASSWORD, MYSQLHOST, MYSQLPORT, MYSQLDATABASE)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Connected to Database")
}
