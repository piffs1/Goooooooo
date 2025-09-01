package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// Демонстрация работы маршалинга json на примере простой структуры
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

/*Маршаллим карты и слайсы*/
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

/*Маршаллим составные значения*/
func json_composite_values() {
	type Address struct {
		Country string
		City    string
	}

	type Person struct {
		Name      string
		Residence Address
	}

	paris := Address{"France", "Paris"}

	alice := Person{
		Name:      "Alice",
		Residence: paris,
	}
	/* MarshalIndent вывод в удобном формате.*/
	b, _ := json.MarshalIndent(alice, "", "    ")
	fmt.Println(string(b))
	/*
			{
		    "Name": "Alice",
		    "Residence": {
		        "Country": "France",
		        "City": "Paris"
		    }
		}
	*/

	b, _ = json.MarshalIndent(alice, "BORIS_", "GRECHA")
	/*
		{
		BORIS_GRECHA"Name": "Alice",
		BORIS_GRECHA"Residence": {
		BORIS_GRECHAGRECHA"Country": "France",
		BORIS_GRECHAGRECHA"City": "Paris"
		BORIS_GRECHA}
		BORIS_}
	*/
	fmt.Println(string(b))
}

/*Маршаллим структуры с указателями*/
func json_pointers() {
	/*Работает и в сложных случаях. Размыеновывает указатель и показывает поля в значении*/
	type Address struct {
		Country string
		City    string
	}

	type Person struct {
		Name    string
		Friends []*Person
	}

	emma := Person{Name: "Emma"}
	grace := Person{Name: "Grace"}

	alice := Person{
		Name:    "Alice",
		Friends: []*Person{&emma, &grace},
	}

	b, _ := json.MarshalIndent(alice, "", "    ")
	fmt.Println(string(b))
	//У Алисы друг Эмма и Грейс. У Эммы и Грейс нет друзей, поэтому Friend = NIULL
	/*
			{
		    "Name": "Alice",
		    "Friends": [
		        {
		            "Name": "Emma",
		            "Friends": null
		        },
		        {
		            "Name": "Grace",
		            "Friends": null
		        }
		    ]
		}*/

}

/*Определяем необходимые поля с помощью json-тегов*/
func json_tags() {
	fmt.Println("=============FUNCTION:JSON_TAGS=============")
	type Person struct {
		Name      string  `json:"namezzz"` //В json будет "namezzz": "Alice"
		Age       int     `json:"age"`
		Weight    float64 `json:"-"` //Заигнорируется
		IsAwesome bool    `json:"is_awesome"`
	}
	alice := Person{
		Name:      "Alice",
		Age:       25,
		Weight:    55.5,
		IsAwesome: true,
	}

	b, err := json.MarshalIndent(alice, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
	/*
		{
		    "namezzz": "Alice",
		    "age": 25,
		    "is_awesome": true
		}
	*/
}

func main() {
	// JSON_work()
	// json_time()
	// json_maps_slices()
	// json_composite_values()
	// json_pointers()
	// stub_json_zadacha()
	// json_tags()
	// custom_unmarshalling()
	выборочные_поля()
}
