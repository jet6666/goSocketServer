package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
)

type tbLog interface {
	Log(...interface{})
}
type testLogger struct {
	tbLog
}

func (tl *testLogger) Output(maxDepth int, s string) error {
	tl.Log(s)
	return nil
}

//func newTestLogger(tbl tbLog) logger {
//	return &testLogger{tbl}
//}

type testHandler struct {
}

func (h *testHandler) HandleMessage(message *nsq.Message) error {
	if len(message.Body) == 0 {
		log.Println("message len = 0")
		return nil
	}
	log.Println("get message " , string(message.Body))
	//switch string(message.Body ) {
	//
	//}
	return nil
}

func main() {
	fmt.Println("this is nsq consumer  ")

	config := nsq.NewConfig()
	//laddr :="192.168.13.65"
	//config.LocalAddr ,_ =net.ResolveIPAddr("tcp",laddr+":0")
	//config.DefaultRequeueDelay =0
	//config.MaxBackoffDuration = time.Millisecond *50
	q, err := nsq.NewConsumer("test", "lc3", config)
	if err != nil {
		log.Fatalln("create nsq consumer fail ", err.Error())
		return
	}
	//q.SetLogger(newTestLogger(t) , nsq.LogLevelDebug)
	q.AddHandler(&testHandler{})
	err = q.ConnectToNSQLookupd("192.168.13.65:5161")
	if err != nil {
		log.Fatalln("ConnectToNSQLookupd error :" , err.Error())
		return
	}
	defer q.Stop()
	select {}

}
