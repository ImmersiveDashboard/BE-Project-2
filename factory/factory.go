package factory

import (
	authDelivery "immersive-dashboard/features/auth/delivery"
	authRepo "immersive-dashboard/features/auth/repository"
	authService "immersive-dashboard/features/auth/service"

	userDelivery "immersive-dashboard/features/user/delivery"
	userRepo "immersive-dashboard/features/user/repository"
	userService "immersive-dashboard/features/user/service"

	teamDelivery "immersive-dashboard/features/team/delivery"
	teamRepo "immersive-dashboard/features/team/repository"
	teamService "immersive-dashboard/features/team/service"

	menteeDelivery "immersive-dashboard/features/mentee/delivery"
	menteeRepo "immersive-dashboard/features/mentee/repository"
	menteeService "immersive-dashboard/features/mentee/service"

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

	teamRepoFactory := teamRepo.New(db)
	teamServiceFactory := teamService.New(teamRepoFactory)
	teamDelivery.New(teamServiceFactory, e)

	menteeRepoFactory := menteeRepo.New(db)
	menteeServiceFactory := menteeService.New(menteeRepoFactory)
	menteeDelivery.New(menteeServiceFactory, e)
}
