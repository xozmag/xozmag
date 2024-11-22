package routers

func (r Router) UploadRouters() {
	filesGroup := r.router.Group("/api")
	filesGroup.POST("/multi-upload", r.handler.MultipleUploadHandler)
}
