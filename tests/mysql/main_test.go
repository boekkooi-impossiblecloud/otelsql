package mysql

import (
	"os"
	"testing"

	"github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"                      // Database driver
	_ "github.com/golang-migrate/migrate/v4/database/mysql" // Database driver
	"github.com/nhatthm/testcontainers-go-extra"
	testcontainersmysql "github.com/nhatthm/testcontainers-go-registry/sql/mysql"

	"github.com/nhatthm/otelsql/tests/suite"
)

const (
	defaultVersion = "8"
	defaultImage   = "mysql"
	defaultDriver  = "mysql"

	databaseName     = "otelsql"
	databaseUsername = "otelsql"
	databasePassword = "OneWrapperToTraceThemAll"
)

func TestIntegration(t *testing.T) {
	suite.Run(t,
		suite.WithTestContainerRequests(
			testcontainersmysql.Request(databaseName, databaseUsername, databasePassword,
				testcontainersmysql.RunMigrations("file://./resources/migrations/"),
				testcontainers.WithImageName(imageName()),
				testcontainers.WithImageTag(imageTag()),
			),
		),
		suite.WithDatabaseDriver(defaultDriver),
		suite.WithDatabaseDSN(testcontainersmysql.DSN(databaseName, databaseUsername, databasePassword)),
		suite.WithDatabasePlaceholderFormat(squirrel.Question),
		suite.WithFeatureFilesLocation("../features"),
		suite.WithCustomerRepositoryConstructor(newRepository()),
	)
}

func imageTag() string {
	v := os.Getenv("MYSQL_VERSION")
	if v == "" {
		return defaultVersion
	}

	return v
}

func imageName() string {
	img := os.Getenv("MYSQL_DIST")
	if img == "" {
		return defaultImage
	}

	return img
}
