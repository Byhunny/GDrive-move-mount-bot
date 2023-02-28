package services

import (
	"os/exec"
	"time"
)

func Mount(Index [2]string, Files [2]string, Rclonename string) string {

	_ = exec.Command("fusermount -uz /" + Files[0])
	time.Sleep(1 * time.Second)
	_ = exec.Command("fusermount -uz /" + Files[1])
	time.Sleep(1 * time.Second)
	_ = exec.Command("rclone mount " + Rclonename + ":/Sch/sc" + Index[0] + " /" + Files[0] + " --vfs-read-chunk-size=128k --poll-interval=1h --dir-cache-time=2h --buffer-size=0 --cache-dir /mnt/temp --vfs-cache-mode full --no-checksum --no-modtime --read-only --vfs-read-wait 0 --max-read-ahead 0 --use-mmap --fast-list --checkers 2 --no-check-certificate --multi-thread-cutoff 0 --multi-thread-streams 2 --vfs-cache-max-age 10000h -q --use-cookies --daemon")
	time.Sleep(1 * time.Second)
	_ = exec.Command("rclone mount " + Rclonename + ":/Sch/sc" + Index[1] + " /" + Files[1] + " --vfs-read-chunk-size=128k --poll-interval=1h --dir-cache-time=2h --buffer-size=0 --cache-dir /mnt/temp --vfs-cache-mode full --no-checksum --no-modtime --read-only --vfs-read-wait 0 --max-read-ahead 0 --use-mmap --fast-list --checkers 2 --no-check-certificate --multi-thread-cutoff 0 --multi-thread-streams 2 --vfs-cache-max-age 10000h -q --use-cookies --daemon")
	time.Sleep(1 * time.Second)

	return "Mount are Done!"

}
