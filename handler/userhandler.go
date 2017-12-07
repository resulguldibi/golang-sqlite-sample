package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"resulguldibi/golang-sqlite-sample/contract"
	"resulguldibi/golang-sqlite-sample/entity"
	"resulguldibi/golang-sqlite-sample/util"
	"strconv"
)

func (handler UserHandler) HandleCreateUserFunc(ctx *gin.Context) {

	defer func() {
		if err := recover(); err != nil {
			responseSatus := PrepareResponseStatusWithMessage(false, fmt.Sprint(err))
			ctx.JSON(http.StatusBadRequest, responseSatus)
		}
	}()

	var user entity.User
	if err := ctx.ShouldBindJSON(&user); err == nil {
		err = handler.userService.CreateUser(user)
		util.CheckErr(err)
		responseSatus := PrepareResponseStatusWithMessage(true, "user created successfully")
		ctx.JSON(http.StatusOK, responseSatus)
	} else {
		responseSatus := PrepareResponseStatusWithMessage(false, err.Error())
		ctx.JSON(http.StatusBadRequest, responseSatus)
	}
}

func (handler UserHandler) HandleDeleteUserFunc(ctx *gin.Context) {

	defer func() {
		if err := recover(); err != nil {

			responseSatus := PrepareResponseStatusWithMessage(false, fmt.Sprint(err))
			ctx.JSON(http.StatusBadRequest, responseSatus)
		}
	}()

	var userToDelete struct {
		Id int64 `json:"id"`
	}

	if err := ctx.ShouldBindJSON(&userToDelete); err == nil {

		affectedRows, err := handler.userService.DeleteUser(userToDelete.Id)
		util.CheckErr(err)

		responseSatus := PrepareResponseStatusWithMessage(true, "user deleted successfully")

		if affectedRows == 0 {
			responseSatus = PrepareResponseStatusWithMessage(false, fmt.Sprintf("no user deleted with id: %d", userToDelete.Id))
			ctx.JSON(http.StatusBadRequest, responseSatus)
		} else {
			ctx.JSON(http.StatusOK, responseSatus)
		}

	} else {
		responseSatus := PrepareResponseStatusWithMessage(false, err.Error())
		ctx.JSON(http.StatusBadRequest, responseSatus)
	}
}

func (handler UserHandler) HandleDeleteMultipleUserFunc(ctx *gin.Context) {

	defer func() {
		if err := recover(); err != nil {
			responseSatus := PrepareResponseStatusWithMessage(false, fmt.Sprint(err))
			ctx.JSON(http.StatusBadRequest, responseSatus)
		}
	}()

	var usersToDelete struct {
		IdList []int64 `json:"idList"`
	}

	if err := ctx.ShouldBindJSON(&usersToDelete); err == nil {
		affectedRows, err := handler.userService.DeleteMultipleUser(usersToDelete.IdList)
		util.CheckErr(err)
		responseSatus := PrepareResponseStatusWithMessage(true, "user deleted successfully")

		if affectedRows == 0 {
			responseSatus = PrepareResponseStatusWithMessage(false, fmt.Sprintf("no user deleted with id: %d", usersToDelete.IdList))
			ctx.JSON(http.StatusBadRequest, responseSatus)
		} else {
			ctx.JSON(http.StatusOK, responseSatus)
		}
	} else {
		responseSatus := PrepareResponseStatusWithMessage(false, err.Error())
		ctx.JSON(http.StatusBadRequest, responseSatus)
	}

}

func (handler UserHandler) HandleUpdateUserFunc(ctx *gin.Context) {

	defer func() {
		if err := recover(); err != nil {
			responseSatus := PrepareResponseStatusWithMessage(false, fmt.Sprint(err))
			ctx.JSON(http.StatusBadRequest, responseSatus)
		}
	}()

	var user entity.User

	if err := ctx.ShouldBindJSON(&user); err == nil {
		affectedRows, err := handler.userService.UpdateUser(user)
		util.CheckErr(err)
		responseSatus := PrepareResponseStatusWithMessage(true, "user updated successfully")

		if affectedRows == 0 {
			responseSatus = PrepareResponseStatusWithMessage(false, fmt.Sprintf("no user updated with id: %d", user.Id))
			ctx.JSON(http.StatusBadRequest, responseSatus)
		} else {
			ctx.JSON(http.StatusOK, responseSatus)
		}
	} else {
		responseSatus := PrepareResponseStatusWithMessage(false, err.Error())
		ctx.JSON(http.StatusBadRequest, responseSatus)
	}
}

func (handler UserHandler) HandleGetUserFunc(ctx *gin.Context) {

	defer func() {
		if err := recover(); err != nil {
			responseSatus := PrepareResponseStatusWithMessage(false, fmt.Sprint(err))
			ctx.JSON(http.StatusBadRequest, responseSatus)
		}
	}()

	id, err := strconv.ParseInt(ctx.Request.URL.Query().Get("id"), 10, 64)
	if err != nil {
		panic(err)
	}

	user, err := handler.userService.GetUser(id)
	ctx.JSON(http.StatusOK, user)
}

func (handler UserHandler) HandleSendMoneyFunc(ctx *gin.Context) {

	defer func() {
		if err := recover(); err != nil {
			responseSatus := PrepareResponseStatusWithMessage(false, fmt.Sprint(err))
			ctx.JSON(http.StatusBadRequest, responseSatus)
		}
	}()

	var sendMoneyRequest contract.SendMoneyRequest

	if err := ctx.ShouldBindJSON(&sendMoneyRequest); err == nil {
		err = handler.userService.SendMoney(sendMoneyRequest)
		util.CheckErr(err)
		responseSatus := PrepareResponseStatusWithMessage(true, "money sended successfully")
		ctx.JSON(http.StatusOK, responseSatus)
	} else {
		responseSatus := PrepareResponseStatusWithMessage(false, err.Error())
		ctx.JSON(http.StatusBadRequest, responseSatus)
	}
}

func (handler UserHandler) HandleGetAllUserFunc(ctx *gin.Context) {

	defer func() {
		if err := recover(); err != nil {
			responseSatus := PrepareResponseStatusWithMessage(false, fmt.Sprint(err))
			ctx.JSON(http.StatusBadRequest, responseSatus)
		}
	}()

	users, err := handler.userService.GetAllUsers()
	util.CheckErr(err)
	ctx.JSON(http.StatusOK, users)
}

func PrepareResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func PrepareResponseStatus(err interface{}) entity.ResponseStatus {
	return entity.ResponseStatus{
		IsSucccess: false,
		Message:    fmt.Sprint(err),
	}
}

func PrepareResponseStatusWithMessage(isSucccess bool, message string) entity.ResponseStatus {
	return entity.ResponseStatus{
		IsSucccess: isSucccess,
		Message:    message,
	}
}
