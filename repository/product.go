package repository

import (
	"errors"
	"main/domain"
	"math"
)

// TDD 第一步：定义接口以及领域模型
type ProductRepository interface {
	CreateByCost(name string, description string, cost float64) (*domain.Product, error)
}

// TDD 第二步：接口的实现，实现内容为空或者使用AI生成的内容
func NewProductRepository() ProductRepository {
	return &InMemoryProductRepository{}
}

type InMemoryProductRepository struct {
	products []domain.Product
}

// TDD 第二步：接口的实现，实现内容为空或者使用AI生成的内容
/*

func (r *InMemoryProductRepository) CreateByCost(name string, description string, cost float64) (*domain.Product, error) {
    price := cost * 1.7 // 利润率为70%
    product := &domain.Product{
        ID:          len(r.products) + 1,
        Name:        name,
        Description: description,
        Price:       price,
    }
    return product, nil
}

*/

// TDD 第四步：接口的实现，实现内容需要通过单元测试
func (r *InMemoryProductRepository) CreateByCost(name string, description string, cost float64) (*domain.Product, error) {
	if cost <= 0 || math.IsNaN(cost) || math.IsInf(cost, 0) {
		return nil, errors.New("cost cannot be zero, negative, NaN, or Infinity")
	}
	price := cost * 1.7 // 利润率为70%
	product := &domain.Product{
		ID:          len(r.products) + 1,
		Name:        name,
		Description: description,
		Price:       price,
	}
	return product, nil
}
