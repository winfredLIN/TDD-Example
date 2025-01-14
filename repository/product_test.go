package repository_test

import (
	"math"
	"testing"

	"main/repository"
)

// TDD 第三步：为接口创建测试，单元测试也是其中的一种
/* Prompt
请根据以下规则，完成repo.CreateByCost方法的单元测试
这是需要测试的方法的相关定义：

CreateByCost(name string, description string, cost float64) (*domain.Product, error)

type Product struct {
    ID          int
    Name        string
    Description string
    Price       float64
}

该函数的目的是根据商品的成本创建商品, 创建的商品应满足：
1. 结果中商品的id需要非空
2. 结果中商品的名称和描述应与输入的一致
3. 结果中商品的价格应该保证使得最终的利润率大于70%
4. 成本为0或其他异常情况，应该返回错误

必须满足：
1. 测试用例需要覆盖尽可能多的场景，需要考虑边界条件
2. 测试用例提取为结构体,结构体应该包含该测试用例测试的目的以及预期的结果，并且在开始进行初始化
3. 在单元测试失败时应立即停止继续测试
4. 单元测试报告的错误必须包含测试用例的目的、预期结果以及测试的参数
*/

func TestProductRepository_CreateByCost(t *testing.T) {
	repo := repository.NewProductRepository()

	// 定义测试用例结构体
	type testCase struct {
		name        string
		description string
		cost        float64
		wantErr     bool
		purpose     string
		expected    string
	}
	// 初始化测试用例
	testCases := []testCase{
		{name: "Test Product", description: "This is a test product", cost: 10.0, wantErr: false, purpose: "Test normal product creation", expected: "Product should be created with correct name, description, and price"},
		{name: "Zero Cost Product", description: "This is a zero cost product", cost: 0.0, wantErr: true, purpose: "Test product creation with zero cost", expected: "An error should be returned indicating that the cost cannot be zero"},
		{name: "Negative Cost Product", description: "This is a negative cost product", cost: -5.0, wantErr: true, purpose: "Test product creation with negative cost", expected: "An error should be returned indicating that the cost cannot be negative"},
		{name: "NaN Cost Product", description: "This is a NaN cost product", cost: math.NaN(), wantErr: true, purpose: "Test product creation with NaN cost", expected: "An error should be returned indicating that the cost is invalid"},
		{name: "Infinity Cost Product", description: "This is an infinity cost product", cost: math.Inf(1), wantErr: true, purpose: "Test product creation with infinity cost", expected: "An error should be returned indicating that the cost is invalid"},
	}

	// 遍历测试用例，执行测试
	for _, tc := range testCases {
		product, err := repo.CreateByCost(tc.name, tc.description, tc.cost)
		if tc.wantErr {
			if err == nil {
				t.Fatalf("Expected an error for test case '%s' with expected '%s', but got nil", tc.name, tc.expected)
			}
		} else {
			if err != nil {
				t.Fatalf("Expected no error for test case '%s' with purpose '%s', but got: %v", tc.name, tc.purpose, err)
			} else {
				if product.ID == 0 {
					t.Fatalf("Expected non-zero product ID for test case '%s' with purpose '%s', but got: %d. Expected: %v", tc.name, tc.purpose, product.ID, tc.expected)
				}
				if product.Name != tc.name {
					t.Fatalf("Expected product name to be '%s' for test case '%s' with purpose '%s', but got: %s. Expected: %v", tc.name, tc.name, tc.purpose, product.Name, tc.expected)
				}
				if product.Description != tc.description {
					t.Fatalf("Expected product description to be '%s' for test case '%s' with purpose '%s', but got: %s. Expected: %v", tc.description, tc.name, tc.purpose, product.Description, tc.expected)
				}
				if math.Abs((product.Price-tc.cost)/tc.cost-0.70) > 0.005 {
					t.Fatalf("Expected price to be greater than %.2f for test case '%s' with purpose '%s', but got: %.2f. The profit margin is only %.2f%%. Expected: %v", tc.cost*1.7, tc.name, tc.purpose, product.Price, (product.Price-tc.cost)/tc.cost*100, tc.expected)
				}
			}
		}
	}
}
