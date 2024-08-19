package routers

func (r Router) AdminRouters() {
	adminGroup := r.router.Group("/api/v1/admin")
	adminGroup.GET("/test", r.handler.Test)
}