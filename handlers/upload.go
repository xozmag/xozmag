package handlers

import (
	"delivery/pkg/http"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *Handler) MultipleUploadHandler(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		h.handleResponse(c, http.BadRequest, "error in FileUpload")
		return
	}

	files := form.File["files"]

	fmt.Println("files: ", files)
	for _, file := range files{
		fmt.Println("file name: ", file.Filename)
		fmt.Println("file header: ", file.Header)
		fmt.Println("file size: ", file.Size)
	}

}
