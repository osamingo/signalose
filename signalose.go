package signalose

import (
	"fmt"
	"io"
	"os"
	"os/signal"
)

// Writer writes signalose logs.
var Writer io.Writer = os.Stderr

// AddCloser performs close method when the registered signal is generated.
func AddCloser(tag string, closer io.Closer, trigger ...os.Signal) (chan os.Signal, error) {

	if tag == "" {
		return nil, fmt.Errorf("signalose: tag should not be empty")
	}

	if closer == nil {
		return nil, fmt.Errorf("signalose: closer should not be nil")
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, trigger...)

	go wait(tag, closer, sigChan)

	return sigChan, nil
}

func wait(tag string, closer io.Closer, sigChan chan os.Signal) {
	for {
		<-sigChan
		if closer != nil {
			err := closer.Close()
			msg := fmt.Sprintf("signalose: %s is close", tag)
			if err != nil {
				msg = msg + fmt.Sprintf(", err = %v", err)
			}
			Writer.Write([]byte(msg + "\n"))
		}
	}
}
