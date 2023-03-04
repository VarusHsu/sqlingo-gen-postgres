package generator

import (
	"flag"
	"fmt"
)

type option struct {
	User       string
	Password   string
	Dbname     string
	Host       string
	Port       string
	TableNames []string
	ForceCases []string
}

var opt option

func init() {
	flag.StringVar(&opt.User, "user", "postgres", "postgres user")
	flag.StringVar(&opt.Password, "password", "", "postgres password")
	flag.StringVar(&opt.Host, "host", "localhost", "postgres host")
	flag.StringVar(&opt.Port, "port", "5432", "postgres port")
	flag.StringVar(&opt.Dbname, "db", "postgres", "postgres database name")
}

func ParseArgs() string {
	flag.Parse()
	dataSourceName := fmt.Sprintf(
		"host=%s sslmode=disable user=%s dbname=%s password=%s port=%s",
		opt.Host,
		opt.User,
		opt.Dbname,
		opt.Password,
		opt.Port,
	)
	return dataSourceName
}
