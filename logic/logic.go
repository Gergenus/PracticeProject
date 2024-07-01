package logic

import (
	"errors"
	"strconv"
	"strings"
)

var (
	TooManyPages = errors.New("Слишком большое количество вакансий")
)
var template string = "https://api.hh.ru/vacancies?"

func AddText(text string) {
	ans := make([]string, len(text)) //Результат ручного тестирования показал недостаток в формировании запроса с последующим фиксом
	for _, d := range strings.Split(text, "") {
		if d != " " {
			ans = append(ans, d)
		}
	}
	ans1 := strings.Join(ans, "")
	tmpl := template + "text=" + ans1
	template = tmpl
}

func ReturnFinal() string {
	return template
}

func Addpages(pages int) error {
	if pages > 50 {
		return TooManyPages
	}
	tmpl := template + "&per_page=" + strconv.Itoa(pages)
	template = tmpl
	return nil
}

func AddCity(city string) {
	var ans []string
	z := strings.Split(city, "")
	for _, d := range z {
		if d == " " {
			continue
		}
		ans = append(ans, d)
	}
	city = strings.Join(ans, "")
	template = template + city
}

func Reset() {
	template = "https://api.hh.ru/vacancies?"
}
