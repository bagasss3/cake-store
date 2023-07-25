package constant

import (
	"cake-store/src/util"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// http errors
var (
	ErrInvalidArgument = echo.NewHTTPError(http.StatusBadRequest, "invalid argument")
	ErrAlreadyDeleted  = echo.NewHTTPError(http.StatusBadRequest, "record already deleted")
	ErrNotFound        = echo.NewHTTPError(http.StatusNotFound, "record not found")
	ErrInternal        = echo.NewHTTPError(http.StatusInternalServerError, "internal system error")
	ErrFieldEmpty      = echo.NewHTTPError(http.StatusBadRequest, "requirement field empty")
)

// httpValidationOrInternalErr return valdiation or internal error
func HttpValidationOrInternalErr(err error) error {
	switch t := err.(type) {
	case validator.ValidationErrors:
		_ = t
		errVal := err.(validator.ValidationErrors)

		fields := map[string]interface{}{}
		for _, ve := range errVal {
			fields[ve.Field()] = fmt.Sprintf("Failed on the '%s' tag", ve.Tag())
		}

		return echo.NewHTTPError(http.StatusBadRequest, util.Dump(fields))
	default:
		return ErrInternal
	}
}
