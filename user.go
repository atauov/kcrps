package dashboard

type User struct {
	Id          int    `json:"-" db:"id"`
	CompanyName string `json:"company_name" binding:"required"`
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Role        int    `json:"role"`
}
