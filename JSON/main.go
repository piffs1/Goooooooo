package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// Демонстрация работы маршалинга json
func JSON_work() {
	/*Кодирование в JSON. Опишем тип Person*/
	type Person struct {
		Name      string
		Age       int
		Weight    float64
		IsAwesome bool
		secret    string //<! Поля с маленькой буквой игнорируются.
	}

	/*Содзаем и заполняем структуру*/
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

/*Определяемые типы json*/
func json_time() {
	type Person struct {
		Name      string
		BirthDate time.Time
	}
	date, _ := time.Parse("2006-01-02", "2000-05-25")
	alice := Person{
		Name:      "Alice",
		BirthDate: date,
	}

	b, err := json.Marshal(alice)
	fmt.Println(err, string(b)) //<! <nil> {"Name":"Alice","BirthDate":"2000-05-25T00:00:00Z"}
	/*
		тип time.Time закодировался в строку по стандарту RFC 3339 (ISO 8601).
		time.Time можем замаршаллить, потому что он реализует интерфейс

		type Marshaler interface {
			MarshalJSON() ([]byte, error)
		}


		func (t Time) MarshalJSON() ([]byte, error) {
			b := make([]byte, 0, len(RFC3339Nano)+len(`""`))
			b = append(b, '"')
			b, err := t.appendStrictRFC3339(b)
			b = append(b, '"')
			if err != nil {
				return nil, errors.New("Time.MarshalJSON: " + err.Error())
			}
			return b, nil
		}

		@attention: Значение времени кодируется в строку ОБЯЗАТЕЛЬНО С КАВЫЧКАМИ!
		"BirthDate":"2000-05-25T00:00:00Z"
	*/
}

func json_maps_slices() {
	nums := []int{1, 3, 5}
	b, err := json.Marshal(nums)
	fmt.Println(err, string(b)) ///<! <nil> [1,3,5]

	m := map[string]int{
		"one":   1,
		"three": 3,
		"five":  5,
	}
	b, err = json.Marshal(m)
	fmt.Printf("Err = %v, Marshalled map: %s", err, string(b))
	///<! Err = <nil>, Marshalled map: {"five":5,"one":1,"three":3}

}

func main() {
	JSON_work()
	json_time()
	json_maps_slices()
}
