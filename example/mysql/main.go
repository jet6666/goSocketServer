package main

import (
	"fmt"
	//_ "github.com/go-sql-driver/mysql"
	"log"
	"gorm.io/driver/mysql"
	"time"

	//go get -u no success ? "xorm.io/xorm"
	//expired "github.com/go-xorm/xorm"
	"gorm.io/gorm"
)

type Name2  struct {
Id    int  `gorm:"primaryKey;AUTO_INCREMENT"`
Name  string
//TestField string
}

type Name3  struct {
	Id    int  `gorm:"primaryKey;AUTO_INCREMENT"`
	Name  string
	//TestField string
}
func main() {
	log.Println("mysql example ")

	dsn := "root:@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database ")
	}

	//auto create table
	 //db.AutoMigrate(&Name3{})


	//v :=  time.Unix( time.Now().Unix() ,0 )
	currentTime := time.Now().UTC() //time.Now()
	v2 := currentTime.Format("2006-01-02 15:04:05") //currentTime.Format ("YYYY-MM-DD hh:mm:ss"  )
	fmt.Println(v2 )
	user := Name2{
		//Id:   2,
		Name : "xxx" + v2 ,
		//TestField:"yyyyy",
	}
	result :=db.Create(&user)
	if result.Error ==nil {
		log.Println("create success !")
	}else {
		log.Println("create failed ",result.Error.Error() )
	}
	 log.Println("db.create " ,result.Error ,result.RowsAffected )



	var r Result
	db.Raw("SELECT * FROM name2   WHERE id =99  ").Scan(&r)

	fmt.Println( "result = " ,r.Name ,r.Id)

	var r2 Result
	rows,err :=db.Raw("SELECT * FROM name2 ORDER BY id DESC ").Rows()
	defer rows.Close()
	for rows.Next() {
		 rows.Scan(&r2 )
		db.ScanRows(rows, &r2 )
		fmt.Println( "result2  = " ,r2.Name ,r2.Id)
	}






}

type Result struct {
	Id int
	Name string
}
