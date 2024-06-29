package telegrambot

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Gergenus/logic"
	"github.com/Gergenus/requests"
	"github.com/Gergenus/storage"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Pizdec struct {
	enum int // 0 рабочее состояние, 1 ждем
}

func changeTo(p *Pizdec, code int) {
	p.enum = code
}

func InitBot() error {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("bot"))
	if err != nil {
		return err
	}
	bot.Debug = true
	log.Println(bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30
	updates := bot.GetUpdatesChan(u)
	Suffer := Pizdec{0}
	db, err := storage.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	for update := range updates {
		if update.Message == nil {
			continue
		}
		switch update.Message.Text {
		case "/start":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Пора за работу\n Чтобы пользоваться ботом используйте последовательно команды:\n /descjob0 \n /descjob1 \n /descjob2 \n после них них используйте /findjob")
			_, err := bot.Send(msg)
			if err != nil {
				return err
			}
			_, err = db.AddUserToDB(int(update.Message.Chat.ID))
			if err != nil {
				return err
			}
		case "/descjob0":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Кем хочешь работать? (Введите название работы)")
			_, err := bot.Send(msg)
			if err != nil {
				return err
			}
			changeTo(&Suffer, 1)
			fmt.Println(Suffer)
		case "/descjob2":
			msg1 := tgbotapi.NewMessage(update.Message.Chat.ID, "Количество вакансий? (Введите количество предложений одним числом)")
			_, err = bot.Send(msg1)
			if err != nil {
				log.Fatal("quantity error", err)
			}
			changeTo(&Suffer, 2)
		case "/descjob1":
			msg1 := tgbotapi.NewMessage(update.Message.Chat.ID, "Город (Введите название города)")
			_, err = bot.Send(msg1)
			if err != nil {
				log.Fatal("quantity error", err)
			}
			changeTo(&Suffer, 3)
		case "/findjob":
			access := os.Getenv("access")
			ans, err := requests.NewRequest(access, logic.ReturnFinal())
			if err != nil {
				return err
			}
			for _, data := range ans.Items {
				vacancy := fmt.Sprintf("Вакансия: %s, место работы: город-%s улица-%s строение-%s, зарплата %d, почта работодателя: %s, имя работодателя: %s", data.Name, data.Area.NameArea, data.Address.Street, data.Address.Street, data.Salary.From, data.Contacts.Email, data.Contacts.Name)
				msg2 := tgbotapi.NewMessage(update.Message.Chat.ID, vacancy)
				_, err = bot.Send(msg2)
				if err != nil {
					return err
				}
				logic.Reset()
				err := db.AddVacancyToDB(data.Name, data.Contacts.Email, strconv.Itoa(int(update.Message.Chat.ID)), data.Area.NameArea, data.Salary.From)
				if err != nil {
					return err
				}
			}
		default:
			if Suffer.enum == 1 {
				logic.AddText(update.Message.Text)
				changeTo(&Suffer, 0)
			}
			if Suffer.enum == 2 {
				chr, err := strconv.Atoi(update.Message.Text)
				if err != nil {
					return err
				}
				logic.Addpages(chr)
				changeTo(&Suffer, 0)
			}
			if Suffer.enum == 3 {
				logic.AddCity(update.Message.Text)
				changeTo(&Suffer, 0)
			}
		}

	}
	return nil
}
