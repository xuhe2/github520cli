package main

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/xuhe2/github520cli/utils"
)

func main() {
	// update hosts file once
	if len(os.Args) == 1 {
		UpdateHosts()
	}
	// has more args
	// if arg has `--auto/-a`, update hosts file every 1 hour
	args := os.Args[1:]
	for _, arg := range args {
		switch arg {
		case "--auto", "-a":
			handleAutoArg()
		}
	}
}

func handleAutoArg() {
	ticker := time.NewTicker(time.Hour)
	for range ticker.C {
		UpdateHosts()
	}
}

func UpdateHosts() {
	// call api for hosts file content
	url := "https://raw.hellogithub.com/hosts"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	newHostsContent := string(body)

	if utils.CheckConfigAvailable(newHostsContent) {
		slog.Info("config available")
	} else {
		slog.Error("config not available")
		os.Exit(1)
	}

	// get hosts config file
	hostFilePath := utils.GetHostsFilePath()

	hostFile, err := os.Open(hostFilePath)
	if err != nil {
		panic(err)
	}
	defer hostFile.Close()
	oldHostsContentBytes, err := io.ReadAll(hostFile)
	if err != nil {
		panic(err)
	}
	oldHostsContent := string(oldHostsContentBytes)

	// analy diff and update hosts file
	// analy diff
	diffs := utils.AnalyConfigDiff(oldHostsContent, newHostsContent)
	for _, diff := range diffs {
		fmt.Println(diff)
	}
	// update hosts file
	updatedHostsContent := utils.UpdateConfigFileContent(oldHostsContent, newHostsContent)
	// `0644`表示文件所有者有读写权限，其他用户只有读权限
	if err := os.WriteFile(hostFilePath, []byte(updatedHostsContent), 0644); err != nil {
		panic(err)
	}
}
