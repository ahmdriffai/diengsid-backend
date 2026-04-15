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
		Message: "success send otp",
	})
}
