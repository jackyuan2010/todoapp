package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/cors"
	todogorm "github.com/jackyuan2010/todoapp/server/gorm"
	todogormpg "github.com/jackyuan2010/todoapp/server/gorm/postgres"
	model "github.com/jackyuan2010/todoapp/server/model"
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
	// initDB()
	router := SetupRoutes()
	router.Run(":8081")
}

func initDB() {
	dbconfig := todogorm.DbConfig{
		Host:     "172.17.0.2",
		Username: "gormuser", Password: "gormuser",
		DbName: "todo_db", Port: "5432",
		Config: "sslmode=disable TimeZone=Asia/Shanghai",
	}

	pgdbconfig := todogormpg.Converte2PostgresDbConfig(&dbconfig)

	var dbdsn todogorm.DbDsn = *pgdbconfig

	var dbcontext todogorm.DbContext = todogormpg.PostgresDbContext{}

	var repository todogorm.Repository = todogormpg.NewPostgresRepository(&dbdsn)

	db := dbcontext.GetDb(&dbdsn)

	sqlDB, _ := db.DB()

	defer sqlDB.Close()

	// sqlDB.SetMaxIdleConns(5)
	// sqlDB.SetMaxOpenConns(100)

	// db.AutoMigrate(&model.ToDoTask{})

	// object := model.ToDoTask{Task_Description: "to do task 1", Is_Finished: false, Is_Delay: false}
	// object.Id = "1"

	// db.Create(&object)

	// object := model.ToDoTask{BaseModel: model.BaseModel{ Id: "2" }, Task_Description: "to do task 2", Is_Finished: false, Is_Delay: false}

	// db.Create(&object)


	dbObj := repository.QueryById("2")

	fmt.Println(dbObj)
}