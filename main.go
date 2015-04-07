package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/tonyhb/gogdb/gui"
)

const LOG_FILENAME = "/tmp/gogdb.log"

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	app := cli.NewApp()
	app.Name = "gogdb"
	app.Usage = "debug go apps using devtools"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "headless, x",
			Usage: "Run without a UI",
		},
		cli.BoolFlag{
			Name:  "server",
			Usage: "Run the app's UI server without node-webkit",
		},
	}
	app.Action = run
	app.Run(os.Args)
}

func run(c *cli.Context) {
	if c.Bool("headless") {
		logrus.Println("Running headless")
	} else {
		logToFile()
		gui.Run(c)
	}

	// ... init gdb
}

func logToFile() {
	f, err := os.OpenFile(LOG_FILENAME, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	logrus.SetOutput(f)
}
