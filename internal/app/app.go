package app

import (
	"activity/internal/app/config"
)

func Init() {
	config.MustLoad()
}
