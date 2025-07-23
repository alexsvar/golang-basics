package main

import "fmt"

func main() {
	m := map[string]string{
		"Github": "https://github.com",
		"Google": "https://google.com",
		"Yandex": "yandex",
	}
	fmt.Println(m)
	fmt.Println(m["Github"])

	m["Yandex"] = "https://yandex.ru"
	fmt.Println(m["Yandex"])
	
	m["VK"] = "https://vk.com"
	fmt.Println(m["VK"])
	
	delete(m, "VK")
	fmt.Println(m)


	for key, value := range m {
		fmt.Println(key, value)
	}
}
