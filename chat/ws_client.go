package chat

import (
	"bytes"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 10 * time.Second //60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

type WSClient struct {
	Hub    *Hub
	Conn   *websocket.Conn
	Send   chan []byte
	Id     string
	Ip     string
	Online bool
}

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

func (c *WSClient) WsReader(conn *websocket.Conn) {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()
	c.Conn.SetReadLimit(maxMessageSize)
	//ping pong heart break(embed in server/client(chrome/edge) )
	//c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error {
		//log.Println("got  pong ")
		//c.Conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	//read message
	for {
		messageType, message, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println("read message err :", string(message), string(messageType))
			log.Println(err)
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error --- : %v ", err)
			}
			quitMessage := fmt.Sprintf("%s is quiting.online members :%d", c.Id, c.Hub.OnlineMembers()-1)
			c.Hub.broadcast <- []byte(quitMessage)
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		message2 := []byte ("[" + c.Id + "] says : ")
		message2 = append(message2, message...)
		c.Hub.broadcast <- message2
	}

}

func (c *WSClient) WsWriter(conn *websocket.Conn) {
	//ping pong
	//newTicker (!!!not NewTimer )
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()
	//从通道中获取写入消息
	for {
		select {
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				log.Println("write close !!!!!!!!! ")
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			//Coalesce outbound messages in chat example
			n := len(c.Send)
			log.Println("c.send outbound message len = ", n)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.Send)
			}
			if err := w.Close(); err != nil {
				return
			}

		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
			//log.Println("ping message ")
		}
	}
}
