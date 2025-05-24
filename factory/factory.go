package factory

// Factory defines an abstract interface for generating and optionally persisting test entities.
type Factory[T any] interface {
	// Make returns a new instance of T with optional field overrides (not persisted).
	Make(attrs map[string]any) *T

	// MakeMany returns a slice of T instances (not persisted).
	MakeMany(count int, baseAttrs map[string]any) []*T

	// Create returns a new instance of T and persists it (e.g., via repository).
	Create(attrs map[string]any) *T

	// CreateMany returns a slice of T instances and persists them.
	// Implementations should optimize for batch inserts where possible.
	CreateMany(count int, baseAttrs map[string]any) []*T
}

// GenerateMany is a reusable helper to simplify bulk object generation (e.g., for MakeMany or pre-Create).
func GenerateMany[T any](count int, fn func(i int) *T) []*T {
	out := make([]*T, count)
	for i := 0; i < count; i++ {
		out[i] = fn(i)
	}
	return out
}
