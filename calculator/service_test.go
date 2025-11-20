package calculator

import (
	"reflect"
	"testing"
)

func TestEdgeCase(t *testing.T) {
	repo := NewMemoryRepo([]int{23, 31, 53})
	svc := NewService(repo)

	amount := 500000

	res, err := svc.CalculatePacks(amount)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := map[int]int{23: 2, 31: 7, 53: 9429}

	if !reflect.DeepEqual(res.PackCounts, want) {
		t.Fatalf("unexpected pack distribution:\n got=%v\nwant=%v", res.PackCounts, want)
	}
}

func TestCalculatePacks_InvalidAmount(t *testing.T) {
	repo := NewMemoryRepo([]int{23, 31, 53})
	svc := NewService(repo)

	_, err := svc.CalculatePacks(-10)
	if err == nil {
		t.Fatal("expected error for negative amount, got nil")
	}
}
