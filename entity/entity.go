package entity


type ResponseStatus struct {
	IsSucccess bool   `json:"issuccess"`
	Message    string `json:"message"`
}

type User struct {
	Id   int64  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Age  int    `db:"age" json:"age"`	
}

type UserBalance struct {
	UserId  int64 `db:"user_id" json:"userId"`
	Balance int64 `db:"balance" json:"balance"`
}

type IEntity interface{
	Do()
}


func (user User)Do(){}

func (userBalance UserBalance)Do(){}



