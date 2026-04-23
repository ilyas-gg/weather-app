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
	l := logger.New(true)   // исправлено: добавили bool
	wi := getProvider(c, l) // wi может быть любого типа, который нужен app
	app := cli.New(l)       // исправлено: только один аргумент

	// Если app должен знать wi и c — добавьте их через методы
	// Например (названия методов предположительные):
	// app.SetWeatherInfo(wi)
	// app.SetConfig(c)

	err = app.Run()
	if err != nil {
		l.Error(err) // исправлено: только ошибка
		os.Exit(1)
	}
	os.Exit(0)
}

// Временно убираем конкретный тип возврата, если cli.WeatherInfo не определён
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
