package factory

import (
	authDelivery "capstone-alta1/features/auth/delivery"
	authRepo "capstone-alta1/features/auth/repository"
	authService "capstone-alta1/features/auth/service"

	userDelivery "capstone-alta1/features/user/delivery"
	userRepo "capstone-alta1/features/user/repository"
	userService "capstone-alta1/features/user/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	userRepoFactory := userRepo.New(db)
	userServiceFactory := userService.New(userRepoFactory)
	userDelivery.New(userServiceFactory, e)

	authRepoFactory := authRepo.New(db)
	authServiceFactory := authService.New(authRepoFactory)
	authDelivery.New(authServiceFactory, e)
}
