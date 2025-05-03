package suite

import (
	"context"
	"testing"
)

type BaseSuite struct {
	T   *testing.T
	Ctx context.Context
}

func NewBaseSuite(t *testing.T) *BaseSuite {
	return &BaseSuite{
		T:   t,
		Ctx: context.Background(),
	}
}
