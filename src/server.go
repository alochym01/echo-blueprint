package main

import (
	"controllers/user"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"helpers"
	"models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Init database connection instance
	// DbConfig := config.GetConfig()
	// gormParameters := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable", DbConfig.DbHost, DbConfig.DbPort, DbConfig.DbName, DbConfig.DbUsername, DbConfig.DbPassword)
	// gormDB, err := gorm.Open("mysql", gormParameters)
	// fmt.Println(gormParameters)
	// fmt.Println(DbConfig)
	// fmt.Println(gormDB)
	gormDBLite, err := gorm.Open("sqlite3", "alochym.db")
	// fmt.Println(gormDBLite)
	if err != nil {
		fmt.Println(err)
	}

	helpers.DbGorm = gormDBLite // setting db connection
	// fmt.Println(helpers.DbGorm)
	helpers.DbGorm.LogMode(true)                            // Enable Logger, show detailed log
	helpers.DbGorm.SetLogger(log.New(os.Stdout, "\r\n", 0)) //Using os.Stdout as the backend

	// auto create table if the table is not exist in database
	helpers.DbGorm.AutoMigrate(&models.User{})

	// create user demo
	helpers.DbGorm.Create(&models.User{
		Name:     "Do Nguyen Ha",
		Email:    "hadn4@fpt.com.vn",
		Password: "alochym",
	})

	e := echo.New()             // create Echo instance
	e.Use(middleware.Recover()) //https://echo.labstack.com/middleware/recover
	e.Use(middleware.Logger())  // log
	e.Use(middleware.CORS())    // CORS from Any Origin, Any Method

	// routing for user controller
	e.GET("/users", user.Index)
	e.POST("/users", user.Create)
	e.GET("/users/:id", user.Show)
	e.PUT("/users/:id", user.Update)
	e.DELETE("/users/:id", user.Delete)

	// dump all routes to route.json
	route, _ := json.MarshalIndent(e.Routes(), "", "  ")
	ioutil.WriteFile("routes.json", route, 0644)

	defer helpers.DbGorm.Close()     // close database connection
	e.Logger.Fatal(e.Start(":8000")) // starting Echo instance
}
