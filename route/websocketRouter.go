package router

import (
	"log"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"github.com/gorilla/websocket"
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
 
 func handleWebSocket(c echo.Context) error {                //ХЭНДЛЕРРРРРРРРРРРРРРРР
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	clients[conn] = true
	numberOfOnlineUsers++
	broadcastOnlineUsersMessage()
//	sendMessageToClient()
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



func handlePublicChat(c echo.Context) error {
    conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
    if err != nil {
        return err
    }
    clients[conn] = true

    defer func() {
        delete(clients, conn)
        conn.Close()
    }()

    // Прочитать сообщение от клиента
    _, msgBytes, err := conn.ReadMessage()
    if err != nil {
        log.Printf("Error reading message from client: %v", err)
        return err
    }

    // Передать сообщение в функцию handlePublicMessages
    handlePublicMessages(string(msgBytes))

    return nil
}

func handlePublicMessages(message string) error {
    fmt.Println("Received message CHATTTTTTTTTTTTTTTTTTTTTTTTTTTT:", message)
    fmt.Println(clients)
    for client := range clients {
        err := client.WriteMessage(websocket.TextMessage, []byte(message))
        if err != nil {
            log.Printf("Error sending message to client: %v", err)
        }
    }
    return nil
}

func InitWebsocketRoutes(e *echo.Echo) {
	e.GET("/ws", handleWebSocket)
	e.GET("/publicMessages",handlePublicChat  )
 
}