package calculator

type Repository interface {
	GetPackSizes() []int
	SavePackSizes([]int) error
}
