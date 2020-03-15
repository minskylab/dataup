package dataup

import (
	"github.com/caarlos0/env/v6"
	"github.com/gofiber/fiber"
)

type DataService struct {
	app *fiber.App
	Port string `env:"PORT" envDefault:"8080"`
	Host string `env:"ASSISTANCE_HOST" envDefault:"assistance"`
}

func NewDataService() (*DataService, error) {
	s := new(DataService)
	if err := env.Parse(s); err != nil {
		return nil, err
	}

	s.app = fiber.New()
	return s, nil
}


func (data *DataService) Run() error {
	return data.app.Listen(data.Port)
}
