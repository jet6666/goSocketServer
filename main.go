package main

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/gorilla/websocket"
	"gorock/chat"
	"log"
	"net/http"
	"strconv"
)

var hub = chat.NewHub()

func main() {
	fmt.Println("websocket start")
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "homepage")

	})

	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println("snowflake 算法失败", err)
		return
	}

	go hub.Run()

	upgrader := websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}
	http.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		upgrader.CheckOrigin = func(r *http.Request) bool {
			return true
		}
		ws, err := upgrader.Upgrade(writer, request, nil)
		if err != nil {
			log.Println("upgrade failed ", err)
		}
		id := node.Generate().String()
		online := strconv.Itoa(hub.OnlineMembers() + 1 )

		log.Println("client connected ,RemoteAddr = ", request.RemoteAddr+",uid = "+id)

		//给当前client hi client
		err = ws.WriteMessage(1, []byte("Welcome : "+id+"("+request.RemoteAddr+") ,online members:"+online ))
		if err != nil {
			log.Println("write 1 hello err : ", err)
		}
		hub.SendAll([]byte ( fmt.Sprintf("%s is coming ,online members :%s", id, online)))

		client := &chat.WSClient{
			Hub:  hub,
			Conn: ws,
			Send: make(chan []byte, 256),
			Id:   id,
			Ip:   request.RemoteAddr,
		}
		hub.Register <- client

		//处理当前conn的消息
		go client.WsWriter(ws)
		go client.WsReader(ws)
	})
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln("unable to start websocket server 8080 ")
	}

}
