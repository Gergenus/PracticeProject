package logic

import (
	"strconv"
	"strings"
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

func Addpages(pages int) {
	tmpl := template + "&per_page=" + strconv.Itoa(pages)
	template = tmpl
}

func AddCity(city string) {
	template = template + city
}

func Reset() {
	template = "https://api.hh.ru/vacancies?"
}
