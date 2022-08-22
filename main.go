package main

import (
	"fmt"

	"github.com/Byhunny/GDrive-move-mount-bot/services/mount"
)

var (
	Index      = [2]string{"1", "2"}
	Driveid    = [2]string{"0ACxWzIqJUAjtUk9PVA", ""}
	Files      = [2]string{"atlas1", "atlas2"}
	Rclonename = string("ahmetsoya444")
)

func main() {

	fmt.Print("New drive ID: ")
	_, err := fmt.Scanln(&Driveid[1])
	if err != nil {
		fmt.Println("New drive ID error: ", err)
	}
	mount.Mount()

}
