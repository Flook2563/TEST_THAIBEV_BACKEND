package handler

import (
	"net/http"
	"thaibev_backend/internal/common"
	"thaibev_backend/internal/domain"

	"github.com/labstack/echo/v4"
)

func (h *handler) GetAllUser(c echo.Context) error {
	ctx := c.Request().Context()
	users, err := h.services.GetAllUser(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get users"})
	}
	return c.JSON(http.StatusOK, users)
}

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

func (h *handler) DeleteUserProfile(c echo.Context) error {
	ctx := c.Request().Context()
	userID := c.Param("user_id")
	if userID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "user_id is required"})
	}
	err := h.services.DeleteUserProfile(ctx, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "deleted successfully"})
}
