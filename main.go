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
  */




  /*
  package main

import (
	"fmt"
	"log"
	"net/http"
	"server/db"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/gorilla/websocket"
	r "server/route"
)

 
func main() {
	fmt.Println("STATIC DIR:", "static")
	e := echo.New()
	db.Connect()
	e.Use(middleware.CORS())
	r.InitRoutes(e)

	// Добавление обработчика веб-сокетов
	//e.GET("/ws", handleWebSocket)
	wsUpgrader := websocket.Upgrader{
        ReadBufferSize:  1024,
        WriteBufferSize: 1024,
        CheckOrigin: func(r *http.Request) bool {
            return true // Allow any origin
        },
    }

    // Listen for WebSocket connections on port 8080.
    http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
        // Upgrade the HTTP connection to a WebSocket connection.
        conn, err := wsUpgrader.Upgrade(w, r, nil)
        if err != nil {
            fmt.Println(err)
            return
        }
        // Read messages from the client.
        for {
            // Read a message from the client.
            messageType, message, err := conn.ReadMessage()
            if err != nil {
                fmt.Println(err)
                return
            }
            // Print the message to the console.
            fmt.Println("Received:", message)
            // Send a message back to the client.
            err = conn.WriteMessage(messageType, []byte("Hello, client!"))
            if err != nil {
                fmt.Println(err)
                return
            }
        }
    })






	err := e.Start(":5000")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
 */











 
 package main

import (
	"fmt"
	"log"
	"net/http"
	"server/db"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/gorilla/websocket"
	r "server/route"
)
var numberOfOnlineUsers = 0 
var clients = make(map[*websocket.Conn]bool)
var onlineUsersMessage = "%d"
var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true 
		},
	}
)

 func handleWebSocket(c echo.Context) error {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	clients[conn] = true
	numberOfOnlineUsers++
	broadcastOnlineUsersMessage()
	defer func() {
		numberOfOnlineUsers--
		delete(clients, conn)
		broadcastOnlineUsersMessage()
		conn.Close()
	}()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			return err
		}
		fmt.Printf("Received message: %s\n", p)
		fmt.Println("CONNNN")
		fmt.Println(conn)
		fmt.Println(messageType)
		fmt.Println(p)
		err = conn.WriteMessage(messageType, []byte("Hello, client!"))
		if err != nil {
			return err
		}
	}
}
 
func broadcastOnlineUsersMessage() {
	message := fmt.Sprintf(onlineUsersMessage, numberOfOnlineUsers)
	fmt.Println("MESSSSSSSSSAAAGE")
	fmt.Println(message)
	fmt.Println(clients)
	for client := range clients {
		err := client.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			log.Printf("Error sending message to client: %v", err)
		}
	}
}

func main() {
	fmt.Println("STATIC DIR:", "static")
	e := echo.New()
	db.Connect()
	e.Use(middleware.CORS())
	r.InitRoutes(e)
 
	e.GET("/ws", handleWebSocket)

	err := e.Start(":5000")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
 




