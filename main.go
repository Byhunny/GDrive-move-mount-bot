package main

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/Byhunny/go-gdrive-mount-move-bot/services"
)

var (
	Index      = [2]string{"1", "2"}
	Driveid    = [3]string{"0ACxWzIqJUAjtUk9PVA", "AAAABBBB", "reverse"}
	Files      = [2]string{"atlas1", "atlas2"}
	Rclonename = string("ahmetsoya444")
)

func main() {
	for true {
		fmt.Print("New drive ID: ")
		_, err := fmt.Scanln(&Driveid[1])
		if err != nil {
			fmt.Println("New drive ID error: ", err)
		}

		_ = exec.Command("sed -i 's/" + Driveid[2] + "/" + Driveid[1] + "/g' /root/.config/rclone/rclone.conf")

		time.Sleep(1 * time.Second)

		services.Move(Files)

		time.Sleep(1 * time.Second)
		r_driveid := services.Reverse(Driveid[1])
		_ = exec.Command("sed -i 's/" + Driveid[1] + "/" + r_driveid + "/g' /root/.config/rclone/rclone.conf")

		time.Sleep(1 * time.Second)

		_ = exec.Command("sed -i 's/" + Driveid[0] + "/" + Driveid[1] + "/g' /root/.config/rclone/rclone.conf")
		time.Sleep(1 * time.Second)

		services.Mount(Index, Files, Rclonename)
		time.Sleep(1 * time.Second)

		Driveid[2] = r_driveid

		fmt.Println("All process success...")

	}

}
