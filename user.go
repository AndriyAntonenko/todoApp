package todo

type User struct {
	Id int `json:"id"`

	// Make next fields required
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
