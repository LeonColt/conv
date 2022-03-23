package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/leoncolt/conv/service"
)

// Operations about Conv
type ConvController struct {
	Controller
	service *service.ConvService
}

func (controller *ConvController) Register(group *echo.Group) {
	controller.service = new(service.ConvService)
	controller.service.Init()
	group.POST("/docx-to-pdf", controller.DocxToPdf)
	group.POST("/docx-to-html", controller.DocxToHtml)
	group.POST("/html-to-docx", controller.HtmlToDocx)
	group.POST("/html-to-pdf", controller.HtmlToPdf)
}

// DocxToPdf godoc
// Post DocxToPdf
// @ID DocxToPdf
// @Summary  Convert file from docx to pdf
// @Description Convert file from docx to pdf
// @Tags DocxToPdf
// @Accept mpfd
// @Product octet-stream
// @Param file formData file true "this is a test file"
// @Success 200
// @Failure 400 {object} models.HTTPError
// @Failure 500 {object} models.HTTPError
// @Router /docx-to-pdf [post]
func (controller *ConvController) DocxToPdf(ctx echo.Context) error {
	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		return controller.InternalServerError(ctx, err)
	}
	file, err := fileHeader.Open()
	if err != nil {
		return controller.InternalServerError(ctx, err)
	}
	defer file.Close()

	res, err := controller.service.DocxToPdf(file)
	if err != nil {
		return controller.InternalServerError(ctx, err)
	}
	return controller.OkBlob(ctx, "application/pdf", res)
}

// DocxToHtml godoc
// Post DocxToHtml
// @ID DocxToHtml
// @Summary  Convert file from docx to html
// @Description Convert file from docx to html
// @Tags DocxToHtml
// @Accept mpfd
// @Product octet-stream
// @Param file formData file true "this is a test file"
// @Success 200
// @Failure 400 {object} models.HTTPError
// @Failure 500 {object} models.HTTPError
// @Router /docx-to-html [post]
func (controller *ConvController) DocxToHtml(ctx echo.Context) error {
	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		return controller.InternalServerError(ctx, err)
	}
	file, err := fileHeader.Open()
	if err != nil {
		return controller.InternalServerError(ctx, err)
	}
	defer file.Close()

	res, err := controller.service.DocxToHtml(file)
	if err != nil {
		return controller.InternalServerError(ctx, err)
	}
	return controller.OkBlob(ctx, "text/html", res)
}

// HtmlToDocx godoc
// Post HtmlToDocx
// @ID HtmlToDocx
// @Summary  Convert file from html to docx
// @Description Convert file from html to docx
// @Tags HtmlToDocx
// @Accept mpfd
// @Product octet-stream
// @Param file formData file true "this is a test file"
// @Success 200
// @Failure 400 {object} models.HTTPError
// @Failure 500 {object} models.HTTPError
// @Router /html-to-docx [post]
func (controller *ConvController) HtmlToDocx(ctx echo.Context) error {
	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		return controller.InternalServerError(ctx, err)
	}
	file, err := fileHeader.Open()
	if err != nil {
		return controller.InternalServerError(ctx, err)
	}
	defer file.Close()

	res, err := controller.service.HtmlToDocx(file)
	if err != nil {
		return controller.InternalServerError(ctx, err)
	}
	return controller.OkBlob(ctx, "application/vnd.openxmlformats-officedocument.wordprocessingml.document", res)
}

// HtmlToPdf godoc
// Post HtmlToPdf
// @ID HtmlToPdf
// @Summary  Convert file from html to pdf
// @Description Convert file from html to pdf
// @Tags HtmlToPdf
// @Accept mpfd
// @Product octet-stream
// @Param file formData file true "this is a test file"
// @Success 200
// @Failure 400 {object} models.HTTPError
// @Failure 500 {object} models.HTTPError
// @Router /html-to-pdf [post]
func (controller *ConvController) HtmlToPdf(ctx echo.Context) error {
	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		return controller.InternalServerError(ctx, err)
	}
	file, err := fileHeader.Open()
	if err != nil {
		return controller.InternalServerError(ctx, err)
	}
	defer file.Close()

	res, err := controller.service.HtmlToPdf(file)
	if err != nil {
		return controller.InternalServerError(ctx, err)
	}
	return controller.OkBlob(ctx, "application/pdf", res)
}
