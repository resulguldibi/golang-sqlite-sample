package service

import(
	"rguldibi.com/golang-sqlite-sample/entity"
	"rguldibi.com/golang-sqlite-sample/util"
)

func (service UserBalanceService) GetAllUserBalances() ([]*entity.UserBalance, error) {
	defer service.userBalanceRepository.CloseConnection()
	userBalances, err := service.userBalanceRepository.GetAll("UserBalance", "select user_id,balance from user_balance")
	util.CheckErr(err)

	list := make([]*entity.UserBalance, 0, len(userBalances.([]interface{})))
	for _, item := range userBalances.([]interface{}) {
		list = append(list, item.(*entity.UserBalance))
	}
	return list, err
}