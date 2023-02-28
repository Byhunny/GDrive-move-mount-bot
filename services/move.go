package services

import (
	"os/exec"
	"time"
)

func Move(Index [2]string) string {
	_ = exec.Command("./gclone move movesource:/Sch/sc" + Index[0] + "/ movetarget:/Sch/sc" + Index[0] + "/ --drive-server-side-across-configs -v")
	time.Sleep(1 * time.Second)
	_ = exec.Command("./gclone move movesource:/Sch/sc" + Index[1] + "/ movetarget:/Sch/sc" + Index[1] + "/ --drive-server-side-across-configs -v")
	time.Sleep(1 * time.Second)

	return "Move are Success!"

}
