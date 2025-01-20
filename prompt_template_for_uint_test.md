# 为方法生成单元测试

## 需求

1. 生成一个单元测试函数，名称为 `Test<接口名称>_<方法名称>`。例如：TestProductRepository_CreateByCost
2. 单元测试包含：
   - 初始化必要的存储库或依赖项。例如：repo := repository.NewProductRepository()
   - 定义测试用例结构体，用于描述测试的输入和期望输出。
   - 遍历测试用例，逐一执行，并根据期望输出检查结果。
3. 请根据下列模板进行生成单元测试代码：

### 模板

```go
func Test<接口名称>_<方法名称>(t *testing.T) {
 // 初始化存储库或依赖项
 repo := <初始化代码>

 // 定义测试用例结构体
 type testCase struct {
  name        string // 测试用例的名称
  wantErr     bool   // 是否期望返回错误
  purpose     string // 测试的目的
  expected    string // 对测试结果的期望描述
  <input参数> // 根据实际功能定义输入参数
  <output参数> // 定义预期的输出参数或结果
 }

 // 初始化测试用例
 testCases := []testCase{
  // 使用者可在此定义测试用例
  {
   name:        "<测试用例名称>",
   wantErr:     <是否期望错误>,
   purpose:     "<测试目的描述>",
   expected:    "<期望描述>",
   <input参数>: <具体值>,
   <output参数>: <预期值>,
  },
 }

 // 遍历测试用例，执行测试逻辑
 for _, tc := range testCases {
  t.Run(tc.name, func(t *testing.T) {
   // 调用待测方法
   result, err := repo.<方法名称>(<输入参数>)

   // 根据 wantErr 判断是否正确处理错误
   if tc.wantErr {
    if err == nil {
     t.Fatalf("Expected an error for test case '%s', but got nil, the test case purpose is %s expected %s", tc.name, tc.purpose, tc.expected)
    }
   } else {
    if err != nil {
     t.Fatalf("Unexpected error for test case '%s': %v, the test case purpose is %s expected %s", tc.name, err, tc.purpose, tc.expected)
    }

    // 校验返回结果（具体逻辑由使用者自行定义）
    if result != tc.<output参数> {
     t.Fatalf("For test case '%s', expected %v but got %v, the test case purpose is %s expected %s", tc.name, tc.<output参数>, result, tc.purpose, tc.expected)
    }
   }
  })
 }
}
