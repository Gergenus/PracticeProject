package requests

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type AreaSTR struct {
	NameArea string `json:"name"`
}

type AddressSTR struct {
	Building string `json:"building"`
	Street   string `json:"street"`
}

type ContactsSTR struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type DepartmentSTR struct {
	Name string `json:"name,omitempty"`
}

type EmployerSTR struct {
	Name string `json:"name"`
}

type SalarySTR struct {
	From int `json:"from"`
}

type Items struct {
	Name       string        `json:"name"`
	Area       AreaSTR       `json:"area"`
	Address    AddressSTR    `json:"address"`
	Contacts   ContactsSTR   `json:"contacts"`
	Department DepartmentSTR `json:"department,omitempty"`
	Employer   EmployerSTR   `json:"employer"`
	Salary     SalarySTR     `json:"salary"`
	Site       string        `json:"url"`
}

type Ans struct {
	Items []Items `json:"items"`
}

func NewRequest(code string, reqstring string) (Ans, error) {
	access := fmt.Sprintf("Bearer %s", code)
	repl := Ans{}
	req, err := http.NewRequest(http.MethodGet, reqstring, nil)
	if err != nil {
		return Ans{}, err
	}
	req.Header.Add("Authorization", access)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return Ans{}, err
	}
	defer res.Body.Close()
	ans, err := io.ReadAll(res.Body)
	if err != nil {
		return Ans{}, err
	}
	err = json.Unmarshal(ans, &repl)
	if err != nil {
		return Ans{}, err
	}
	return repl, nil
}
