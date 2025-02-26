package migrations

import (
	"github.com/golang-migrate/migrate/v4"

	"github.com/augustus281/cqrs-pattern/global"
	"github.com/pkg/errors"
)

func RunMigrations() (version uint, dirty bool, err error) {
	if !global.Config.Migrations.Enabled {
		return 0, false, nil
	}

	m, err := migrate.New(global.Config.Migrations.SourceUrl, global.Config.Migrations.DBUrl)
	if err != nil {
		return 0, false, err
	}
	defer func() {
		sourceErr, dbErr := m.Close()
		if sourceErr != nil {
			err = sourceErr
		}
		if dbErr != nil {
			err = dbErr
		}
	}()

	if global.Config.Migrations.Recreate {
		if err = m.Down(); err != nil {
			return 0, false, err
		}
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return 0, false, err
	}

	return m.Version()
}
