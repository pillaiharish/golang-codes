package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(80 * time.Millisecond):
			return errors.New("task one failed")
		}
	})
	g.Go(func() error {
		select {
		case <-ctx.Done():
			fmt.Println("task two noticed cancellation")
			return ctx.Err()
		case <-time.After(time.Second):
			return nil
		}
	})
	fmt.Println("errgroup returned:", g.Wait())
}
