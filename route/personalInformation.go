/*
package router
import (
	"server/controller"
	"github.com/labstack/echo/v4"

	


	"fmt"
	"net/http"
	"io"
	"os"
	"path/filepath"



)


func uploadHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		src, hdr, err := req.FormFile("my-file")
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		defer src.Close()

		err = saveFile(src, hdr.Filename)
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}

		fmt.Fprintf(res, "File %s uploaded successfully!", hdr.Filename)
		return
	}
}

func saveFile(src io.Reader, filename string) error {
	dst, err := os.Create(filepath.Join("static", filename))
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}

	return nil
}

func SetPersonalInformation(e *echo.Echo) {
	e.POST("/worklist.com/getPersonalInformation/editPersonalData", controller.EditPersonalInformation)
	e.GET("/worklist.com/getPersonalInformation", controller.GetPersonalInformation)
	e.POST("/worklist.com/getPersonalInformation/setAvatar", controller.SetAvatar)
	e.POST("/test", uploadHandler(e *echo.Echo))
}

*/


package router

import (
	"server/controller"
	"github.com/labstack/echo/v4"
	"fmt"
	"io"
	m "server/db"
	e "server/middleware"
//	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)



func img(c echo.Context) error {
	token := c.QueryParam("token")
	decodedToken, _ := e.Decode(token,  "key")
	if user, err := m.FindUserByUsername(decodedToken.Username); err == nil {
		filePath := "static/"+user.Avatar
	//	filePath := "static/heart.png"
		fmt.Println("USER AVATAR")
		fmt.Println(decodedToken)
		fmt.Println(user)
		fmt.Println("FILEEEPATH "+filePath)
		fmt.Println( c.File(filePath))
		response := map[string]string{
            "avatar": filePath,
        }

		return c.JSON(http.StatusOK, response)
	//	return c.File(filePath)
	} 
	return c.JSON(http.StatusBadRequest, "{message: error}")
}

 func uploadHandler(c echo.Context) error {
	token := c.QueryParam("token")
	 
	decodedToken, _ := e.Decode(token,  "key")
	fmt.Println(decodedToken)
	src, err := c.FormFile("my-file")
	fmt.Println("SRCCCCCCCCC"+src.Filename) // НАЗВАНИЕ КОТОРОЕ НАДО ЗАГНАТЬ В БД
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	file, err := src.Open()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	defer file.Close()
	fmt.Println("FILE")
fmt.Println(file)
	err = saveFile(file, src.Filename)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if user, err := m.FindUserByUsername(decodedToken.Username); err == nil {

		fmt.Println("USER AVATAR")
		fmt.Println(user)
//user.Avatar=string(src.Filename)
fmt.Println("New avaaaaaaaaaaaaaaaaaaaa")
fmt.Println(user.Avatar)
newAvatar :=string(src.Filename)

errr := m.UpdateUserAvatar(user, newAvatar)
fmt.Println("ERROR ", errr)
	fmt.Println("Everything is clear")
	fmt.Println(user)
	}
	return c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully!", src.Filename))
}

func saveFile(src io.Reader, filename string) error {
	dst, err := os.Create(filepath.Join("static", filename))

	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}

	return nil
}

func SetPersonalInformation(e *echo.Echo) {
	e.POST("/worklist.com/getPersonalInformation/editPersonalData", controller.EditPersonalInformation)
	e.GET("/worklist.com/getPersonalInformation", controller.GetPersonalInformation)
	e.POST("/worklist.com/getPersonalInformation/setAvatar", controller.SetAvatar)
	e.POST("/test", uploadHandler)
	e.GET("/worklist.com/getPersonalInformation/getAvatar", img)
}
//http://localhost:5000/worklist.com/getPersonalInformation/getAvatar