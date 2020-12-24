package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gin_test/pkg/app"
	"gin_test/pkg/logging"
	"gin_test/pkg/upload"
)

// @Summary 选择文件
// @Produce  json
// @Param image formData file true "Image File"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /upload/selectFile [get]
func SelectFile(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.C.HTML(http.StatusOK, "upload.html", gin.H{
		"title": "Main website",
	})
}

// @Summary 上传图片
// @Produce  json
// @Param image formData file true "Image File"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/tags/import [post]
func UploadImage(c *gin.Context) {
	appG := app.Gin{C: c}
	file, image, err := c.Request.FormFile("image")
	if err != nil {
		logging.LogWarning(err)
		appG.Response(http.StatusInternalServerError, e.ERROR, "", nil, false)
		return
	}

	if image == nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, "", nil, false)
		return
	}

	imageName := upload.GetImageName(image.Filename)
	fullPath := upload.GetImageFullPath()
	savePath := upload.GetImagePath()
	src := fullPath + imageName

	if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
		appG.Response(http.StatusBadRequest, e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT, "", nil, false)
		return
	}

	err = upload.CheckImage(fullPath)
	if err != nil {
		logging.LogWarning(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_UPLOAD_CHECK_IMAGE_FAIL, "", nil, false)
		return
	}

	if err := c.SaveUploadedFile(image, src); err != nil {
		logging.LogWarning(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_UPLOAD_SAVE_IMAGE_FAIL, "", nil, false)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, "", map[string]string{
		"image_url":      upload.GetImageFullUrl(imageName),
		"image_save_url": savePath + imageName,
	}, false)
}
