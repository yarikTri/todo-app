package todo

// User declares context for active user
type User struct {
	ID       int    `json:"-"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
