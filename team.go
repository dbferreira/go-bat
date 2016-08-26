package gobat

import "fmt"

func createTeam(userKey string) (map[string]string, error) {
	fmt.Println("Creating a new team...")
	newTeam := map[string]string{
		"name": "newname",
		"user": userKey,
	}
	return newTeam, nil
}
