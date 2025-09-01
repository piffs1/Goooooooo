package main

import (
	"encoding/json"
	"fmt"
)

/*декодируем стуктуру со вложенными структурами*/
func unmarshall_structure() {
	/*Предположим есть json-чик*/
	const src = `{
    "name": "Alice",
    "is_awesome": true,
    "residence": {
        "country": "France",
        "city": "Paris"
    },
    "friends": [
        { "name": "Emma" },
        { "name": "Grace" }
    ]
}`
	fmt.Println(json.Valid([]byte(src))) //true
	/*Создаем структуру для ВЫЧИТКИ*/
	type Address struct {
		Country string
		City    string
	}

	type Person struct {
		Name      string
		IsAwesome bool `json:"is_awesome"`
		Residence Address
		Friends   []*Person
	}
	/*ДЕКОДИРУЕМ*/
	var alice Person
	err := json.Unmarshal([]byte(src), &alice)

	fmt.Println(err, alice)
	// <nil> {Alice true {France Paris} [0xc00012a050 0xc00012a0a0]}

	fmt.Println(alice.Friends[0])
	// &{Emma false { } []}

	fmt.Println(alice.Friends[1])
	// &{Grace false { } []}
}

type AncientNumber string

func (an *AncientNumber) UnmarshalJSON(data []byte) error {
	// Go рекомендует игнорировать значения null
	if string(data) == "null" {
		return nil
	}
	// декодируем исходное число
	var n int
	err := json.Unmarshal(data, &n)
	if err != nil {
		return err
	}
	// преобразуем в значение типа AncientNumber
	switch {
	case n <= 0:
		*an = "impossible!"
	case n == 1:
		*an = "one"
	case n == 2:
		*an = "two"
	case n > 2:
		*an = "many"
	}
	return nil
}

func custom_unmarshalling() {
	var n AncientNumber

	err := json.Unmarshal([]byte("1"), &n)
	fmt.Println(err, n)
	// <nil> one

	err = json.Unmarshal([]byte("2"), &n)
	fmt.Println(err, n)
	// <nil> two

	err = json.Unmarshal([]byte("42"), &n)
	fmt.Println(err, n)
	// <nil> many

	err = json.Unmarshal([]byte("-1"), &n)
	fmt.Println(err, n)
	// <nil> impossible!
}
