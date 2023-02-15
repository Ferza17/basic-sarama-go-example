package main

import (
	"log"
	"runtime"

	"github.com/ferza17/kafka-basic/consumer/cmd"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("err execute command err : %v", err)
	}
}
