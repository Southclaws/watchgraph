package main

import (
	"bytes"
	"container/ring"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/guptarohit/asciigraph"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
	"github.com/wayneashleyberry/terminal-dimensions"
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.DurationFlag{
			Name:  "interval",
			Value: time.Second,
		},
	}

	app.Action = func(c *cli.Context) (err error) {
		interval := c.Duration("interval")
		ticker := time.NewTicker(interval)

		command := c.Args().First()
		if command == "" {
			return errors.New("Usage: ascigraphwatch [options] command")
		}

		x, err := terminaldimensions.Width()
		if err != nil {
			panic(err)
		}
		y, err := terminaldimensions.Height()
		if err != nil {
			panic(err)
		}
		x -= 14

		data := ring.New(int(x))

		for range ticker.C {
			cmd := exec.Command("bash", "-c", command)
			out, err := cmd.CombinedOutput()
			if err != nil {
				return errors.Wrap(err, "failed to run command")
			}

			value, err := strconv.Atoi(string(bytes.Trim(out, "\n\r\t ")))
			if err != nil {
				return errors.Wrap(err, "failed to parse output as number")
			}

			data.Value = value
			data = data.Next()

			raw := make([]float64, 0, x)
			data.Do(func(v interface{}) {
				asint, ok := v.(int)
				if !ok {
					return
				}
				raw = append(raw, float64(asint))
			})

			graph := asciigraph.Plot(raw, asciigraph.Height(int(y-2)))
			print("\033[H\033[2J")
			fmt.Println(graph)
		}

		return
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
