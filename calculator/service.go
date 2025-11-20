package calculator

import (
	"errors"
	"fmt"
	"math"
	"sort"
)

const maxAmount = 5_000_000

type Result struct {
	PackCounts map[int]int `json:"packs"`
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetPackSizes() []int {
	return s.repo.GetPackSizes()
}

func (s *Service) UpdatePackSizes(sizes []int) error {
	err := s.repo.SavePackSizes(sizes)
	return err
}

func (s *Service) CalculatePacks(amount int) (*Result, error) {
	sizes := s.repo.GetPackSizes()
	return calculateUsingSizes(amount, sizes)
}

// calculateUsingSizes finds the optimal way to reach (or exceed) the requested
// amount using the available pack sizes. It chooses the combination with:
//
//  1. The smallest over-delivery (sum - amount)
//  2. If multiple combinations have equal over-delivery, the one
//     with the fewest number of packs.
//
// This function uses a classic dynamic programming approach similar to
// the "unbounded knapsack" or "coin change (min coins)" problem.
func calculateUsingSizes(amount int, sizes []int) (*Result, error) {
	// Basic validation
	if amount <= 0 {
		return nil, errors.New("amount must be positive")
	}
	if len(sizes) == 0 {
		return nil, errors.New("pack sizes missing")
	}
	if amount > maxAmount {
		return nil, fmt.Errorf("amount too large (max supported: %d)", maxAmount)
	}

	// Ensure pack sizes are sorted ascending for consistency and reconstruction.
	sort.Ints(sizes)

	// We allow sums up to amount + largestPack.
	// This is enough to find the minimal over-delivery.
	maxSize := sizes[len(sizes)-1]
	limit := amount + maxSize

	// "INF" represents an unreachable state inside DP.
	// Using MaxInt32 is safe and large enough for pack counts.
	const INF = math.MaxInt32

	// dp[s] = minimum number of packs required to reach exact sum 's'.
	// choice[s] = index of the pack size used last to achieve 's'.
	dp := make([]int, limit+1)
	choice := make([]int, limit+1)

	// Initialize DP arrays: everything unreachable except dp[0].
	for i := range dp {
		dp[i] = INF
		choice[i] = -1
	}
	dp[0] = 0

	// Try to build each sum 's' using all available pack sizes.
	for s := 1; s <= limit; s++ {
		for i, pack := range sizes {
			// Check if pack fits and if previous sum is reachable.
			if s >= pack && dp[s-pack] < dp[s] {
				dp[s] = dp[s-pack] + 1
				choice[s] = i
			}
		}
	}

	// Search for the best achievable sum >= requested amount.
	// Criteria:
	//   1. Minimal over-delivery
	//   2. Minimal number of packs (tie-breaker)
	bestSum := -1
	bestExcess := INF
	bestCount := INF

	for s := amount; s <= limit; s++ {
		if dp[s] == INF {
			continue // unreachable sum
		}

		excess := s - amount
		count := dp[s]

		// Choose best candidate based on defined priority.
		if excess < bestExcess || (excess == bestExcess && count < bestCount) {
			bestSum = s
			bestExcess = excess
			bestCount = count
		}
	}

	if bestSum == -1 {
		return nil, errors.New("no combination found")
	}

	// Reconstruct pack usage by walking backwards through `choice`
	// starting from bestSum.
	counts := make(map[int]int)
	for cur := bestSum; cur > 0; {
		i := choice[cur]
		p := sizes[i]
		counts[p]++
		cur -= p
	}

	return &Result{PackCounts: counts}, nil
}
