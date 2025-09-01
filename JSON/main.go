package main

import (
	"encoding/json"
	"fmt"
)

/*Кодирование в JSON. Опишем тип Person*/
type Person struct {
	Name      string
	Age       int
	Weight    float64
	IsAwesome bool
	secret    string //<! Поля с маленькой буквой игнорируются.
}

func main() {
	alice := Person{
		Name:      "Alice",
		Age:       25,
		Weight:    55.5,
		IsAwesome: true,
		secret:    "42",
	}

	b, err := json.Marshal(alice) //<! Возвращается []byte, error
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b)) // {"Name":"Alice","Age":25,"Weight":55.5,"IsAwesome":true}

}
