
/*
package controller

import (
	"fmt"
	"path/filepath"
	"net/http"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io"
	"os"
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

type AvatarParams struct {
	Token    string `json:"token"`
	Avatar string `json:"avatar"`
}
func SetAvatar(c echo.Context) error {
	var avatarParams AvatarParams
	err := json.NewDecoder(c.Request().Body).Decode(&avatarParams)
	if err != nil {
		fmt.Println("Error decoding request body:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"Error": "Error generating JWT"})
	}
	 
	fmt.Println("HELOOOOOOOOOOOOOOOOOOO")
	fmt.Println("AVATAR"+avatarParams.Avatar)
	fmt.Println("Tomen"+avatarParams.Token)
uploadHandler()
	return c.JSON(http.StatusOK, "{message: Success}")
} */


package controller

import (
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"os"
	"github.com/labstack/echo/v4"
)

func uploadHandler(c echo.Context) error {
	src, err := c.FormFile("my-file")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	file, err := src.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	defer file.Close()

	err = saveFile(file, src.Filename)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": fmt.Sprintf("File %s uploaded successfully!", src.Filename)})
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

type AvatarParams struct {
	Token  string `json:"token"`
	Avatar string `json:"avatar"`
}

func SetAvatar(c echo.Context) error {
	var avatarParams AvatarParams
	err := c.Bind(&avatarParams)
	if err != nil {
		fmt.Println("Error decoding request body:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Error decoding request body"})
	}

	fmt.Println("HELOOOOOOOOOOOOOOOOOOO")
	fmt.Println("AVATAR" + avatarParams.Avatar)
	fmt.Println("Token" + avatarParams.Token)

	// Call the file upload handler
	err = uploadHandler(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Success"})
}
