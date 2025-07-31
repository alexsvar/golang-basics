package output

import "fmt"

func PrintError(value any) {
	switch t := value.(type) {
	case string:
		fmt.Println(t)
	case int:
		fmt.Printf("Код ошибки: %d", t)
	case error:
		fmt.Println(t.Error())
	default:
		fmt.Println("Неизвестный тип ошибки")
	}
}
