package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "magic bean buyer",
		Usage: "Quote Shel Silverstein.",
		Action: func(*cli.Context) error {
			fmt.Println("If you are a dreamer come in\nIf you are a dreamer a wisher a liar\nA hoper a pray-er a magic-bean-buyer\nIf youre a pretender com sit by my fire\nFor we have some flax golden tales to spin\nCome in!\nCome in!")
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
