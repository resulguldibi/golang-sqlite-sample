package contract

type SendMoneyRequest struct {
	SenderId      int64 `db:"senderId"`
	BeneficiaryId int64 `db:"beneficiaryId"`
	Amount        int64 `db:"amount"`
}