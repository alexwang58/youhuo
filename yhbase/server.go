package yhbase

import (
	"fmt"
	"github.com/alexwang58/youhuo/components"
	"log"
	"net/http"
	"strconv"
)

type YouHuoEnv struct {
	port    int
	docRoot string
}

var (
	YHSENV = new(YouHuoEnv)
	//logFileName = flag.String("log", "cServer.log", "Log file name")
)

// Start YouHuo Http Server
// set port and web server root path
func StartYhServer(port int, docRoot string) error {
	YHSENV.port = port
	YHSENV.docRoot = docRoot
	components.GlobalComponentInit()
	initRouter()
	if err := httpStart(port); err != nil {
		//fmt.Printf("YouHuo Server Start faild: %v", err.Error())
		log.Fatalf("YouHuo Server Start faild: %v", err.Error())
	}
	return nil
}

func httpStart(port int) error {
	fmt.Printf("YouHuo Server Start :%v\n", port)
	log.Printf("YouHuo Server Start :%v", port)
	return http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
