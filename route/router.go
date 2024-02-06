package router

import (
	"github.com/labstack/echo/v4"
)




func InitRoutes(e *echo.Echo) {
	SetAuth(e)
	SetBookRoutes(e)
	SetOffersRoutes(e)
	SetPersonalInformation(e)
	SetVacancy(e)
}
//    return offersApiInstance.get(`/getOffers`)