package config

import (
	"github.com/m-mizutani/alertchain/pkg/domain/interfaces"
	"github.com/m-mizutani/alertchain/pkg/domain/model"
	"github.com/m-mizutani/alertchain/pkg/domain/types"
	"github.com/m-mizutani/alertchain/pkg/infra/firestore"
	"github.com/m-mizutani/alertchain/pkg/infra/memory"
	"github.com/m-mizutani/alertchain/pkg/utils"
	"github.com/m-mizutani/goerr"
	"github.com/urfave/cli/v2"
)

type Database struct {
	dbType              string
	firestoreProjectID  string
	firestoreDatabaseID string
}

func (x *Database) Flags() []cli.Flag {
	category := "Database"

	return []cli.Flag{
		&cli.StringFlag{
			Name:        "db-type",
			Usage:       "Database type (memory, firestore)",
			Category:    category,
			Aliases:     []string{"t"},
			EnvVars:     []string{"ALERTCHAIN_DB_TYPE"},
			Value:       "memory",
			Destination: &x.dbType,
		},
		&cli.StringFlag{
			Name:        "firestore-project-id",
			Usage:       "Project ID of Firestore",
			Category:    category,
			EnvVars:     []string{"ALERTCHAIN_FIRESTORE_PROJECT_ID"},
			Destination: &x.firestoreProjectID,
		},
		&cli.StringFlag{
			Name:        "firestore-database-id",
			Usage:       "Prefix of Firestore database ID",
			Category:    category,
			EnvVars:     []string{"ALERTCHAIN_FIRESTORE_DATABASE_ID"},
			Destination: &x.firestoreDatabaseID,
		},
	}
}

func (x *Database) New(ctx *model.Context) (interfaces.Database, func(), error) {
	nopCloser := func() {}

	switch x.dbType {
	case "memory":
		return memory.New(), nopCloser, nil

	case "firestore":
		if x.firestoreProjectID == "" {
			return nil, nopCloser, goerr.Wrap(types.ErrInvalidOption, "firestore-project-id is required for firestore")
		}
		if x.firestoreDatabaseID == "" {
			return nil, nopCloser, goerr.Wrap(types.ErrInvalidOption, "firestore-collection-prefix is required for firestore")
		}

		client, err := firestore.New(ctx, x.firestoreProjectID, x.firestoreDatabaseID)
		if err != nil {
			return nil, nopCloser, goerr.Wrap(err, "failed to initialize firestore client")
		}
		return client, func() { utils.SafeClose(client) }, nil

	default:
		return nil, nopCloser, goerr.Wrap(types.ErrInvalidOption, "invalid db-type").With("db-type", x.dbType)
	}
}
