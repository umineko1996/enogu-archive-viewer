package main

import (
	"fmt"
	"os"

	"github.com/umineko1996/enogu-archive-viewer/no6"
)

func main() {
	config := no6.Config{
		Email:    "XXXX",
		Password: "XXXX",
	}
	if _, err := os.Stat(no6.ArchivesDir); err == nil {
		fmt.Println("既にアーカイブページがDLされています")
		os.Exit(0)
	}

	client, err := no6.NewClient(config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := client.GetALLArchivesPage(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
