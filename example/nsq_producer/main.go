package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
	"strconv"
	"sync"
)

func main ()  {
	config :=nsq.NewConfig()
	p,err :=nsq.NewProducer("192.168.13.65:5150" ,config )
	if err!=nil {
		log.Fatalln("create nsq producer failed " ,err.Error())
		return
	}

	/*for {
		time.Sleep(time.Second)
		now :=time.Now().String()

		if err :=p.Publish("test" ,[]byte(now)) ;err !=nil {
			log.Println("nsql publish fail " , err.Error() )
		}
	}*/

	var wg sync.WaitGroup
	messageCount :=5
	wg.Add(messageCount)
	//跑完5个
	for i:=0;i<messageCount   ;i++ {

		go func() {
			defer wg.Done()
			if err :=p.Publish("test" ,[]byte("aaaa" + strconv.Itoa(i))) ;err !=nil {
				log.Println("nsql publish fail " , err.Error() )
			}
			//wg.Done()
		}()


	}

	wg.Wait()
	fmt.Println("send finished ")

	
}
