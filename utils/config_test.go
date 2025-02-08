package utils

import (
	"io"
	"os"
	"testing"
)

func TestConfig(t *testing.T) {
	hostsFilePath := "./hosts"
	hostsFile, err := os.Open(hostsFilePath)
	if err != nil {
		t.Fatal(err)
	}
	defer hostsFile.Close()

	hostsFileContentBytes, err := io.ReadAll(hostsFile)
	if err != nil {
		t.Fatal(err)
	}
	hostsFileContent := string(hostsFileContentBytes)

	responseFilePath := "./response"
	responseFile, err := os.Open(responseFilePath)
	if err != nil {
		t.Fatal(err)
	}
	defer responseFile.Close()

	responseFileContentBytes, err := io.ReadAll(responseFile)
	if err != nil {
		t.Fatal(err)
	}
	responseFileContent := string(responseFileContentBytes)

	updatedFileContent := UpdateConfigFileContent(hostsFileContent, responseFileContent)

	if updatedFileContent == hostsFileContent {
		t.Fatal("File content is not updated")
	}

	emptyHostsContent := ""
	if CheckConfigAvailable(emptyHostsContent) {
		t.Fatal("Config is available when it should not be")
	}
	if UpdateConfigFileContent(emptyHostsContent, responseFileContent) == emptyHostsContent {
		t.Fatal("File content is not updated")
	}
}
