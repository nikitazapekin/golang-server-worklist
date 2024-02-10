package controller

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
	m "server/db"
)
func GetAmountOfOffers(c echo.Context) error {
  
elems,err :=m.GetAmountOfOffers()
if(err!=nil) {
	return c.JSON(http.StatusBadRequest, "{error: error}")
} 
	amountOfOffers:=0;
	for i := range elems {
	
		  fmt.Println(i)
		  amountOfOffers++;
}

resObj := map[string]interface{}{
	"amountOfOffers": amountOfOffers,
	"error": nil,
	 

}
return c.JSON(http.StatusOK,resObj) 
	
} 