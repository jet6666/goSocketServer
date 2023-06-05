package chat

import "log"

//其实是客户端管理器
type Hub struct {
	clients    map[*WSClient]bool
	broadcast  chan []byte
	Register   chan *WSClient
	Unregister chan *WSClient
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*WSClient]bool),
		broadcast:  make(chan []byte),
		Register:   make(chan *WSClient),
		Unregister: make(chan *WSClient),
	}
}
func (h *Hub) SendAll(message []byte) {
	h.broadcast <- message
}

func (h *Hub) OnlineMembers() int {
	return len(h.clients)
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.clients[client] = true
		case client := <-h.Unregister:
			if _, ok := h.clients[client]; ok {
				log.Println("Unregister  client ")
				delete(h.clients, client)
				close(client.Send)
			}
		case message := <-h.broadcast:
			log.Println("Broadcast message : " + string(message))
			for client := range h.clients {
				select {
				case client.Send <- message:
					//log.Println("got new message ", message)
				default:
					//client offline ? cant receive message ??
					log.Println("unreachable  broadcast ,  close client  ", message)
					close(client.Send)
					delete(h.clients, client)

				}
			}

		}
	}

}
