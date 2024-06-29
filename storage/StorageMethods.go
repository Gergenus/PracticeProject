package storage

import (
	"database/sql"
	"errors"
	"strconv"
)

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
