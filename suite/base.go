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

func WithBaseTestSuite(t *testing.T, fn func(s *IntegrationSuite)) {
	s := NewIntegrationSuite(t)
	defer s.Cleanup()
	fn(s)
}
