package command

import "context"

type Sneaker struct{}

type SneakerCommand interface {
	Create(ctx context.Context, sneaker *Sneaker) error
	Update(sneaker *Sneaker) error // want "context.Context is required as the first argument"
}
