package cmd

import (
	"github.com/sahasajib/mini_atm/config"
	"github.com/sahasajib/mini_atm/rest"
)

func Serve() {
	cnf := config.GetConfig()
	rest.Start(cnf)
}
