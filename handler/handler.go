package handler

import (
	"resulguldibi/golang-sqlite-sample/service"
)

type UserHandler struct{
	userService	service.UserService
}

type UserBalanceHandler struct{
	userBalanceService	service.UserBalanceService
}

func NewUserHandler(userService service.UserService) UserHandler{
	return UserHandler{userService : userService}
}

func NewUserBalanceHandler(userBalanceService service.UserBalanceService) UserBalanceHandler{
	return UserBalanceHandler{userBalanceService : userBalanceService}
}
