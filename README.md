# golang-sqlite-sample
simple web application using golang and sqlite.

- To create required tables on sqlite

CREATE TABLE "user" ( `name` TEXT NOT NULL, `age` INTEGER NOT NULL, `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE );
CREATE TABLE "user_balance" ( `user_id` INTEGER NOT NULL UNIQUE, `balance` INTEGER NOT NULL );

- To create new user
POST /user HTTP/1.1
Host: localhost:8080
Content-Type: application/json

{
  "id":1,
  "age":35,
  "name":"elon musk"
}

- To delete existing user

DELETE /user HTTP/1.1
Host: localhost:8080
Content-Type: application/json

{
  "id":1
}

-To update existing user

PUT /user HTTP/1.1
Host: localhost:8080
Content-Type: application/json

{
  "id":1,
  "age":36,
  "name":"elon musk aged"
}

- To get an existing user

GET /user?id=83 HTTP/1.1
Host: localhost:8080
Content-Type: application/json

- To get all users

GET /users HTTP/1.1
Host: localhost:8080
Content-Type: application/json

- To delete multiple user

DELETE /users HTTP/1.1
Host: localhost:8080
Content-Type: application/json

{
  "idList":[1,2]
}

- To send money to beneficiary (sqlx.Transaction.Commit example)

POST /send-money HTTP/1.1
Host: localhost:8080
Content-Type: application/json

{
  "senderId":67,
  "beneficiaryId":66,
  "amount":1
}

- To get all user's balance

GET /user-balances HTTP/1.1
Host: localhost:8080
Content-Type: application/json
Cache-Control: no-cache
Postman-Token: bab28f47-b402-8e25-5bfe-07f88a9ace4b


