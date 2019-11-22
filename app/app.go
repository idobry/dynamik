package app

import (
	"github.com/idobry/dynamik/model"
	"github.com/sirupsen/logrus"
)

type App struct {
	Config *Config
	Repository *model.Repository
}

func (a *App) NewContext() *Context {
	return &Context{
		Logger: logrus.New(),
	}
}

func New() (app *App,err error) {
	app = &App{}
	app.Config, err = InitConfig()
	if err != nil {
		return nil, err
	}
	app.Repository = model.NewRepository()
	return app, err
}