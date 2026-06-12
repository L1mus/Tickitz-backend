package router

import (
	"github.com/L1mus/Tickitz-backend/internal/controller"
	"github.com/L1mus/Tickitz-backend/internal/middleware"
	"github.com/L1mus/Tickitz-backend/internal/repository"
	"github.com/L1mus/Tickitz-backend/internal/service"
	"github.com/L1mus/Tickitz-backend/pkg"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func AuthRouter(router *gin.RouterGroup, db *pgxpool.Pool, rdb *redis.Client, mailer pkg.Mailer) {
	authRouter := router.Group("/auth")

	authRepository := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepository, rdb, mailer)
	// authService := service.NewAuthService(authRepository, rdb)
	authController := controller.NewAuthController(authService)

	authRouter.POST("", authController.Login)
	authRouter.POST("/register", authController.Register)
	authRouter.POST("/register/activate", authController.Activate)
	authRouter.POST("/register/resend-otp", authController.ResendOTP)
	authRouter.POST("/check-email", authController.ForgotPassword)
	authRouter.POST("/check-email/verify-otp", authController.VerifyResetOTP)
	authRouter.POST("/check-email/verify-otp/reset", authController.ResetPassword)
	authRouter.DELETE("/logout", middleware.VerifyToken, middleware.CheckBlacklist(rdb), authController.Logout)
}
