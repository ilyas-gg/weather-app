package main

import (
	"os"

	"github.com/ilyas-gg/weather-app/internal/adapters/weather"
	"github.com/ilyas-gg/weather-app/internal/pkg/app/cli"
	"github.com/ilyas-gg/weather-app/internal/pkg/config"
	"github.com/ilyas-gg/weather-app/pkg/logger"
)

func main() {
	r, err := os.Open("./config/config.yaml")
	if err != nil {
		panic(err)
	}
	c, err := config.Parse(r)
	if err != nil {
		panic(err)
	}
	l := logger.New(true) // или logger.New() если без аргументов, смотрите ваш логгер
	_ = getProvider(c, l) // используем blank identifier, чтобы убрать warning (если wi не нужен)
	app := cli.New(l)     // предполагаем, что New принимает только логгер
	err = app.Run()
	if err != nil {
		l.Error(err.Error()) // исправлено: передаём строку
		os.Exit(1)
	}
	os.Exit(0)
}

func getProvider(c config.Config, l cli.Logger) interface{} {
	var wi interface{}
	switch c.P.Type {
	case "open-meteo":
		wi = weather.New(l)
	default:
		wi = weather.New(l)
	}
	return wi
}
