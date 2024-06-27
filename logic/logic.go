package logic

import "strconv"

var template string = "https://api.hh.ru/vacancies?"

func AddText(text string) string {
	tmpl := template + "text=" + text
	template = tmpl
	return tmpl
}

func ReturnFinal() string {
	return template
}

func Addpages(pages int) string {
	tmpl := template + "&per_page=" + strconv.Itoa(pages)
	template = tmpl
	return tmpl
}

func Reset() {
	template = "https://api.hh.ru/vacancies?"
}
