package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
)

type VacancyData struct {
	Vacancy string
	Email   string
	City    string
	Salary  int
}

func (s DBsource) AddVacancyToDB(vacancy, email, telegramId, city string, salary int) error {
	var id int
	err := s.db.QueryRow("Select id from users WHERE Telegram_id=$1", telegramId).Scan(&id)
	if err != nil {
		return err
	}
	_, err = s.db.Exec("INSERT INTO vacancies (vacancy, salary, user_id, email, city) VALUES($1, $2, $3, $4, $5)", vacancy, salary, id, email, city)
	if err != nil {
		return err
	}
	return nil
}

func (s DBsource) AddUserToDB(telegramId int) (int, error) {
	var id int
	svo := strconv.Itoa(telegramId)
	err := s.db.QueryRow("Select id from users WHERE Telegram_id=$1", svo).Scan(&id)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			_, err = s.db.Exec("INSERT INTO users (Telegram_id) VALUES($1)", svo)
			if err != nil {
				return 0, err
			}
			return 0, nil
		}
		return 0, err
	}
	return id, nil
}

func (s DBsource) GetUsersvacancies(telegramId int) ([]VacancyData, error) {
	var foreign_id int
	var vacancy string
	var email string
	var city string
	var salary int
	var ans []VacancyData

	s.db.QueryRow("SELECT id FROM users WHERE telegram_id=$1", strconv.Itoa(telegramId)).Scan(&foreign_id)
	fmt.Println(foreign_id)
	row, err := s.db.Query("SELECT vacancy, salary, email, city FROM vacancies WHERE user_id=$1", foreign_id)
	if err != nil {
		return []VacancyData{}, err
	}
	defer row.Close()
	for row.Next() {
		err = row.Scan(&vacancy, &salary, &email, &city)
		if err != nil {
			return []VacancyData{}, err
		}
		z := VacancyData{Vacancy: vacancy, Email: email, City: city, Salary: salary}
		ans = append(ans, z)
	}
	return ans, err
}
