
package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"resulguldibi/golang-sqlite-sample/util"
)

func (handler UserBalanceHandler) HandleGetAllUserBalanceFunc(ctx *gin.Context) {

	defer func() {
		if err := recover(); err != nil {
			responseSatus := PrepareResponseStatusWithMessage(false, fmt.Sprint(err))
			ctx.JSON(http.StatusBadRequest, responseSatus)
		}
	}()

	users, err := handler.userBalanceService.GetAllUserBalances()
	util.CheckErr(err)
	ctx.JSON(http.StatusOK, users)
}