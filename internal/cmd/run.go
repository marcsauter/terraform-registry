package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/alecthomas/kong"
	"github.com/marcsauter/terraform-registry/internal/registry"
	"github.com/marcsauter/terraform-registry/internal/version"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
)

type runCmd struct {
}

func (r runCmd) Run(app *kong.Context, g *Globals, l *zap.SugaredLogger, v *version.Version) error {
	l.Info("start Terraform registry")

	ctx := contextWithSignal(context.Background(), func(s os.Signal) {
		l.Infow("stopping server", "signal", s.String())
	}, syscall.SIGINT, syscall.SIGTERM)

	reg, err := registry.New(l, prometheus.NewRegistry(),
		registry.WithHTTPListen(g.HTTPListen),
	)
	if err != nil {
		return err
	}

	return reg.Run(ctx)
}

func contextWithSignal(ctx context.Context, f func(s os.Signal), s ...os.Signal) context.Context {
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, s...)

		defer signal.Stop(c)

		select {
		case <-ctx.Done():
		case sig := <-c:
			if f != nil {
				f(sig)
			}

			cancel()
		}
	}()

	return ctx
}