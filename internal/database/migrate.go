package database

import (
	"embed"
	"fmt"
	"github.com/fatih/color"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"strings"
	"violettes-cms/utility/helper"

	_ "github.com/golang-migrate/migrate/v4/database/mysql"
)

//go:embed migrations/*.sql
var fs embed.FS

func Migrate() {
	command := gcmd.GetArg(2).String()

	link, err := g.Cfg().Get(gctx.New(), "database.default.link")
	if err != nil || link == nil {
		helper.ConsoleErrorMessageExit("get database link error", err)
	}
	linkStr := link.String()
	linkStr = strings.TrimLeft(linkStr, "mysql:")
	linkStr = "mysql://" + linkStr

	d, err := iofs.New(fs, "migrations")
	if err != nil {
		helper.ConsoleErrorMessageExit("new migrations error", err)
	}

	m, err := migrate.NewWithSourceInstance(
		"iofs",
		d, linkStr)
	if err != nil {
		helper.ConsoleErrorMessageExit("new database instance error", err)
	}

	switch command {
	case "up":
		err = m.Up()
		if err != nil {
			helper.ConsoleErrorMessageExit("migrate up error", err)
		}
		color.Green("migrate up success")
	case "down":
		err = m.Down()
		if err != nil {
			helper.ConsoleErrorMessageExit("migrate down error", err)
		}
		color.Green("migrate down success")
	case "force":
		version := gconv.Int(gcmd.GetArg(3).String())
		err = m.Force(version)
		if err != nil {
			helper.ConsoleErrorMessageExit("migrate force error", err)
		}
		color.Green(fmt.Sprintf("force version %v susscess", version))
	case "create":
	case "new":
		fileName := gconv.String(gcmd.GetArg(3).String())
		err = newMigration(fileName)
		if err != nil {
			helper.ConsoleErrorMessageExit("create migrate file error", err)
		}
		color.Green(fmt.Sprintf("create migtaion file %v susscess", fileName))
	}
}
