package domain

// TDD 第一步：定义接口以及领域模型
type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
}
