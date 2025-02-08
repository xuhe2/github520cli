package utils

import (
	"regexp"
	"runtime"
	"strings"

	"github.com/samber/lo"
)

const (
	CONFIG_REGEX = `(?s)# GitHub520 Host Start.*?# GitHub520 Host End`
	ITEM_REGEX   = `^(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})\s+(.+)$`
)

func CheckConfigAvailable(config string) bool {
	re := regexp.MustCompile(CONFIG_REGEX)
	return re.MatchString(config)
}

func GetHostsFilePath() string {
	os := runtime.GOOS
	switch os {
	case "windows":
		return "C:\\Windows\\System32\\drivers\\etc\\hosts"
	case "darwin":
		return "/etc/hosts"
	case "linux":
		return "/etc/hosts"
	}
	return ""
}

func GetConfig(content string) string {
	if !CheckConfigAvailable(content) {
		return ""
	}

	return regexp.MustCompile(CONFIG_REGEX).FindString(content)
}

func ParseConfig(config string) map[string]string {
	if !CheckConfigAvailable(config) {
		return make(map[string]string)
	}

	configMap := make(map[string]string)
	re := regexp.MustCompile(ITEM_REGEX)
	configLines := FilterEmptyLines(GetLines(config))
	for _, line := range configLines {
		match := re.FindStringSubmatch(line)
		if match != nil {
			configMap[match[2]] = match[1]
		}
	}
	return configMap
}

func AnalyConfigDiff(oldConfig, newConfig string) []string {
	var diff []string

	oldConfigMap := ParseConfig(GetConfig(oldConfig))
	newConfigMap := ParseConfig(GetConfig(newConfig))

	oldConfigMapKeys := lo.Keys(oldConfigMap)
	newConfigMapKeys := lo.Keys(newConfigMap)

	// check if there are any new entries in the new config
	for _, key := range newConfigMapKeys {
		if _, ok := oldConfigMap[key]; !ok {
			diff = append(diff, "[New entry] "+key+": "+newConfigMap[key])
		} else if oldConfigMap[key] != newConfigMap[key] {
			diff = append(diff, "[Updated entry] "+key+": "+oldConfigMap[key]+" -> "+newConfigMap[key])
		}
	}

	// check if there are any entries that have been removed from the new config
	for _, key := range oldConfigMapKeys {
		if _, ok := newConfigMap[key]; !ok {
			diff = append(diff, "[Removed entry] "+key+": "+oldConfigMap[key])
		}
	}

	return diff
}

func UpdateConfigFileContent(oldContent, newContent string) string {
	oldConfig := GetConfig(oldContent)
	newConfig := GetConfig(newContent)

	if newConfig == "" {
		return oldContent
	}

	if oldConfig == "" {
		return oldContent + "\n" + newConfig
	}

	return strings.Replace(oldContent,
		oldConfig,
		newConfig,
		1)
}
