package cli

import (
	"github.com/m-mizutani/alertchain/pkg/controller/cli/flag"
	"github.com/m-mizutani/alertchain/pkg/domain/model"
	"github.com/m-mizutani/alertchain/pkg/domain/types"
	"github.com/m-mizutani/alertchain/pkg/utils"
	"github.com/m-mizutani/goerr"
	"github.com/urfave/cli/v2"
	"golang.org/x/exp/slog"
)

type CLI struct {
	app *cli.App
}

func New() *CLI {
	var (
		logLevel  flag.LogLevel
		logFormat flag.LogFormat
		logOutput flag.LogOutput

		cfg model.Config
	)

	app := &cli.App{
		Name: "alertchain",
		Flags: []cli.Flag{
			&cli.GenericFlag{
				Name:        "log-level",
				Category:    "logging",
				Aliases:     []string{"l"},
				EnvVars:     []string{"ALERTCHAIN_LOG_LEVEL"},
				Usage:       "Set log level [debug|info|warn|error]",
				Destination: &logLevel,
			},
			&cli.GenericFlag{
				Name:        "log-format",
				Category:    "logging",
				Aliases:     []string{"f"},
				EnvVars:     []string{"ALERTCHAIN_LOG_FORMAT"},
				Usage:       "Set log format [console|json]",
				Destination: &logFormat,
			},
			&cli.GenericFlag{
				Name:        "log-output",
				Category:    "logging",
				Aliases:     []string{"o"},
				EnvVars:     []string{"ALERTCHAIN_LOG_OUTPUT"},
				Usage:       "Set log output (create file other than '-', 'stdout', 'stderr')",
				Destination: &logOutput,
			},
		},

		Before: func(ctx *cli.Context) error {
			utils.ReconfigureLogger(
				logOutput.Writer(),
				logLevel.Level(),
				logFormat.Format(),
			)

			utils.Logger().Debug("config loaded", slog.Any("config", cfg))

			return nil
		},

		Commands: []*cli.Command{
			cmdServe(&cfg),
			cmdRun(&cfg),
			cmdPlay(&cfg),
			{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "Show version",
				Action: func(c *cli.Context) error {
					println("alertchain", types.AppVersion)
					return nil
				},
			},
		},
	}
	return &CLI{app: app}
}

func (x *CLI) Run(argv []string) error {
	if err := x.app.Run(argv); err != nil {
		attrs := []any{
			slog.String("error", err.Error()),
		}

		if goErr := goerr.Unwrap(err); goErr != nil {
			for k, v := range goErr.Values() {
				attrs = append(attrs, slog.Any(k, v))
			}
		}

		utils.Logger().Error("cli failed", attrs...)
		return err
	}

	return nil
}
