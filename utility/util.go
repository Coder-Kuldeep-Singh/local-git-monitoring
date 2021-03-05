package utility

import (
	"log"
	"os"
	"strings"
)

// OpenFile opens the file and folder
func OpenFile(path string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// GetAllFiles returns the list of the all file into repo
func GetAllFiles(file *os.File) []string {
	// checkConfigFileInDir return the all files into the directory
	defer file.Close()
	list, err := file.Readdirnames(0) // 0 to read all files and folders
	if err != nil {
		log.Println(err)
		return nil
	}
	return list
}

// CheckExpectedFile return the bool if file name matched
func CheckExpectedFile(name string, files []string) bool {
	for _, file := range files {
		if file == name {
			return true
		}
	}
	return false
}

// Split splits the array by common character
func Split(character string, str []byte) []string {
	return strings.Split(string(str), character)
}

// ChangeDirectory changes the pwd to destination direcotory
func ChangeDirectory(dest string) {
	err := os.Chdir(dest)
	if err != nil {
		log.Printf("error to change to the destination directory %s", err.Error())
	}
}

// Distinct returns the distinct values
func Distinct(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
