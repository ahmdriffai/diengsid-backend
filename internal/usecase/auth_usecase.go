package usecase

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"id.diengs.backend/internal/entity"
	"id.diengs.backend/internal/model"
	"id.diengs.backend/internal/pkg"
	"id.diengs.backend/internal/pkg/mailview"
	"id.diengs.backend/internal/repository"
)

type AuthUseCase struct {
	DB           *gorm.DB
	Log          *logrus.Logger
	Validate     *validator.Validate
	Mail         *pkg.Mail
	UserRepo     *repository.UserRepo
	EmailOtpRepo *repository.EmailOtpRepo
}

func NewAuthUseCase(
	db *gorm.DB,
	log *logrus.Logger,
	validate *validator.Validate,
	mail *pkg.Mail,
	userRepo *repository.UserRepo,
	emailOtpRepo *repository.EmailOtpRepo,
) *AuthUseCase {
	return &AuthUseCase{
		DB:           db,
		UserRepo:     userRepo,
		Log:          log,
		Validate:     validate,
		Mail:         mail,
		EmailOtpRepo: emailOtpRepo,
	}
}

// Send Email OTP
func (u *AuthUseCase) SendOtp(ctx context.Context, request *model.AuthSendOtpReq) error {
	// transaction
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	// validate request
	if err := u.Validate.Struct(request); err != nil {
		u.Log.WithError(err).Error("FAILED TO VALIDATE REQUEST.")
		return fiber.ErrBadRequest
	}

	// create otp
	otp := rand.Intn(900000) + 100000 // 6 digit
	otpString := fmt.Sprint(otp)
	hashedOtp, _ := bcrypt.GenerateFromPassword([]byte(otpString), bcrypt.DefaultCost)

	// expired code
	expiredCode := time.Now().Add(time.Duration(5) * time.Minute).UnixMilli()

	// create to db
	emailOtp := &entity.EmailOtp{
		Email:     request.Email,
		OtpCode:   string(hashedOtp),
		ExpiredAt: expiredCode,
	}

	if err := u.EmailOtpRepo.Create(tx, emailOtp); err != nil {
		u.Log.WithError(err).Error("FAILED TO CREATE DB.")
		return fiber.ErrInternalServerError
	}

	// send mail
	bodyEmail := mailview.RegisterOtpMailView(otpString)
	err := u.Mail.SendMail([]string{request.Email}, "Kode pengamanan anda "+otpString, bodyEmail)
	if err != nil {
		u.Log.WithError(err).Error("FAILED TO SEND EMAIL.")
		return fiber.ErrInternalServerError
	}

	// commit transaction
	if err := tx.Commit().Error; err != nil {
		u.Log.WithError(err).Error("FAILED TO COMMIT TRANSACTION.")
		return fiber.ErrInternalServerError
	}

	return nil
}

// Verify Email OTP
func (u *AuthUseCase) VerifyOtp(ctx context.Context, requet *model.AuthVerifyOtpRequest) error {
	// transaction
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	// find OTP
	emailOtp := new(entity.EmailOtp)
	err := u.EmailOtpRepo.FindActiveAndEmail(tx, emailOtp, requet.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return fiber.ErrNotFound
		}
		u.Log.WithError(err).Error("FAILED TO FIND EMAIL OTP.")
		return fiber.ErrInternalServerError
	}

	if emailOtp.AttemptCount >= emailOtp.MaxAttempt {
		return fiber.NewError(fiber.StatusTooManyRequests, "TO MANY ATTEMPS.")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(emailOtp.OtpCode), []byte(requet.Otp)); err != nil {
		u.Log.WithError(err).Error("FAILED TO FIND EMAIL OTP.")
		return fiber.ErrNotFound
	}

	// Update OTP
	emailOtp.IsUsed = true
	if err := u.EmailOtpRepo.Update(tx, emailOtp); err != nil {
		u.Log.WithError(err).Error("FAILED TO UPDATE EMAIL OTP.")
		return fiber.ErrInternalServerError
	}

	// commit transaction
	if err := tx.Commit().Error; err != nil {
		u.Log.WithError(err).Error("FAILED TO COMMIT TRANSACTION.")
		return fiber.ErrInternalServerError
	}

	return nil
}
