package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"id.diengs.backend/internal/model"
	"id.diengs.backend/internal/usecase"
)

type AuthController struct {
	AuthUseCase *usecase.AuthUseCase
	Log         *logrus.Logger
}

func NewAuthController(authUseCase *usecase.AuthUseCase, log *logrus.Logger) *AuthController {
	return &AuthController{
		AuthUseCase: authUseCase,
		Log:         log,
	}
}

// Send OTP
func (c *AuthController) SendOtp(ctx *fiber.Ctx) error {
	request := new(model.AuthSendOtpReq)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parse request body")
		return fiber.ErrBadRequest
	}

	if err := c.AuthUseCase.SendOtp(ctx.UserContext(), request); err != nil {
		c.Log.Error(err)
		return err
	}

	return ctx.JSON(model.WebResponse[string]{
		Success: true,
		Message: "success send otp",
	})
}

// Verify OTP
func (c *AuthController) VeriftOtp(ctx *fiber.Ctx) error {
	request := new(model.AuthVerifyOtpRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parse request body")
		return fiber.ErrBadRequest
	}

	if err := c.AuthUseCase.VerifyOtp(ctx.UserContext(), request); err != nil {
		c.Log.Error(err)
		return err
	}

	return ctx.JSON(model.WebResponse[string]{
		Success: true,
		Message: "success verify otp",
	})
}

// Auth Google
func (c *AuthController) AuthGoogle(ctx *fiber.Ctx) error {
	userAgent := ctx.Get(fiber.HeaderUserAgent)
	ip := ctx.IP()

	request := new(model.AuthGoogleRequest)
	request.IP = ip
	request.UserAgent = userAgent

	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parse request body")
		return fiber.ErrBadRequest
	}

	response, err := c.AuthUseCase.AuthGoogle(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to login user")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.AuthResponse]{
		Success: true,
		Message: "success",
		Data:    response,
	})
}

// Logout
func (c *AuthController) Logout(ctx *fiber.Ctx) error {
	token := ctx.Query("token", "")
	if token == "" {
		return ctx.Status(403).JSON("invalid token")
	}

	err := c.AuthUseCase.Logout(ctx.UserContext(), token)
	if err != nil {
		c.Log.WithError(err).Error("failed to logout user")
		return err
	}

	return ctx.JSON(model.WebResponse[any]{
		Message: "logout succesfully",
		Success: true,
		Data:    nil,
	})
}

// Get Current
func (c *AuthController) Current(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*model.UserResponse)
	return ctx.JSON(model.WebResponse[*model.UserResponse]{
		Data: user,
	})
}
