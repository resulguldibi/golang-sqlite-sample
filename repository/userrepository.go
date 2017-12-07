package repository

import(
	"rguldibi.com/SQLiteDemo/entity"
	"rguldibi.com/SQLiteDemo/util"
	"fmt"
	"strings"
)



func (userRepository UserRepository) SendMoney(senderId int64, beneficiaryId int64, amount int64) error {

	db := userRepository.GetConnection()
	params := make([]int64,0,2)
	params = append(params,senderId)
	params = append(params,beneficiaryId)

	query :="SELECT id,name,age FROM user where id in(?"+strings.Repeat(",?", len(params)-1)+")"
	users, err := userRepository.GetByIdList("User",query,senderId, beneficiaryId)

	util.CheckErr(err)

	if len(users) == 0 {
		return fmt.Errorf("invalid sender and beneficiary")
	}

	if len(users) == 1 {
		return fmt.Errorf("invalid sender or beneficiary")
	}

	userBalanceRepository := NewUserBalanceRepository(userRepository.dbClient)

	senderBalance, err := userBalanceRepository.GetByUserId(senderId)
	util.CheckErr(err)

	if senderBalance.(entity.UserBalance).Balance < amount {
		return fmt.Errorf("not enough balance")
	}

	beneficiaryBalance, err := userBalanceRepository.GetByUserId(beneficiaryId)
	util.CheckErr(err)

	beneficiaryNewBalance := beneficiaryBalance.(entity.UserBalance).Balance + amount
	senderNewBalance := senderBalance.(entity.UserBalance).Balance - amount

	tx, err := db.Begin()
	util.CheckErr(err)

	res, err := UpdateTransaction(tx, "update user_balance set balance=? where user_id=?", senderNewBalance, senderId)

	if err != nil {
		tx.Rollback()
		return err
	}

	if cnt, err := res.RowsAffected(); err != nil || cnt == 0 {
		tx.Rollback()
		return fmt.Errorf("system error")
	}

	res, err = UpdateTransaction(tx, "update user_balance set balance=? where user_id=?", beneficiaryNewBalance, beneficiaryId)

	if err != nil {
		tx.Rollback()
		return err
	}

	if cnt, err := res.RowsAffected(); err != nil || cnt == 0 {
		tx.Rollback()
		return fmt.Errorf("system error")
	}

	tx.Commit()

	return nil
}

func findUser(userList []interface{}, id int64) entity.User {
	var foundedUser entity.User
	if len(userList) > 0 {
		for _, item := range userList {
			if (item.(entity.User)).Id == id {
				foundedUser = item.(entity.User)
				break
			}
		}
	}
	return foundedUser
}