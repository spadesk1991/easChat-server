// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/websocket"
// 	uuid "github.com/satori/go.uuid"
// 	"gopkg.in/mgo.v2"
// )

// // ClientManager 客户端管理
// type ClientManager struct {
// 	clients    map[*Client]bool
// 	broadcast  chan []byte
// 	register   chan *Client
// 	unregister chan *Client
// }

// // Client 客户端
// type Client struct {
// 	id     string
// 	socket *websocket.Conn
// 	send   chan []byte
// }

// // Message 消息
// type Message struct {
// 	Sender    string `json:"sender,omitempty"`
// 	Recipient string `json:"recipient,omitempty"`
// 	Content   string `json:"content,omitempty"`
// }

// var manager = ClientManager{
// 	broadcast:  make(chan []byte),
// 	register:   make(chan *Client),
// 	unregister: make(chan *Client),
// 	clients:    make(map[*Client]bool),
// }

// func (manager *ClientManager) start() {
// 	for {
// 		select {
// 		case client := <-manager.register:
// 			manager.clients[client] = true
// 			jsonMessage, _ := json.Marshal(&Message{Content: "/A new socket has connected."})
// 			manager.send(jsonMessage, client)
// 		case client := <-manager.unregister:
// 			if _, ok := manager.clients[client]; ok {
// 				close(client.send)
// 				delete(manager.clients, client)
// 				jsonMessage, _ := json.Marshal(&Message{Content: "/A socket has disconnected."})
// 				manager.send(jsonMessage, client)
// 			}
// 		case message := <-manager.broadcast:
// 			for conn := range manager.clients {
// 				select {
// 				case conn.send <- message:
// 				default:
// 					close(conn.send)
// 					delete(manager.clients, conn)
// 				}
// 			}
// 		}
// 	}
// }

// func (manager *ClientManager) send(message []byte, ignore *Client) {
// 	for conn := range manager.clients {
// 		if conn != ignore {
// 			conn.send <- message
// 		}
// 	}
// }

// func (c *Client) read() {
// 	defer func() {
// 		manager.unregister <- c
// 		c.socket.Close()
// 	}()

// 	for {
// 		_, message, err := c.socket.ReadMessage()
// 		if err != nil {
// 			manager.unregister <- c
// 			c.socket.Close()
// 			break
// 		}
// 		jsonMessage, _ := json.Marshal(&Message{Sender: c.id, Content: string(message)})
// 		log.Println(string(jsonMessage))
// 		manager.broadcast <- jsonMessage
// 	}
// }

// func (c *Client) write() {
// 	defer func() {
// 		c.socket.Close()
// 	}()

// 	for {
// 		select {
// 		case message, ok := <-c.send:
// 			if !ok {
// 				c.socket.WriteMessage(websocket.CloseMessage, []byte{})
// 				return
// 			}

// 			c.socket.WriteMessage(websocket.TextMessage, message)
// 		}
// 	}
// }

// func main() {
// 	fmt.Println("Starting application...")
// 	go manager.start()
// 	http.HandleFunc("/ws", wsPage)
// 	http.ListenAndServe("localhost:12345", nil)
// }

// func wsPage(res http.ResponseWriter, req *http.Request) {
// 	conn, error := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(res, req, nil)
// 	if error != nil {
// 		http.NotFound(res, req)
// 		return
// 	}
// 	client := &Client{id: uuid.NewV4().String(), socket: conn, send: make(chan []byte)}

// 	manager.register <- client

// 	go client.read()
// 	go client.write()
// }
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	http.HandleFunc("/ws", ws)
	http.ListenAndServe("localhost:12345", nil)
}
func ws(res http.ResponseWriter, req *http.Request) {
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(res, req, nil)
	if err != nil {
		log.Fatalln(err)
		return
	}
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		if string(p) == "connent" {

		}
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}
