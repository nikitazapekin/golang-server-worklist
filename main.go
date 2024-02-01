
/*

package main

import (
	"fmt"
	"log"
"server/db"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	r "server/route"
	"github.com/labstack/echo/v4/middleware"





	 
	"io"
	"net/http"
	"path/filepath"
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
func main() {
	//codeEmail := "582455"
	e := echo.New()
	db.Connect()
	e.Use(middleware.CORS())
	r.InitRoutes(e)

	err := e.Start(":5000")
	//e.uploadHandler(e)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
} 
 */

 /*

// РАБОТАААААААААААААААААААААААААААААААААААААААЕТ
 package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("STATIC DIR:", "static")

	http.ListenAndServe(":5000", http.HandlerFunc(uploadHandler))
}

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
 */




 package main

import (
	"fmt"
	



	"log"
"server/db"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	r "server/route"
	"github.com/labstack/echo/v4/middleware"
)


func main() {
	fmt.Println("STATIC DIR:", "static")
	e := echo.New()
	db.Connect()
	e.Use(middleware.CORS())
	r.InitRoutes(e)
	err := e.Start(":5000")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
 