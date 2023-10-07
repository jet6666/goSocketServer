package main

import (
	"goSocketServer/example/log/loglib"
	"io"
	"log"
	"os"
)

func main ()  {
	logfile ,err :=os.OpenFile("./s1.log" ,os.O_RDWR |os.O_CREATE |os.O_APPEND ,0766)
	if err !=nil {
		panic(err.Error() )
	}
	defer func() {
		logfile.Close()
	}()
	//multiwriter
	multiWriter :=io.MultiWriter(os.Stdout ,logfile)
	log.SetOutput(multiWriter)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("line1111 ")
	loglib.LogLibExample()

	
}
