package service

import(
	"rguldibi.com/SQLiteDemo/repository"
)

type UserService struct {
	userRepository repository.UserRepository
}

type UserBalanceService struct {
	userBalanceRepository repository.UserBalanceRepository
}

func NewUserService(userRepository repository.UserRepository) UserService{
	return UserService{userRepository:userRepository}
}

func NewUserBalanceService(userBalanceRepository repository.UserBalanceRepository) UserBalanceService{
	return UserBalanceService{userBalanceRepository:userBalanceRepository}
}

