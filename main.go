package main

import (
	_ "violettes-cms/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"violettes-cms/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
