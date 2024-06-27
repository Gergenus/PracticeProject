package telegrambot

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Gergenus/logic"
	"github.com/Gergenus/requests"
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
	for update := range updates {
		if update.Message == nil {
			continue
		}
		switch update.Message.Text {
		case "/start":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Иди работай сука")
			_, err := bot.Send(msg)
			if err != nil {
				log.Fatal("greet error", err)
			}
		case "/descjob0":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Кем хочешь работать? (Введите название работы)")
			_, err := bot.Send(msg)
			if err != nil {
				log.Fatal("greet error", err)
			}
			changeTo(&Suffer, 1)
			fmt.Println(Suffer)
		case "/descjob1":
			msg1 := tgbotapi.NewMessage(update.Message.Chat.ID, "Количество вакансий? (Введите количество предложений одним числом)")
			_, err = bot.Send(msg1)
			if err != nil {
				log.Fatal("quatity error", err)
			}
			changeTo(&Suffer, 2)
		case "/findjob":
			access := os.Getenv("access")
			ans, err := requests.NewRequest(access, logic.ReturnFinal())
			if err != nil {
				log.Fatal(err)
			}
			for _, data := range ans.Items {
				vacancy := fmt.Sprintf("Вакансия: %s, место работы: город-%s улица-%s строение-%s, зарплата %d, почта работодателя: %s, имя работодателя: %s ", data.Name, data.Area.NameArea, data.Address.Street, data.Address.Street, data.Salary.From, data.Contacts.Email, data.Contacts.Name)
				msg2 := tgbotapi.NewMessage(update.Message.Chat.ID, vacancy)
				_, err = bot.Send(msg2)
				if err != nil {
					log.Fatal("Return job error", err)
				}
			}
		case "/reset":
			logic.Reset()
		default:
			if Suffer.enum == 1 {
				logic.AddText(update.Message.Text)
				changeTo(&Suffer, 0)
			}
			if Suffer.enum == 2 {
				chr, err := strconv.Atoi(update.Message.Text)
				if err != nil {
					log.Fatal("Ошибка конвертации", err)
				}
				logic.Addpages(chr)
				changeTo(&Suffer, 0)
			}

		}
	}

	return nil
}
