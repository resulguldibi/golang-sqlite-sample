package service

import(
	"strings"
	"resulguldibi/golang-sqlite-sample/entity"
	"resulguldibi/golang-sqlite-sample/contract"
	"resulguldibi/golang-sqlite-sample/util"
)

func (service UserService) CreateUser(user entity.User) error {
	
	defer service.userRepository.CloseConnection()
	_, err := service.userRepository.Insert("insert into user(name, age) values(?,?)", user.Name, user.Age)
	util.CheckErr(err)
	return nil
}

func (service UserService) DeleteUser(userId int64) (int64, error) {
	defer service.userRepository.CloseConnection()
	res, err := service.userRepository.Delete("delete from user where id=?", userId)
	util.CheckErr(err)
	affectedRows, err := res.RowsAffected()
	util.CheckErr(err)
	return affectedRows, nil
}

func (service UserService) DeleteMultipleUser(userIdList []int64) (int64, error) {
	
	defer service.userRepository.CloseConnection()

	data := getInnnerStatementParameters(len(userIdList))

	new := toInterfaceArray(userIdList)

	res, err := service.userRepository.Delete("delete from user where id in("+data+")", new...)
	util.CheckErr(err)
	affectedRows, err := res.RowsAffected()
	util.CheckErr(err)
	return affectedRows, nil
}

func getInnnerStatementParameters(length int) string {
	list := make([]string, 0, length)
	for i := 0; i < length; i++ {
		list = append(list, "?")
	}
	return strings.Join(list, ",")
}

func toInterfaceArray(list []int64) []interface{} {
	interfaceList := make([]interface{}, len(list))
	for i, v := range list {
		interfaceList[i] = v
	}
	return interfaceList
}

func (service UserService) UpdateUser(user entity.User) (int64, error) {
	defer service.userRepository.CloseConnection()
	res, err := service.userRepository.Update("update user set name=?, age=? where id=?", user.Name, user.Age, user.Id)
	util.CheckErr(err)
	affectedRows, err := res.RowsAffected()
	util.CheckErr(err)
	return affectedRows, nil
}

func (service UserService) GetUser(userId int64) (*entity.User, error) {
	defer service.userRepository.CloseConnection()
	user, err := service.userRepository.GetById("User",userId,"SELECT id,name,age FROM user where id =?")
	util.CheckErr(err)
	return user.(*entity.User), nil
}

func (service UserService) SendMoney(sendMoneyRequest contract.SendMoneyRequest) error {
	defer service.userRepository.CloseConnection()
	err := service.userRepository.SendMoney(sendMoneyRequest.SenderId, sendMoneyRequest.BeneficiaryId, sendMoneyRequest.Amount)
	util.CheckErr(err)
	return err
}

func (service UserService) GetAllUsers() ([]*entity.User, error) {
	defer service.userRepository.CloseConnection()
	users, err := service.userRepository.GetAll("User", "select id,name,age from user")
	util.CheckErr(err)

	list := make([]*entity.User, 0, len(users.([]interface{})))
	for _, item := range users.([]interface{}) {
		list = append(list, item.(*entity.User))
	}
	return list, err
}