package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	println(getGreeting(), "\n")
	println(getRandomAnswer())
}

func getGreeting() string {
	var greetings = []string{
		"Привет, дорогой друг. Отвечаю на твой вопрос...",
		"Кто вопрошает, тот получит ответ:",
		"Здравствуй, смертный. Сегодня для тебя такой ответ:",
	}
	return greetings[rand.Intn(len(greetings))]
}

func getRandomAnswer() string {
	var answers = []string{
		"Бесспорно",
		"Предрешено",
		"Никаких сомнений",
		"Определённо да",
		"Можешь быть уверен в этом",
		"Мне кажется — «да»",
		"Вероятнее всего",
		"Хорошие перспективы",
		"Знаки говорят — «да»",
		"Да",
		"Пока не ясно, попробуй снова",
		"Спроси позже",
		"Лучше не рассказывать",
		"Сейчас нельзя предсказать",
		"Сконцентрируйся и спроси опять",
		"Даже не думай",
		"Мой ответ — «нет»",
		"По моим данным — «нет»",
		"Перспективы не очень хорошие",
		"Весьма сомнительно",
	}

	return answers[rand.Intn(len(answers))]
}
