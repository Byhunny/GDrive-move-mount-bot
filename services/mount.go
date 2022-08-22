package services

import (
	"fmt"
	"time"

	"github.com/Byhunny/GDrive-move-mount-bot/models/models"
)

func Mount(models.Var) string {
	time.Sleep(1 * time.Second)
	fmt.Println(models.Files, models.Index, models.Rclonename)

	result := "Mount are Done!"
	return result

}
