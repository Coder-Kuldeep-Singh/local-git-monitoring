package utility

import (
	"log"
	"os"
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
