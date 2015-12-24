package main

import (
	"log"
	"syscall"
	"time"

	"github.com/fluent/fluent-logger-golang/fluent"
	"github.com/osamingo/signalose"
)

func main() {

	f, err := fluent.New(fluent.Config{})
	if err != nil {
		log.Fatal(err)
	}

	sigChan, err := signalose.AddCloser("fluentdClient", f, syscall.SIGUSR2)
	if err != nil {
		log.Fatal(err)
	}

	err = f.Post("example_signalise", map[string]string{"message": "test"})
	if err != nil {
		log.Fatal(err)
	}

	sigChan <- syscall.SIGUSR2

	time.Sleep(time.Second)
}
