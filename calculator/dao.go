package calculator

import "sync"

type MemoryRepo struct {
	mu    sync.RWMutex
	sizes []int
}

func NewMemoryRepo(initial []int) *MemoryRepo {
	return &MemoryRepo{
		sizes: append([]int(nil), initial...),
	}
}

func (r *MemoryRepo) GetPackSizes() []int {
	r.mu.RLock()
	defer r.mu.RUnlock()

	out := make([]int, len(r.sizes))
	copy(out, r.sizes)
	return out
}

func (r *MemoryRepo) SavePackSizes(sizes []int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.sizes = append([]int(nil), sizes...)
	return nil
}
