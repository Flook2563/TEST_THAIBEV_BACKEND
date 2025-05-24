package common

import "github.com/labstack/echo/v4"

func GetAndValidateRequestBody(c echo.Context, request interface{}) error {
	if err := BindValidateBody(c, request); err != nil {
		return err
	}
	return nil
}

func BindValidateBody(c echo.Context, request interface{}) error {
	if err := c.Bind(request); err != nil {
		return err
	}
	if err := c.Validate(request); err != nil {
		return err
	}
	return nil
}
