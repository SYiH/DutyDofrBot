package main

import (
	"log"
	"time"
	"fmt"

	"gopkg.in/telebot.v3"
)

func main() {
	// Заменить на свой токен от BotFather
	botToken := "tockenBot"

	// Заменить на ID вашей группы (например: -1001234567890)
	groupID := int64(-GID)

	pref := telebot.Settings{
		Token:  botToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Клавиатура и кнопка
	menu := &telebot.ReplyMarkup{}
	btnPost1 := menu.Data("Прием дежурства", "post_btn")
	btnPost2 := menu.Data("Сдача дежурства", "post_btn")
	menu.Inline(menu.Row(btnPost1, btnPost2))

	// Команда /start
	bot.Handle("/start", func(c telebot.Context) error {
		return c.Send("Привет! Нажми кнопку для публикации в группу:", menu)
	})

	// Обработка нажатия кнопки приема дежурства 
	bot.Handle(&btnPost1, func(c telebot.Context) error {
		//Данные о пользователе
		user := c.Sender()
		
		// Формируем ссылку на пользователя в HTML
		userLink := fmt.Sprintf(
			`<a href="tg://user?id=%d">%s</a>`,
			user.ID,
			user.FirstName,
		)
		
		msg := fmt.Sprintf("%s дежурство в СОУ ИТС сдал. Проверки СОУ ИТС по чек-листу выполнены успешно, чек-лист заполнен", userLink)
		// Отправка сообщения в указанную группу
		_, err := bot.Send(&telebot.Chat{ID: groupID}, msg, telebot.ModeHTML,)
		if err != nil {
			log.Println("Ошибка отправки:", err)
			return c.Send("Не удалось отправить в группу 😢")
		}
		return c.Send("Сообщение отправлено в группу ✅")
	})
		// Обработка нажатия кнопки сдачи дежурства
		bot.Handle(&btnPost2, func(c telebot.Context) error {
			//Данные о пользователе
			user := c.Sender()
			
			// Формируем ссылку на пользователя в HTML
			userLink := fmt.Sprintf(
				`<a href="tg://user?id=%d">%s</a>`,
				user.ID,
				user.FirstName,
			)
			
			msg := fmt.Sprintf("%s дежурство в СОУ ИТС принял. Проверки СОУ ИТС по чек-листу выполнены успешно, чек-лист заполнен", userLink)
			// Отправка сообщения в указанную группу
			_, err := bot.Send(&telebot.Chat{ID: groupID}, msg, telebot.ModeHTML,)
			if err != nil {
				log.Println("Ошибка отправки:", err)
				return c.Send("Не удалось отправить в группу 😢")
			}
			return c.Send("Сообщение отправлено в группу ✅")
		})

	bot.Start()
}
