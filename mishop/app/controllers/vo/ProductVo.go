package vo

type ProductVo struct {
	loginVo LoginVo
}

type LoginVo struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
