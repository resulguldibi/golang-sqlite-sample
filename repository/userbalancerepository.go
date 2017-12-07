package repository

import(
	"rguldibi.com/golang-sqlite-sample/entity"
)

func (userBalanceRepository UserBalanceRepository) GetByUserId(userId int64) (interface{}, error) {

	db := userBalanceRepository.GetConnection()

	result, err := Query(db, "SELECT user_id,balance FROM user_balance where user_id =?", userId)
	if err != nil {
		return nil, err
	}

	userBalance := entity.UserBalance{}

	convertMapToStruct(result[0], &userBalance)

	return userBalance, err
}