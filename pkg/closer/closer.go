package closer

import (
	"context"
	"fmt"
	"log/slog"
	"sync"

	"go.uber.org/multierr"
)

type closeFn func(context context.Context) error

type item struct {
	name string
	fn   closeFn
}

type Closer struct {
	mu     sync.Mutex
	items  []item
	logger slog.Logger
}

func New(logger slog.Logger) *Closer {
	return &Closer{
		logger: logger,
	}
}

func (c *Closer) Add(name string, fn closeFn) {
	c.mu.Lock()
	c.items = append(c.items, item{name, fn})
	c.mu.Unlock()
}

func (c *Closer) AddFunc(name string, fn func()) {
	c.Add(name, func(ctx context.Context) error {
		fn()
		return nil
	})
}

func (c *Closer) Close(ctx context.Context) (e error) {
	c.mu.Lock()
	items := append([]item(nil), c.items...)
	c.mu.Unlock()

	for _, item := range items {
		if err := item.fn(ctx); err != nil {
			e = multierr.Append(e, err)
			c.logger.Error(fmt.Sprintf("closer.Close: failed to close %s", item.name), "error", err)
		}
	}
	c.logger.Info("closer.Close: closing finished", "result", e)

	return e
}
