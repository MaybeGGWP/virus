package main

import (
	"fmt"
	"github.com/shirou/gopsutil/process"
	"os/exec"
	"sergogan/tasks"
	"time"
)

func browser() {

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
			if psName == "browser.exe" {
				exec.Command("rundll32", "url.dll,FileProtocolHandler", "https://gist.github.com/nanmu42/4fbaf26c771da58095fa7a9f14f23d27").Start()
				return
			}
		}
	}
}

func main() {
	tasks.Game()
	browser()
}
