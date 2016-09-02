package gobat

import (
	"fmt"

	"github.com/zabawaba99/firego"
)

func createTeam(userKey, teamName string, fb *firego.Firebase) (map[string]interface{}, error) {
	fmt.Println("Creating a new team for username", userKey)
	team := map[string]string{
		"name": teamName,
		"user": userKey,
	}
	newTeam := map[string]interface{}{
		userKey: team,
	}

	return newTeam, nil
}
