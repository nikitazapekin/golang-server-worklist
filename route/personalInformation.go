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
	"net/http"
	"os"
	"path/filepath"
)

func uploadHandler(c echo.Context) error {
	src, err := c.FormFile("my-file")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	file, err := src.Open()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	defer file.Close()

	err = saveFile(file, src.Filename)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
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
}
