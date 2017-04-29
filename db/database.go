package db


import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
	"fmt"
	"golang.org/x/crypto/bcrypt"

)

func InitDb(){
	db := Database()
	defer db.Close()
	//db.AutoMigrate(&api.User{})
	//db.AutoMigrate(&api.Post{})
	//db.AutoMigrate(&api.Comment{})
	db.DropTableIfExists(&User{})
	db.DropTableIfExists(&Post{})
	db.DropTableIfExists(&Comment{})
	db.CreateTable(&User{})
	db.CreateTable(&Post{})
	db.CreateTable(&Comment{})

	Seed()

}

func Seed( ) {
	db := Database()
	defer db.Close()
	hashedPw, _ :=bcrypt.GenerateFromPassword([]byte("password"),bcrypt.DefaultCost)
	user:= User{
		Email:"admin@q.q",
		Name: "admin",
		Password: string(hashedPw),
	}
	db.Create(&user)

	post := Post{
		Message:"first post",
		User:user,
	}
	db.Create(&post)
	comment := Comment{
		Message:"firs comment",
		User:user,
		Post:post,
	}
	db.Create(&comment)

	user2:= User{
		Email:"adminq@q.q",
		Name: "admin",
		Password: string(hashedPw),
	}
	db.Create(&user2)
}

func Database() *gorm.DB {
	db, err := gorm.Open("mysql", "root:adminmysqlpassword@tcp(192.168.0.23:3306)/vcelin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("Error connecting to DB: <%s> \n", err)

		panic(fmt.Errorf("failed to connect database with error  <%s> \n",err))
	}
	return db



}
