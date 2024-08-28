package routers

func(r Router) UserRouters() {
	authGroup := r.router.Group("/api/auth")
	authGroup.POST("/sendcode", r.handler.SendCode)
	authGroup.POST("/login", r.handler.Login)
	authGroup.POST("/registr", r.handler.SignUp)
	
	authGroup.PUT("/users", r.handler.UpdateProfile)
	

}