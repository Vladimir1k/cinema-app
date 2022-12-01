package cinema

type User struct {
	Id       int    `json:"-"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Age      uint16 `json:"age"`
}
