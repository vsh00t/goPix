package ui

import (
	"time"

	"github.com/schollz/progressbar/v3"
)

func ProgBar() {
	bar := progressbar.Default(int64(30))
	for i := 0; i < 120; i++ {
		bar.Add(1)
		time.Sleep(1 * time.Second)
	}
}
