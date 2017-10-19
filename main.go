package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Go REST"
	app.Usage = "A stateless REST API for your relational database"

	app.Commands = []cli.Command{{
		Name:    "run",
		Aliases: []string{"start"},
		Usage:   "start the server",
		Action:  run_server,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "host",
				Value: "127.0.0.1",
				Usage: "hostname of the database to connect to",
			},
			cli.StringFlag{
				Name:  "port",
				Value: "5432",
				Usage: "port of the database to connect to",
			},
			cli.StringFlag{
				Name:  "user",
				Value: "",
				Usage: "username to connect with",
			},
			cli.StringFlag{
				Name:  "password",
				Value: "",
				Usage: "password to connect with",
			},
			cli.StringFlag{
				Name:  "database",
				Value: "postgres",
				Usage: "database to connect with",
			},
			cli.StringFlag{
				Name:  "schema",
				Value: "public",
				Usage: "schema to query against",
			},
		},
	}}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err.Error())
		os.Exit(1)
	}
}

func run_server(ctx *cli.Context) error {
	database := postgres_db{
		Host:     ctx.String("host"),
		Port:     ctx.String("port"),
		Username: ctx.String("user"),
		Password: ctx.String("password"),
		Database: ctx.String("database"),
		Schema:   ctx.String("schema"),
	}
	results, err := describe_tables(database)
	fmt.Printf("%v\n", results)
	return err
}
