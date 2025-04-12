package main

import (
	"github.com/knightoverwatch/C2N-Launchpad-go/core"
	"github.com/knightoverwatch/C2N-Launchpad-go/global"
)

func main() {
	global.LP_VP = core.Viper()
	print(global.LP_VP)
}
