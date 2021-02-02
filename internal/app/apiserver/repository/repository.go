package repository

type Repositories struct {
	ACMEAccounts ACMEAccountRepository
	ACMEServers ACMEServerRepository
}

func NewRepositories() *Repositories {
	return &Repositories{}
}
