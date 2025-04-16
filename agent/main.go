package main

import (
	"agent/clicommand"
	"fmt"
	"github.com/urfave/cli"
	"os"
)

const appHelpTemplate = `Usage:

  {{.Name}} <command> [options...]

Available commands are:

  {{range .Commands}}{{.Name}}{{with .ShortName}}, {{.}}{{end}}{{ "\t" }}{{.Usage}}
  {{end}}
Use "{{.Name}} <command> --help" for more information about a command.
`

const subcommandHelpTemplate = `Usage:

  {{.Name}} {{if .VisibleFlags}}<command>{{end}} [options...]

Available commands are:

  {{range .Commands}}{{.Name}}{{with .ShortName}}, {{.}}{{end}}{{ "\t" }}{{.Usage}}
  {{end}}{{if .VisibleFlags}}

Options:

{{range .VisibleFlags}}  {{.}}
{{end}}{{ end -}}
`

const commandHelpTemplate = `{{.Description}}

Options:

{{range .VisibleFlags}}  {{.}}
{{ end -}}
`

var version = "1.0.0"

func printVersion(c *cli.Context) {
	fmt.Fprintf(c.App.Writer, "%s version %s\n", c.App.Name, version)
}

func main() {

	cli.AppHelpTemplate = appHelpTemplate
	cli.CommandHelpTemplate = commandHelpTemplate
	cli.SubcommandHelpTemplate = subcommandHelpTemplate
	cli.VersionPrinter = printVersion

	app := cli.NewApp()
	app.Name = "owl-agent"
	app.Version = version
	app.Commands = clicommand.AgentCommands
	app.ErrWriter = os.Stderr

	// When no sub command is used
	app.Action = func(c *cli.Context) {
		_ = cli.ShowAppHelp(c)
		os.Exit(1)
	}

	// When a sub command can't be found
	app.CommandNotFound = func(c *cli.Context, command string) {
		_ = cli.ShowAppHelp(c)
		os.Exit(1)
	}

	if err := app.Run(os.Args); err != nil {
		os.Exit(clicommand.PrintMessageAndReturnExitCode(err))
	}
}
