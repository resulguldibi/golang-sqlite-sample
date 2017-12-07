package server

import (
	"github.com/gin-gonic/gin"
	"resulguldibi/golang-sqlite-sample/handler"
	"resulguldibi/golang-sqlite-sample/repository"
	"resulguldibi/golang-sqlite-sample/service"
	"resulguldibi/golang-sqlite-sample/factory"
)

func NewServer() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)
	server := gin.New()
	factory.InitFactoryList()
	AddDefaultMiddlewaresToEngine(server)
	//TODO : get connection info from config
	dbClientFactory := repository.NewDbClientFactory("sqlite3","./SQLiteDB.db")

	server.POST("/user", func(ctx *gin.Context) {		
		dbClient := dbClientFactory.NewDBClient()
		userHandler := handler.NewUserHandler(service.NewUserService(repository.NewUserRepository(dbClient)))
		userHandler.HandleCreateUserFunc(ctx)
	})

	server.DELETE("/user", func(ctx *gin.Context) {
		dbClient := dbClientFactory.NewDBClient()
		userHandler := handler.NewUserHandler(service.NewUserService(repository.NewUserRepository(dbClient)))
		userHandler.HandleDeleteUserFunc(ctx)
	})

	server.DELETE("/users", func(ctx *gin.Context) {
		dbClient := dbClientFactory.NewDBClient()
		userHandler := handler.NewUserHandler(service.NewUserService(repository.NewUserRepository(dbClient)))
		userHandler.HandleDeleteMultipleUserFunc(ctx)
	})

	server.PUT("/user", func(ctx *gin.Context) {
		dbClient := dbClientFactory.NewDBClient()
		userHandler := handler.NewUserHandler(service.NewUserService(repository.NewUserRepository(dbClient)))
		userHandler.HandleUpdateUserFunc(ctx)
	})

	server.GET("/user", func(ctx *gin.Context) {
		dbClient := dbClientFactory.NewDBClient()
		userHandler := handler.NewUserHandler(service.NewUserService(repository.NewUserRepository(dbClient)))
		userHandler.HandleGetUserFunc(ctx)
	})

	server.GET("/users", func(ctx *gin.Context) {
		dbClient := dbClientFactory.NewDBClient()
		userHandler := handler.NewUserHandler(service.NewUserService(repository.NewUserRepository(dbClient)))
		userHandler.HandleGetAllUserFunc(ctx)
	})

	server.POST("/send-money", func(ctx *gin.Context) {
		dbClient := dbClientFactory.NewDBClient()
		userHandler := handler.NewUserHandler(service.NewUserService(repository.NewUserRepository(dbClient)))
		userHandler.HandleSendMoneyFunc(ctx)
	})


	server.GET("/user-balances", func(ctx *gin.Context) {
		dbClient := dbClientFactory.NewDBClient()
		userBalanceHandler := handler.NewUserBalanceHandler(service.NewUserBalanceService(repository.NewUserBalanceRepository(dbClient)))
		userBalanceHandler.HandleGetAllUserBalanceFunc(ctx)
	})

	return server
}

func AddDefaultMiddlewaresToEngine(server *gin.Engine){
	//engine.Use(secure.Secure(secure.Options))
	server.Use(gin.Logger())
	server.Use(gin.Recovery())
}
