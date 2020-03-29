package exithandler

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/boilerplate/pkg/logger"
)

// Init accepts a callback function that will be invoked when program exits
// unexpectedly or is terminated by user. Used to perform cleanup: close DB
// connection, close network connections etc
func Init(cb func()) {
	sigs := make(chan os.Signal, 1)
	terminate := make(chan bool)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		logger.Error.Println("exit reason: ", sig)
		close(terminate)
	}()

	<-terminate
	cb()
}
