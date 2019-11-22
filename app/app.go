package app

import (
	"github.com/idobry/dynamik/model"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type App struct {
	Config *Config
	Repository model.Repository
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
	logrus.Info(viper.GetString("giturl"), viper.GetString("username"),viper.GetString("token"))
	app.Repository = model.NewRepository(viper.GetString("giturl"), viper.GetString("username"),viper.GetString("token"))
	return app, err
}