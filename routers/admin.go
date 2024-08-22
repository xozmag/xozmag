package routers

func (r Router) AdminRouters() {
	adminGroup := r.router.Group("/api/v1/admin")
	adminGroup.POST("/xozmak", r.handler.CreateXozmak)
	adminGroup.GET("/test", r.handler.Test)

	authGroup := r.router.Group("/api/auth")
	authGroup.POST("/sign-up", r.handler.SignUp)
	authGroup.POST("/login", r.handler.Login)
	authGroup.POST("/verifyCode", r.handler.VerifyCode)
}
