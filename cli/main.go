package main

import (
    "fmt"
    "log"
    "os"

    "github.com/urfave/cli/v2"
)

func main() {
    app := &cli.App{
        Name:    "example",
        Usage:   "A simple example CLI application",
        Version: "1.0.0",
        Commands: []*cli.Command{
            {
                Name:    "greet",
                Aliases: []string{"g"},
                Usage:   "Prints a greeting",
                Flags: []cli.Flag{
                    &cli.StringFlag{
                        Name:  "name",
                        Aliases: []string{"n"},
                        Usage: "Name of the person to greet",
                        Value: "World",
                    },
                },
                Action: func(c *cli.Context) error {
                    name := c.String("name")
                    fmt.Printf("Hello, %s!\n", name)
                    return nil
                },
            },
        },
    }

    err := app.Run(os.Args)
    if err != nil {
        log.Fatal(err)
    }
}