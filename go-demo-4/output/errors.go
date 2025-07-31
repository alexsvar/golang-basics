package output

import "fmt"

func PrintError(value any) {
	intValue, ok := value.(int)
	if ok {
		fmt.Printf("Код ошибки: %d \n", intValue)
		return
	}
	strValue, ok := value.(string)
	if ok {
		fmt.Println(strValue)
		return
	}
	errorValue, ok := value.(error)
	if ok {
		fmt.Println(errorValue.Error())
		return
	}
	fmt.Println("Неизвестный тип ошибки")

	// switch t := value.(type) {
	// case string:
	// 	fmt.Println(t)
	// case int:
	// 	fmt.Printf("Код ошибки: %d", t)
	// case error:
	// 	fmt.Println(t.Error())
	// default:
	// 	fmt.Println("Неизвестный тип ошибки")
	// }
}

func sum[T int | float32 | float64 | string](a, b T) T {
	return a + b
}
