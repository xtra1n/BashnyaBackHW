package cli

import (
	"context"
	"os"
	"os/signal"
)

func ProcessSignal() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)

	go func() {
		<-signalChannel
		cancel()
		signal.Stop(signalChannel)
		close(signalChannel)
	}()

	return ctx, cancel
}
