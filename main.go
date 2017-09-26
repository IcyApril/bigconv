package main

import (
	"github.com/urfave/cli"
	"fmt"
	"math/big"
	"os"
	"github.com/asaskevich/govalidator"
)

func main() {
	app := cli.NewApp()
	app.Name = "BigConv"
	app.Usage = "Convert big hexadecimal values to integers"
	app.Version = "1.0.0"

	app.Commands = []cli.Command{
		{
			Name:    "hex",
			Usage:   "convert hexadecimal to integer",
			Action: func(c *cli.Context) error {
				i := new(big.Int)
				num := c.Args().Get(0)
				_, err := fmt.Sscan("0x"+num, i)
				if err != nil {
					return cli.NewExitError("error scanning value:"+err.Error(), 1)
				}
				fmt.Println(i)
				return nil
			},
		},
		{
			Name:    "int",
			Usage:   "convert integer to hexadecimal",
			Action: func(c *cli.Context) error {
				num := c.Args().Get(0)
				if govalidator.IsInt(num) == false {
					return cli.NewExitError("Argument is not an integer", 1)
				}
				i := new(big.Int)
				_, err := fmt.Sscan(num, i)
				if err != nil {
					return cli.NewExitError("error scanning value:"+err.Error(), 1)
				}
				fmt.Println(fmt.Sprintf("%x", i))
				return nil
			},
		},
	}

	app.Run(os.Args)
}
