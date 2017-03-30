package main

import (
	"testing"
	"time"
)

func TestBuildFilterWihoutPrefixAndSufix(t *testing.T) {
	vacuum := Vacuum{Path: "/tmp/"}
	outputFile := buildFilter(vacuum)
	expectedFile := "/tmp/*"
	if outputFile != expectedFile {
		t.Error("Filter " + expectedFile + " isn't equal to " + outputFile)
	}
}

func TestBuildFilterWithPrefix(t *testing.T) {
	vacuum := Vacuum{Path: "/tmp/", FilesPrefix: "logs*"}
	outputFile := buildFilter(vacuum)
	expectedFile := "/tmp/logs*"
	if outputFile != expectedFile {
		t.Error("Filter " + expectedFile + " isn't equal to " + outputFile)
	}
}

func TestBuildFilterWithSufix(t *testing.T) {
	vacuum := Vacuum{Path: "/tmp/", FilesSufix: "*.log"}
	outputFile := buildFilter(vacuum)
	expectedFile := "/tmp/*.log"
	if outputFile != expectedFile {
		t.Error("Filter " + expectedFile + " isn't equal to " + outputFile)
	}
}

func TestBuildFilterWithPrefixAndSufix(t *testing.T) {
	vacuum := Vacuum{Path: "/tmp/", FilesPrefix: "access-", FilesSufix: "*.log"}
	outputFile := buildFilter(vacuum)
	expectedFile := "/tmp/access-*.log"
	if outputFile != expectedFile {
		t.Error("Filter " + expectedFile + " isn't equal to " + outputFile)
	}
}

func TestGenerateOutputFilename(t *testing.T) {
	vacuum := Vacuum{OutputPath: "/tmp/", OutputName: "file.log"}
	outputFile := generateOutputFilename(vacuum)
	expectedFile := "/tmp/file.log"
	if outputFile != expectedFile {
		t.Error("File " + expectedFile + " isn't equal to " + outputFile)
	}
}

func TestGenerateOutputFilenameWithDate(t *testing.T) {
	vacuum := Vacuum{OutputPath: "/tmp/", OutputName: "log-2006.log"}
	now := time.Now()
	expectedFile := now.Format("/tmp/log-2006.log")
	outputFile := generateOutputFilename(vacuum)
	if outputFile != expectedFile {
		t.Error("File " + expectedFile + " isn't equal to " + outputFile)
	}
}

func TestValidVacuum(t *testing.T) {
	vacuum := Vacuum{Path: "/some/log/path/",
		FilesPrefix:  "server*",
		RemoveLogs:   true,
		Compressor:   "tar.gz",
		OutputPath:   "/some/backup/path/",
		OutputName:   "logs.2006-01.tag.gz",
		UpdateOutput: true}

	if !isValid(vacuum) {
		t.Error("Vacuum should be valid ")
	}
}

func TestInvalidVacuumWithoutPath(t *testing.T) {
	vacuum := Vacuum{FilesPrefix: "server*",
		RemoveLogs:   true,
		Compressor:   "tar.gz",
		OutputPath:   "/some/backup/path/",
		OutputName:   "logs.2006-01.tag.gz",
		UpdateOutput: true}

	if isValid(vacuum) {
		t.Error("Vacuum should be invalid because Path is missing. ")
	}
}

func TestInvalidVacuumWithoutOutputName(t *testing.T) {
	vacuum := Vacuum{Path: "/some/log/path/",
		FilesPrefix:  "server*",
		RemoveLogs:   true,
		Compressor:   "tar.gz",
		OutputPath:   "/some/backup/path/",
		UpdateOutput: true}

	if isValid(vacuum) {
		t.Error("Vacuum should be invalid because OutputName is missing. ")
	}
}

func TestInvalidVacuumWithoutOutputPath(t *testing.T) {
	vacuum := Vacuum{Path: "/some/log/path/",
		FilesPrefix:  "server*",
		RemoveLogs:   true,
		Compressor:   "tar.gz",
		OutputName:   "logs.2006-01.tag.gz",
		UpdateOutput: true}

	if isValid(vacuum) {
		t.Error("Vacuum should be invalid because OutputPath is missing. ")
	}
}

func TestInvalidVacuumWithoutCompressor(t *testing.T) {
	vacuum := Vacuum{Path: "/some/log/path/",
		FilesPrefix:  "server*",
		RemoveLogs:   true,
		OutputPath:   "/some/backup/path/",
		OutputName:   "logs.2006-01.tag.gz",
		UpdateOutput: true}

	if isValid(vacuum) {
		t.Error("Vacuum should be invalid because Compressor is missing. ")
	}
}
