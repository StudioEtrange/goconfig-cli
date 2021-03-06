package main

import (
	"bufio"
	"fmt"
	"github.com/Unknwon/goconfig"
	"os"
	//TODO test : "gopkg.in/ini.v1"
	"gopkg.in/codegangsta/cli.v1"
)

var configFileName string
var configFile *goconfig.ConfigFile
var modified bool = false

func main() {
	prettyFlag := cli.BoolFlag{
		Name:  "pretty, p",
		Usage: "active space around = character",
	}

	// main controller
	app := cli.NewApp()
	app.Name = "goconfig-cli"
	app.Usage = "simple cli to manipulate ini file"
	app.Commands = []cli.Command{
		{
			Name:      "getkey",
			ShortName: "g",
			Usage:     "get key value of a section",
			Before: func(c *cli.Context) error {
				return Init(c.Args().Get(0))
			},
			Action: func(c *cli.Context) {
				val := GetKey(c.Args().Get(2), c.Args().Get(3))
				fmt.Print(val)
			},
		},

		{
			Name:      "delkey",
			ShortName: "d",
			Usage:     "delete a key from a section",
			Flags: []cli.Flag{
				prettyFlag,
			},
			Before: func(c *cli.Context) error {
				return Init(c.Args().Get(0))
			},
			Action: func(c *cli.Context) {
				goconfig.PrettyFormat = c.Bool("pretty")
				DelKey(c.Args().Get(2), c.Args().Get(3))
			},
		},

		{
			Name:      "setkey",
			ShortName: "s",
			Usage:     "add/replace a key with a value into a section",
			Flags: []cli.Flag{
				prettyFlag,
			},
			Before: func(c *cli.Context) error {
				return Init(c.Args().Get(0))
			},
			Action: func(c *cli.Context) {
				goconfig.PrettyFormat = c.Bool("pretty")
				SetKey(c.Args().Get(2), c.Args().Get(3), c.Args().Get(4))
			},
		},
	}

	// launch app
	app.RunAndExitOnError()

	// close app
	if modified {
		if err := goconfig.SaveConfigFile(configFile, configFileName); err != nil {
			os.Exit(1)
		}
	}
	os.Exit(0)
}

func Init(filename string) error {
	var err error

	configFileName = filename

	if lineBreak := DetectLineEnding(filename); lineBreak != "" {
		goconfig.LineBreak = lineBreak
	}

	configFile, err = goconfig.LoadConfigFile(filename)
	return err
}

func DetectLineEnding(filename string) string {
	f, err := os.Open(filename)
	if err != nil {
		return ""
	}
	defer f.Close()

	buf := bufio.NewReader(f)
	_, err = buf.ReadString('\r')
	if err != nil {
		buf.Reset(f)
		_, err = buf.ReadString('\n')
		if err != nil {
			return ""
		} else {
			return "\n"
		}
	} else {
		char, _ := buf.ReadByte()
		if char == '\n' {
			return "\r\n"
		} else {
			return "\r"
		}
	}

}

func GetKey(section, key string) string {
	val, err := configFile.GetValue(section, key)
	if err != nil {
		return ""
	}
	return val
}

func DelKey(section, key string) {
	configFile.DeleteKey(section, key)
	modified = true
}

func SetKey(section, key, val string) {
	configFile.SetValue(section, key, val)
	modified = true
}
