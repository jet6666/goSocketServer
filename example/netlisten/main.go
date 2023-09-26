package main

import (
	"log"
	"net"
)

func main() {

	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Println("net.listen error ", err.Error())
	} else {
		log.Println("listening ", ln)
	}

	go func() {
		i:= 1
		 for {
			 i++
		 }

	}()

	select {

	}
}
