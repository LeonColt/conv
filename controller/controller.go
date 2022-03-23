package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/leoncolt/conv/models"
)

type Controller struct{}

func (controller *Controller) OkBlob(ctx echo.Context, contentType string, b []byte) error {
	if err := ctx.Blob(http.StatusOK, contentType, b); err != nil {
		log.Warn("error serving JSON")
		return err
	}
	return nil
}

func (controller *Controller) Ok(ctx echo.Context, data interface{}) error {
	if err := ctx.JSON(http.StatusOK, data); err != nil {
		log.Warn("error serving JSON")
		return err
	}
	return nil
}

func (controller *Controller) Created(ctx echo.Context, data interface{}) error {
	if err := ctx.JSON(http.StatusCreated, data); err != nil {
		log.Warn("error serving JSON")
		return err
	}
	return nil
}

func (controller *Controller) NoContent(ctx echo.Context) error {
	if err := ctx.NoContent(204); err != nil {
		log.Warn("error serving JSON")
		return err
	}
	return nil
}

func (controller *Controller) BadRequestError(ctx echo.Context, err error) error {
	if err := ctx.JSON(http.StatusBadRequest, models.HTTPError{
		Code:    400,
		Message: err.Error(),
	}); err != nil {
		log.Warn("error serving JSON")
		return err
	}
	return nil
}

func (controller *Controller) Unauthorized(ctx echo.Context) error {
	if err := ctx.JSON(http.StatusBadRequest, models.HTTPError{
		Code:    http.StatusUnauthorized,
		Message: "please sign in first",
	}); err != nil {
		log.Warn("error serving JSON")
		return err
	}
	return nil
}

func (controller *Controller) UnauthorizedError(ctx echo.Context, err error) error {
	if err := ctx.JSON(http.StatusBadRequest, models.HTTPError{
		Code:    http.StatusUnauthorized,
		Message: err.Error(),
	}); err != nil {
		log.Warn("error serving JSON")
		return err
	}
	return nil
}

func (controller *Controller) Forbidden(ctx echo.Context) error {
	if err := ctx.JSON(http.StatusBadRequest, models.HTTPError{
		Code:    http.StatusUnauthorized,
		Message: "you are not allowed to do this",
	}); err != nil {
		log.Warn("error serving JSON")
		return err
	}
	return nil
}

func (controller *Controller) ForbiddenError(ctx echo.Context, err error) error {
	if err := ctx.JSON(http.StatusBadRequest, models.HTTPError{
		Code:    http.StatusUnauthorized,
		Message: err.Error(),
	}); err != nil {
		log.Warn("error serving JSON")
		return err
	}
	return nil
}

func (controller *Controller) NotFoundError(ctx echo.Context, err error) error {
	if err := ctx.JSON(http.StatusNotFound, models.HTTPError{
		Code:    404,
		Message: err.Error(),
	}); err != nil {
		log.Warn("error serving JSON")
		return err
	}
	return nil
}

func (controller *Controller) InternalServerError(ctx echo.Context, err error) error {
	log.Warn(err)
	if err := ctx.JSON(http.StatusInternalServerError, models.HTTPError{
		Code:    500,
		Message: "Internal Server Error",
	}); err != nil {
		log.Warn("error serving JSON")
		return err
	}
	return nil
}

func (controller *Controller) ServiceUnavailableError(ctx echo.Context, err error) error {
	log.Warn(err)
	if err := ctx.JSON(http.StatusServiceUnavailable, models.HTTPError{
		Code:    503,
		Message: "Service Unavailable",
	}); err != nil {
		log.Warn("error serving JSON")
		return err
	}
	return nil
}
