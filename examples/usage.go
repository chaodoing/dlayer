//go:build ignore

package main

func main() {
	// 生成代码后，在业务项目中按实际模块路径导入：
	//
	// import "generated/generated/validator"
	//
	// input := validator.UserRequest{
	// 	Name:  "Alice",
	// 	Email: "alice@example.com",
	// 	Age:   18,
	// }
	//
	// if err := validator.Validate(&input, validator.SceneInsert); err != nil {
	// 	return err
	// }
}
