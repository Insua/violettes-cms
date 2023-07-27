package main

import (
	"github.com/gogf/gf/v2/os/gcmd"
	"os"
	"violettes-cms/internal/cmd/consoles"
	_ "violettes-cms/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"violettes-cms/internal/cmd"
)

func main() {
	if len(gcmd.GetArg(1).String()) > 0 {
		consoles.Console()
		os.Exit(0)
	}
	cmd.Main.Run(gctx.GetInitCtx())
}
