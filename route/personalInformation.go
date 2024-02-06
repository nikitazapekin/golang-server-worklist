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
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"server/controller"
	//"encoding/json"
	m "server/db"
	e "server/middleware"
	//	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func img(c echo.Context) error {
	token := c.QueryParam("token")
	decodedToken, _ := e.Decode(token, "key")
	if user, err := m.FindUserByUsername(decodedToken.Username); err == nil {
		filePath := "static/" + user.Avatar
		fmt.Println("USER AVATAR")
		fmt.Println(decodedToken)
		fmt.Println(user)
		fmt.Println("FILEEEPATH " + filePath)
		fmt.Println(c.File(filePath))
		response := map[string]string{
			"avatar": filePath,
		}

		return c.JSON(http.StatusOK, response)
	}
	return c.JSON(http.StatusBadRequest, "{message: error}")
}

func uploadHandler(c echo.Context) error {
	token := c.QueryParam("token")

	decodedToken, _ := e.Decode(token, "key")
	fmt.Println(decodedToken)
	src, err := c.FormFile("my-file")
	fmt.Println("SRCCCCCCCCC" + src.Filename) // НАЗВАНИЕ КОТОРОЕ НАДО ЗАГНАТЬ В БД
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
		fmt.Println(user)
		fmt.Println("New avaaaaaaaaaaaaaaaaaaaa")
		fmt.Println(user.Avatar)
		newAvatar := string(src.Filename)

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
func getAbsoluteURL(c echo.Context, relativePath string) string {
	host := c.Request().Host
	scheme := "http"
	if c.Request().TLS != nil {
		scheme = "https"
	}
	return fmt.Sprintf("%s://%s/%s", scheme, host, relativePath)
}

var globalArray []string

/*
func uploadHandlerMultiple(c echo.Context) error {
	token := c.QueryParam("token")

	decodedToken, _ := e.Decode(token, "key")
	fmt.Println(decodedToken)
	err := c.Request().ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	form := c.Request().MultipartForm
	files := form.File["my-files"]
	for _, file := range files {
		fmt.Println("SRCCCCCCCCC" + file.Filename) // НАЗВАНИЕ КОТОРОЕ НАДО ЗАГНАТЬ В БД

		src, err := file.Open()
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		defer src.Close()

		err = saveFile(src, file.Filename)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		fmt.Println("URLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLL")
		filePath := "static/" + file.Filename
		fmt.Println(filePath)
		fmt.Println("File URL:", getAbsoluteURL(c, filePath))

		//fmt.Println(c.File(file.Filename))
	//	return c.JSON(http.StatusOK,  getAbsoluteURL(c, filePath))
	}

	return c.String(http.StatusOK, "Files uploaded successfully!")
} */

func uploadHandlerMultiple(c echo.Context) error {
	token := c.QueryParam("token")

	decodedToken, _ := e.Decode(token, "key")
	fmt.Println(decodedToken)
	err := c.Request().ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	form := c.Request().MultipartForm
	files := form.File["my-files"]
	var fileURLs []string // Добавленный массив для хранения URL-ов файлов
	for _, file := range files {
		fmt.Println("SRCCCCCCCCC" + file.Filename) // НАЗВАНИЕ КОТОРОЕ НАДО ЗАГНАТЬ В БД

		src, err := file.Open()
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		defer src.Close()

		err = saveFile(src, file.Filename)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		fmt.Println("URLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLL")
		filePath := "static/" + file.Filename
		fmt.Println(filePath)

		fileURL := getAbsoluteURL(c, filePath)
		fmt.Println("File URL:", fileURL)
		globalArray = append(globalArray, fileURL)
		fileURLs = append(fileURLs, fileURL)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "Files uploaded successfully!",
		"fileURLs": fileURLs,
	})
}
func SetPersonalInformation(e *echo.Echo) {
	e.POST("/worklist.com/getPersonalInformation/editPersonalData", controller.EditPersonalInformation)
	e.GET("/worklist.com/getPersonalInformation", controller.GetPersonalInformation)
	e.POST("/worklist.com/getPersonalInformation/setAvatar", controller.SetAvatar)
	e.POST("/test", uploadHandler)
	e.POST("/testMultiple", uploadHandlerMultiple)
	e.POST("/worklist.com/createOffer", controller.CreateVacancy)
	e.GET("/worklist.com/getPersonalInformation/getAvatar", img)
}

//http://localhost:5000/worklist.com/getPersonalInformation/getAvatar
