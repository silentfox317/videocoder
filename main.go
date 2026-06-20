package main

import (
	"github.com/dirshunt/videocoder/cmd"
	"github.com/dirshunt/videocoder/internal/logging"
)

func main() {
	defer logging.RecoverPanic("main", func() {
		logging.ErrorPersist("VideoCoder terminated due to unhandled panic")
	})
	cmd.Execute()
}
