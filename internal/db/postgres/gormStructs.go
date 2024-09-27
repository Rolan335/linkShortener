package postgres

type Clients struct {
	ID       uint
	Login    string `json:"login" gorm:"unique"`
	Password string `json:"password"`
}

/*
{
    "login": "test_client",
    "password": "Pa$$w0rd"
}
*/

type Links struct {
	ID          uint
	Original    string `json:"original"`
	Short       string `json:"short"`
	SearchCount uint   `json:"searchCount"`
	ClientsID   uint   `json:"client_id"`
	Clients     Clients
}
