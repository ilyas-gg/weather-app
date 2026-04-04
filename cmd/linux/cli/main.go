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
l := logger.New()
wi := getProvider(c, l)
app := cli.New(l, wi, c)
err = app.Run()
if err != nil {
l.Error("Some error", err)
os.Exit(1)
}
os.Exit(0)
}
func getProvider(c config.Config, l cli.Logger) cli.WeatherInfo {
var wi cli.WeatherInfo
switch c.P.Type {
case "open-meteo":
wi = weather.New(l)
default:
wi = weather.New(l)
}
return wi
}
r, err := os.Open("./config/config.yaml")
if err != nil {
panic(err)
}