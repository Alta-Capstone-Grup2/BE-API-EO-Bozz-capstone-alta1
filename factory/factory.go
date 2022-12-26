package factory

import (
	authDelivery "capstone-alta1/features/auth/delivery"
	authRepo "capstone-alta1/features/auth/repository"
	authService "capstone-alta1/features/auth/service"

	userDelivery "capstone-alta1/features/user/delivery"
	userRepo "capstone-alta1/features/user/repository"
	userService "capstone-alta1/features/user/service"

	clientDelivery "capstone-alta1/features/client/delivery"
	clientRepo "capstone-alta1/features/client/repository"
	clientService "capstone-alta1/features/client/service"

	partnerDelivery "capstone-alta1/features/partner/delivery"
	partnerRepo "capstone-alta1/features/partner/repository"
	partnerService "capstone-alta1/features/partner/service"

	reviewDelivery "capstone-alta1/features/review/delivery"
	reviewRepo "capstone-alta1/features/review/repository"
	reviewService "capstone-alta1/features/review/service"

	additionalDelivery "capstone-alta1/features/additional/delivery"
	additionalRepo "capstone-alta1/features/additional/repository"
	additionalService "capstone-alta1/features/additional/service"

	cityDelivery "capstone-alta1/features/city/delivery"
	cityRepo "capstone-alta1/features/city/repository"
	cityService "capstone-alta1/features/city/service"

	// serviceDelivery "capstone-alta1/features/service/delivery"
	// serviceRepo "capstone-alta1/features/service/repository"
	// serviceService "capstone-alta1/features/service/service"

	// orderDelivery "capstone-alta1/features/order/delivery"
	// orderRepo "capstone-alta1/features/order/repository"
	// orderService "capstone-alta1/features/order/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	authRepoFactory := authRepo.New(db)
	authServiceFactory := authService.New(authRepoFactory)
	authDelivery.New(authServiceFactory, e)

	userRepoFactory := userRepo.New(db)
	userServiceFactory := userService.New(userRepoFactory)
	userDelivery.New(userServiceFactory, e)

	clientRepoFactory := clientRepo.New(db)
	clientServiceFactory := clientService.New(clientRepoFactory)
	clientDelivery.New(clientServiceFactory, e)

	partnerRepoFactory := partnerRepo.New(db)
	partnerServiceFactory := partnerService.New(partnerRepoFactory)
	partnerDelivery.New(partnerServiceFactory, e)

	reviewRepoFactory := reviewRepo.New(db)
	reviewServiceFactory := reviewService.New(reviewRepoFactory)
	reviewDelivery.New(reviewServiceFactory, e)

	additionalRepoFactory := additionalRepo.New(db)
	additionalServiceFactory := additionalService.New(additionalRepoFactory)
	additionalDelivery.New(additionalServiceFactory, e)

	cityRepoFactory := cityRepo.New(db)
	cityServiceFactory := cityService.New(cityRepoFactory)
	cityDelivery.New(cityServiceFactory, e)

	// serviceRepoFactory := serviceRepo.New(db)
	// serviceServiceFactory := serviceService.New(serviceRepoFactory)
	// serviceDelivery.New(serviceServiceFactory, e)

	// orderRepoFactory := orderRepo.New(db)
	// orderServiceFactory := orderService.New(orderRepoFactory)
	// orderDelivery.New(orderServiceFactory, e)
}
