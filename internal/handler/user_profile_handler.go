package handler

import (
	"net/http"
	"thaibev_backend/internal/common"
	"thaibev_backend/internal/domain"

	"github.com/labstack/echo/v4"
)

func (h *handler) CreateUserProfile(c echo.Context) error {
	ctx := c.Request().Context()

	req := domain.CreateUserProfileRequest{}
	if err := common.GetAndValidateRequestBody(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	resp, err := h.services.CreateUserProfile(ctx, &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create user profile",
		})
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *handler) GetUserProfile(c echo.Context) error {
	ctx := c.Request().Context()

	req := domain.UserProfileRequest{}
	if err := common.GetAndValidateRequestBody(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	resp, err := h.services.GetUserProfile(ctx, &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to retrieve user profile",
		})
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *handler) CheckEmailExists(c echo.Context) error {
	ctx := c.Request().Context()
	email := c.Param("email")
	if email == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "email is required"})
	}

	exists, err := h.services.CheckEmailExists(ctx, email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal error"})
	}

	return c.JSON(http.StatusOK, map[string]bool{"exists": exists})
}
