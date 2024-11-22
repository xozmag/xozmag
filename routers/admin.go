package routers

func (r Router) AdminRouters() {
	adminGroup := r.router.Group("/api/v1/admin")
	adminGroup.POST("/xozmak", r.handler.CreateXozmak)
	adminGroup.GET("/xozmak", r.handler.GetXozmak)
	adminGroup.PUT("/xozmak/:id", r.handler.UpdateXozmak)
	adminGroup.DELETE("/xozmak/:id", r.handler.DeleteXozmak)
	adminGroup.POST("/category", r.handler.CreateCategory)
	adminGroup.GET("/category", r.handler.GetCategory)
	adminGroup.PUT("/category/:id", r.handler.UpdateCategory)
	adminGroup.DELETE("/category/:id", r.handler.DeleteCategory)
	adminGroup.POST("/subcategory", r.handler.CreateSubCategory)
	adminGroup.GET("/subcategory", r.handler.GetSubCategory)
	adminGroup.PUT("/subcategory/:id", r.handler.UpdateSubCategory)
	adminGroup.DELETE("/subcategory/:id", r.handler.DeleteSubCategory)

	adminGroup.POST("/multi-upload", r.handler.MultipleUploadHandler)
}
