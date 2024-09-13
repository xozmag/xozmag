package routers

func (r Router) AdminRouters() {
	adminGroup := r.router.Group("/api/v1/admin")
	adminGroup.POST("/xozmak", r.handler.CreateXozmak)
	adminGroup.GET("/xozmak", r.handler.GetXozmak)
	adminGroup.PUT("/xozmak/:id", r.handler.UpdateXozmak)
	adminGroup.DELETE("/xozmak/:id", r.handler.DeleteXozmak)
}
