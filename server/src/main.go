package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/cors"
	todogorm "github.com/jackyuan2010/todoapp/server/src/gorm"
	todogormpg "github.com/jackyuan2010/todoapp/server/src/gorm/postgres"
	model "github.com/jackyuan2010/todoapp/server/src/model"
)

func index(c *gin.Context) {
	c.Header("Access-Control-Allow", "*")
	c.JSON(http.StatusOK, gin.H{
			"message": "TODO App",
		})
}

func SetupRoutes() *gin.Engine {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/", index)

	return router
}

func main() {
	fmt.Println("to do app starting....")
	initDB()
	router := SetupRoutes()
	router.Run(":8081")
}

func initDB() {
	dbconfig := gpaasgorm.DbConfig{
		Host:     "172.17.0.2",
		Username: "gormuser", Password: "gormuser",
		DbName: "todo_db", Port: "5432",
		Config: "sslmode=disable TimeZone=Asia/Shanghai",
	}

	pgdbconfig := gtodogormpg.Converte2PostgresDbConfig(&dbconfig)

	var dbdsn todogorm.DbDsn = *pgdbconfig

	println("database source name:", dbdsn.Dsn())

	var dbcontext todogorm.DbContext = todogormpg.PostgresDbContext{}

	var repository todogorm.Repository = todogormpg.NewPostgresRepository(&dbdsn)

	fmt.Println("Repository method start......")


	fmt.Println("Repository method end......")

	db := dbcontext.GetDb(&dbdsn)

	sqlDB, _ := db.DB()

	defer sqlDB.Close()

	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(100)

	db.AutoMigrate(&model.ToDoTask{})

	object := model.ToDoTask{id: "1", task_description: "to do task 1", is_finished: false, is_delay: false}

	db.Create(&object)
}