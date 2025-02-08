package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/xuhe2/github520cli/utils"
)

func main() {
	UpdateHosts()
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
		fmt.Println("配置可用")
	} else {
		fmt.Println("配置不可用")
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
