package repo

type postgresql struct{ tc SQLTxConn }

type PostgreSQL interface {
	// UsersRepository
	// GroupsRepositry
	// AccountsRepository
	// LoansRepository
	// TransactionsRepository
	// MambuRegistrationLogsRepository
	// MambuTransactionLogsRepository
	// VirtualAccountsRepository
	// BanksRepository
}

func NewPostgreSQL(txc SQLTxConn) PostgreSQL { return &postgresql{txc} }
