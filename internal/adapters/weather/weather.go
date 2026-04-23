package weather

import "github.com/ilyas-gg/weather-app/internal/pkg/app/cli"

type WeatherAdapter struct{}

func New(l cli.Logger) *WeatherAdapter {
	return &WeatherAdapter{}
}
