package consoles

import (
	"github.com/gogf/gf/v2/os/gcmd"
	"violettes-cms/internal/database"
)

func Console() {
	command := gcmd.GetArg(1).String()
	switch command {
	case "migrate":
		database.Migrate()
	}
}
