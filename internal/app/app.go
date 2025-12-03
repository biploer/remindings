package app

import (
	"context"
)

func Run(ctx context.Context) error {
	go func() {
		<-ctx.Done()
	}()
	return nil
}
