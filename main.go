package main

import (
	"archive/tar"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

// ConfigPath : default config path
const ConfigPath = "./config/vacuums.json"

// Config contains a list of vacuums to be used.
type Config struct {
	Vacuums []Vacuum
}

// Vacuum specifies destination to be cleaned
type Vacuum struct {
	Path         string
	FilesPrefix  string
	FilesSufix   string
	RemoveLogs   bool
	Compressor   string
	OutputPath   string
	OutputName   string
	UpdateOutput bool
}

// Reads vacuums.json into ./config dir
func loadVacuums() Config {
	file, e := ioutil.ReadFile(ConfigPath)
	if e != nil {
		fmt.Printf("It was impossible load vacuums file from %v. Error %v\n", ConfigPath, e)
		os.Exit(1)
	}

	fmt.Printf("Found vacuum file : %v\n", ConfigPath)
	var vacuums Config
	json.Unmarshal(file, &vacuums)
	return vacuums
}

// Checks if a path exists
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

// Builds a files path pattern using to path + prefix/suffix
func buildFilter(vacuum Vacuum) string {
	filesFilter := ""
	filesFilter += vacuum.Path
	if vacuum.FilesPrefix == "" && vacuum.FilesSufix == "" {
		filesFilter += "*"
	} else {
		if vacuum.FilesPrefix != "" {
			filesFilter += vacuum.FilesPrefix
		}
		if vacuum.FilesSufix != "" {
			filesFilter += vacuum.FilesSufix
		}
	}
	return filesFilter
}

// Get output file name formmating with date
func generateOutputFilename(vacuum Vacuum) string {
	t := time.Now()
	output := vacuum.OutputPath
	output += t.Format(vacuum.OutputName)
	return output
}

// Creates output file
func createOutput(vacuum Vacuum) *os.File {
	outputFile := generateOutputFilename(vacuum)
	if outputFile == "" {
		fmt.Printf("Invalid output name.\n")
		return nil
	}

	dirExists, err := exists(outputFile)
	if (!vacuum.UpdateOutput) && dirExists || err != nil {
		fmt.Printf("Output file found at %v but Logs Vacuum isn't allowed to update it.\n", outputFile)
		return nil
	}

	file, err := os.Create(outputFile)
	if err != nil {
		log.Fatalln(err)
	}
	return file
}

// Adding a file into a tarball
func addFile(tw *tar.Writer, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	if stat, err := file.Stat(); err == nil {
		// now lets create the header as needed for this file within the tarball
		header := new(tar.Header)
		header.Name = path
		header.Size = stat.Size()
		header.Mode = int64(stat.Mode())
		header.ModTime = stat.ModTime()

		// write the header to the tarball archive
		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		// copy the file data to the tarball
		if _, err := io.Copy(tw, file); err != nil {
			return err
		}
	}
	return nil
}

// remove a file
func remove(file string) {
	os.Remove(file)
}

// Compress a bunch of files found by vacuum
func compress(vacuum Vacuum, files []string) {
	if vacuum.Compressor != "tar.gz" {
		fmt.Printf("Error: Logs Vacuum only supports tar.gz compressor at the moment.\n")
		return
	}
	outputFile := createOutput(vacuum)
	if outputFile == nil {
		fmt.Printf("It was impossible to create output file.\n")
		return
	}
	defer outputFile.Close()
	fmt.Printf("Using output file %v\n", outputFile.Name())

	// setting up the gzip writer
	gw := gzip.NewWriter(outputFile)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()

	// adding files into output
	for _, file := range files {
		fmt.Printf("Adding log %s to output file\n", file)
		err := addFile(tw, file)
		if err != nil {
			fmt.Printf("It was impossible to add log file to output due to %v\n", err)
			continue
		}
		if vacuum.RemoveLogs {
			fmt.Printf("Removing file %s\n", file)
			remove(file)
		}
	}
}

// Runs a vacuum
func run(vacuum Vacuum) {
	dirExists, err := exists(vacuum.Path)
	if !dirExists || err != nil {
		fmt.Printf("It was impossible clean path %v: It doesn't exit or isn't readable.\n", vacuum.Path)
		return
	}
	fmt.Printf("Running vacuum at: %v\n", vacuum.Path)
	var path = buildFilter(vacuum)
	files, _ := filepath.Glob(path)
	if len(files) == 0 {
		fmt.Printf("There isn't files found using pattern %v\n", path)
		return
	}
	compress(vacuum, files)
}

// Validate if a vacuum has all mandatory fields
func isValid(vacuum Vacuum) bool {
	if vacuum.Path == "" {
		fmt.Printf("You must supply path in your vacuum configuration.\n")
		return false
	}
	if vacuum.OutputName == "" {
		fmt.Printf("You must supply outputName in your vacuum configuration.\n")
		return false
	}
	if vacuum.OutputPath == "" {
		fmt.Printf("You must supply outputPath in your vacuum configuration.\n")
		return false
	}
	if vacuum.Compressor == "" {
		fmt.Printf("You must supply compressor in your vacuum configuration.\n")
		return false
	}
	return true
}

func main() {
	fmt.Println("Running Logs Vacuum Cleaner")
	vacuums := loadVacuums()
	for _, vacuum := range vacuums.Vacuums {
		if isValid(vacuum) {
			run(vacuum)
		}
	}
}
