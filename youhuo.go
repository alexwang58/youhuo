package main

import (
	"flag"
	"fmt"
	"github.com/alexwang58/youhuo/yhbase"
	"log"
	"os"
)

var (
	YH_POART    = 8090
	YH_DOCROOT  = ""
	logFileName = flag.String("log", "cServer.log", "Log file name")
)

func initLog() {
	logFile, logErr := os.OpenFile(*logFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if logErr != nil {
		fmt.Println("Fail to find", *logFile, "cServer start Failed")
		os.Exit(1)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime /* | log.Lshortfile */)
}

func main() {
	flag.Parse()
	initLog()
	fmt.Println("Youhuo Start...")
	log.Println("Youhuo Start...")
	if err := yhbase.StartYhServer(YH_POART, YH_DOCROOT); err != nil {
		log.Fatalf("YHServer satart faild: [Error]: %v", err)
		return
	}
}
