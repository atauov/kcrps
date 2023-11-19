package dashboard

type User struct {
	Id          int    `json:"-" db:"id"`
	CompanyName string `json:"company-name"`
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Role        int    `json:"role" binding:"required"`
}
