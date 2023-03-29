package tasks

import (
	"fmt"
	"github.com/reujab/wallpaper"
	"github.com/shirou/gopsutil/process"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func text() error {
	p, _ := os.UserHomeDir()
	p = filepath.Join(p, "Рабочий стол")
	p = filepath.Join(p, "Открой меня.txt")
	fmt.Println(p)
	tex, err := os.Create(p)
	if err != nil {
		return err
	}
	for i := 1; i < 3000; i++ {
		_, err = tex.WriteString("ДА")
		if err != nil {
			return err
		}
	}
	return nil
}

func Game() {
	var lastImage string
	rand.Seed(time.Now().UnixMicro())
	for {
		<-time.Tick(time.Second)
		psList, err := process.Processes()
		if err != nil {
			fmt.Println(err)
			continue
		}
		for _, ps := range psList {
			psName, err := ps.Name()
			if err != nil {
				continue
			}
			if psName == "dota2.exe" {
				currentImage, _ := wallpaper.Get()
				text()
				if lastImage == "" || (lastImage != currentImage) {
					err = wallpaper.SetFromURL(getRandomImage())
					text()
					if err != nil {
						fmt.Println(err)
						continue
					} else {
						lastImage, _ = wallpaper.Get()
					}
				}
			}
		}
	}
}

func getRandomImage() string {
	imagesURL := []string{
		"https://catherineasquithgallery.com/uploads/posts/2021-02/1612266407_110-p-luchshie-fioletovie-foni-v-stim-143.jpg",
		"https://catherineasquithgallery.com/uploads/posts/2021-03/1614645087_47-p-akvarel-fon-dlya-fotoshopa-55.jpg",
	}
	return imagesURL[rand.Intn(len(imagesURL))]
}
