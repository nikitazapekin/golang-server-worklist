package controller

import (
	//"fmt"
	"net/http"
	"fmt"

	"github.com/labstack/echo/v4"
		m "server/db"
)
func GetOffer(c echo.Context) error {
	id := c.QueryParam("id")
	fmt.Println(id)

	element,err :=m.FindOfferById(id)
if(err!=nil) {
	fmt.Println(err)
}
	fmt.Println(element)
	/*
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
return c.JSON(http.StatusOK,resObj)  */
return c.JSON(http.StatusOK, element)
//return nil
	
} 