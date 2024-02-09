/*package controller

import (
    "net/http"
    "io"
    "os"
    "strings"
)

func Imgg(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Query().Get("title")
    imagePath := "static/" + title
    file, err := os.Open(imagePath)
    if err != nil {
        http.Error(w, "Image not found", http.StatusNotFound)
        return
    }
    defer file.Close()

    // Определение MIME-типа изображения
    var contentType string
    if strings.HasSuffix(title, ".jpg") || strings.HasSuffix(title, ".jpeg") {
        contentType = "image/jpeg"
    } else if strings.HasSuffix(title, ".png") {
        contentType = "image/png"
    } else {
        // По умолчанию
        contentType = "image/jpeg"
    }

    // Установка заголовка Content-Type
    w.Header().Set("Content-Type", contentType)

    // Копирование содержимого файла в ответ
    io.Copy(w, file)
}
*/
package controller

import (
    "github.com/labstack/echo/v4"
    "net/http"
    "io"
    "os"
    "strings"
)

 
func ImgEcho(c echo.Context) error {
    w := c.Response().Writer
    r := c.Request()
    title := r.URL.Query().Get("title")
    imagePath := "static/" + title
    file, err := os.Open(imagePath)
    if err != nil {
        return c.String(http.StatusNotFound, "Image not found")
    }
    defer file.Close()
 
    var contentType string
    if strings.HasSuffix(title, ".jpg") || strings.HasSuffix(title, ".jpeg") {
        contentType = "image/jpeg"
    } else if strings.HasSuffix(title, ".png") {
        contentType = "image/png"
    } else {
 
        contentType = "image/jpeg"
    }
 
    c.Response().Header().Set("Content-Type", contentType)
    _, err = io.Copy(w, file)
    if err != nil {
        return err
    }

    return nil
}
// http://localhost:5000/worklist.com/image?title=heart.png