package api

type User struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Credentials struct {
	ID       int32  `json:"id"`
	UserID   int32  `json:"user_id"`
	Service  string `json:"service"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Vault struct {
	UserID      int32
	Credentials []Credentials
}

type MasterKey struct {
	UserID int32
	Key    []byte
}
