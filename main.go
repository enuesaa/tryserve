package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/enuesaa/tryserve/pkg/repository"
	"github.com/enuesaa/tryserve/pkg/usecase"
	"github.com/urfave/cli/v2"
)

func init() {
	log.SetFlags(0)
}

func main() {
	app := &cli.App{
		Name:      "tryserve",
		Version:   "0.0.1",
		Usage:     "Instance web server",
		Args:      true,
		ArgsUsage: "<path>",
		Action: func(c *cli.Context) error {
			path := c.Args().Get(0)
			if path == "" {
				return fmt.Errorf("Argument <path> is required. Please specify the path to serve, like `tryserve .`")
			}

			repos := repository.New()
			ext := filepath.Ext(path)
			// TODO: change this logic. if path is file, run app.
			if ext == "" {
				return usecase.Serve(repos, path)
			}
			return usecase.RunApp(repos, path)
		},
	}

	// disable default
	app.HideHelpCommand = true
	cli.AppHelpTemplate = `{{.Usage}}

USAGE:
	{{.HelpName}}{{if .Commands}} command [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}} {{if .VisibleFlags}}[global options]{{end}}
	{{if len .Authors}}
AUTHOR:
	{{range .Authors}}{{ . }}{{end}}
	{{end}}{{if .Commands}}
COMMANDS:
{{range .Commands}}{{if not .HideHelp}}   {{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}{{if .VisibleFlags}}
GLOBAL OPTIONS:
	{{range .VisibleFlags}}{{.}}
	{{end}}{{end}}
`

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
