
package factory

import(
	"resulguldibi/golang-sqlite-sample/entity"
)

var factoryList = make(map[string]IFactory)

func InitFactoryList(){
	factoryList["User"] = UserFactory{}
	factoryList["UserBalance"] = UserBalanceFactory{}
}

type IFactory interface{
	GetInstance() entity.IEntity
}

type UserFactory struct{

}

type UserBalanceFactory struct{
	
}

func (userFactory UserFactory) GetInstance() entity.IEntity{
	return &entity.User{}
}

func (userBalanceFactory UserBalanceFactory) GetInstance() entity.IEntity{
	return &entity.UserBalance{}
}

func GetEntityInstance(name string)entity.IEntity{
	return factoryList[name].GetInstance()
}