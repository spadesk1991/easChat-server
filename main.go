package main

import (
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/spadesk1991/easChat-server/controllers"
	"github.com/spadesk1991/easChat-server/middleware"
)

var (
	clientMap = make(map[bson.ObjectId]*websocket.Conn)
)

func main() {
	r := gin.Default()
	r.Use(middleware.Cors())
	r.GET("/ws/:id", wsPage)
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.GET("/user", controllers.GetUser)
	r.GET("/users", controllers.GetUsers)
	r.Run(":12345")
}

type Message struct {
	Sender    bson.ObjectId `json:"sender"`
	Recipient bson.ObjectId `json:"recipient"`
	Content   string        `json:"content"`
}

func wsPage(ctx *gin.Context) {
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Fatalln(err)
		return
	}
	id := ctx.Param("id")
	if id == "" {
		return
	}
	clientMap[bson.ObjectIdHex(id)] = conn
	for {
		jsonMsg := Message{}
		err := conn.ReadJSON(&jsonMsg)
		// messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		// json.Unmarshal(p, jsonMsg)
		recipConn := clientMap[jsonMsg.Recipient]
		jsonMsg.Sender, jsonMsg.Recipient = jsonMsg.Recipient, jsonMsg.Sender
		if recipConn != nil {
			if err := recipConn.WriteJSON(jsonMsg); err != nil {
				log.Println(err)
				return
			}

		}
	}
}
