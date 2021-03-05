package contributors

import (
	"git-monitoring/cmd"
	"git-monitoring/utility"
	"log"
)

// ContributorList returns the list of all the contributors into repo
func ContributorList(path string) ([]string, error) {
	utility.ChangeDirectory(path)
	resp, err := cmd.CmdExecute(`git log --format=%aN`)
	if err != nil {
		log.Println("error to get the contributors list", err)
		return []string{}, err
	}
	// log.Println("r", string(resp))
	return utility.Split("\n", resp), nil
}
