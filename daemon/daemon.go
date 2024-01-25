package daemon

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/Masterminds/log-go"

	"github.com/werks/voyeur-go/agent"
)

type Daemon struct {
	agent *agent.Agent
}

func Run() error {

	// TODO: initialize logging

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	d := &Daemon{
		agent: agent.InitAgent(),
	}

	// hook signals for shutdown
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		log.Infof("received signal: %s", sig.String())
		cancel() // ==> shutdown the agent
	}()

	// start the agent and wait forever
	log.Info("voyeur agent starting")
	d.agent.Start(ctx)
	err := d.agent.AwaitShutdown()
	log.Info("voyeur agent shutdown")
	if err != nil {
		log.Warnf("voyeur agent exited with error: %w", err)
	}

	return err
}
