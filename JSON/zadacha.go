/*https://stepik.org/lesson/776129/step/4?unit=778575*/

package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Duration описывает продолжительность фильма
type Duration time.Duration

/*Пример кастомного маршаллинга СТРУКТУРЫ Duration*/
func (d Duration) MarshalJSON() ([]byte, error) {
	duration := time.Duration(d)
	// fmt.Println("Time is ", duration.Hours())
	hours := duration / time.Hour
	minutes := (duration % time.Hour) / time.Minute

	var b strings.Builder
	b.Grow(5)
	b.WriteString(`"`)
	if hours != 0 {
		b.WriteString(strconv.Itoa(int(hours)))
		b.WriteByte('h')
	}

	if minutes != 0 {
		b.WriteString(strconv.Itoa(int(minutes)))
		b.WriteByte('m')
	}
	b.WriteString(`"`)
	res := []byte(b.String())
	return res, nil
}

type Rating int

/*Пример кастомной СТРУКТУРЫ Rating. НА выходе звезды*/
func (r Rating) MarshalJSON() ([]byte, error) {
	const maxRate = 5
	stars := strings.Repeat("★", int(r))
	empty := strings.Repeat("☆", maxRate-int(r))
	return []byte(fmt.Sprintf(`"%s%s"`, stars, empty)), nil
}

// Movie описывает фильм
type Movie struct {
	Title    string   //`json:"Title"`
	Year     int      //`json:"Year"`
	Director string   //`json:"Director"`
	Genres   []string //`json:"Genres"`
	Duration Duration
	Rating   Rating //`json:"Rating"`
}

// MarshalMovies кодирует фильмы в JSON.
//   - если indent = 0 - использует json.Marshal
//   - если indent > 0 - использует json.MarshalIndent
//     с отступом в указанное количество пробелов.
func MarshalMovies(indent int, movies ...Movie) (string, error) {
	var b []byte
	var err error

	if indent == 0 {
		b, err = json.Marshal(movies)
	} else if indent > 0 {
		b, err = json.MarshalIndent(movies, "", strings.Repeat(" ", indent))
	}

	if err != nil {
		return "", err
	}
	return string(b), nil
}

func stub_json_zadacha() {
	m1 := Movie{
		Title:    "Interstellar",
		Year:     2014,
		Director: "Christopher Nolan",
		Genres:   []string{"Adventure", "Drama", "Science Fiction"},
		Duration: Duration(2*time.Hour + 49*time.Minute),
		Rating:   5,
	}
	m2 := Movie{
		Title:    "Sully",
		Year:     2016,
		Director: "Clint Eastwood",
		Genres:   []string{"Drama", "History"},
		Duration: Duration(time.Hour + 36*time.Minute),
		Rating:   4,
	}

	s, err := MarshalMovies(4, m1, m2)
	fmt.Println(err)
	// nil
	fmt.Println(s)
	/*
		[
		    {
		        "Title": "Interstellar",
		        "Year": 2014,
		        "Director": "Christopher Nolan",
		        "Genres": [
		            "Adventure",
		            "Drama",
		            "Science Fiction"
		        ],
		        "Duration": "2h49m",
		        "Rating": "★★★★★"
		    },
		    {
		        "Title": "Sully",
		        "Year": 2016,
		        "Director": "Clint Eastwood",
		        "Genres": [
		            "Drama",
		            "History"
		        ],
		        "Duration": "1h36m",
		        "Rating": "★★★★☆"
		    }
		]
	*/
}
