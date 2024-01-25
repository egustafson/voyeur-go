package agent

import (
	"context"
	"time"
)

type Agent struct {
	ctx    context.Context
	cancel context.CancelFunc
}

type AgentOps struct {
	timeout time.Duration
}

type AgentOption func(*AgentOps)

func WithTimeout(t time.Duration) AgentOption {
	return func(o *AgentOps) {
		o.timeout = t
	}
}

func processAgentOpts(opts []AgentOption) *AgentOps {
	ao := &AgentOps{}
	for _, opt := range opts {
		opt(ao)
	}
	return ao
}

// InitAgent initializes a new Agent.  The agent is inactive until Agent.Start()
// is invoked.
func InitAgent() *Agent {
	a := &Agent{
		ctx: nil,
	}
	return a
}

// Start starts the Agent's main event loop in a separate goroutine and returns.
func (a *Agent) Start(ctx context.Context) error {
	a.ctx, a.cancel = context.WithCancel(ctx)
	go a.mainEventLoop()
	return nil
}

// mainEventLoop is the main event loop.  It should be started as a new
// goroutine by Start() and exits when the Agent's context is canceled.
func (a *Agent) mainEventLoop() {
	for {
		select {
		case <-a.ctx.Done():
			// TODO: log exiting Agent's main event loop
			return
		}
	}
}

// Shutdown stops the Agent goroutine.
func (a *Agent) Shutdown(opts ...AgentOption) error {
	a.cancel()
	return a.AwaitShutdown(opts...)
}

func (a *Agent) AwaitShutdown(opts ...AgentOption) error {
	ops := processAgentOpts(opts)

	if ops.timeout > 0 {
		timedCtx, timedCancel := context.WithTimeout(context.Background(), ops.timeout)
		defer timedCancel()
		select {
		case <-timedCtx.Done():
			return timedCtx.Err()
		case <-a.ctx.Done():
		}
	} else {
		// wait forever
		<-a.ctx.Done()
	}
	return nil
}
