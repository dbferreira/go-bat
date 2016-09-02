package gobat

import (
	"fmt"
	"log"

	"github.com/zabawaba99/firego"
)

func HandleTask() {
	fmt.Println("Listening for new tasks...")

	fb := firego.New("https://go-bat.firebaseio.com", nil)
	fb.Auth("Q9tem1TtHnOabCUlv8lRNfX8IHtD3JBKoKrBwWJk")
	queuesFB := fb.Child("queues")

	queuesTeamsFB := queuesFB.Child("teams")
	teamsEvent := make(chan firego.Event)
	if err := queuesTeamsFB.Watch(teamsEvent); err != nil {
		log.Fatal(err)
	}

	defer queuesTeamsFB.StopWatching()

	for teamTaskEvent := range teamsEvent {
		data := teamTaskEvent.Data
		if data != nil {
			var err = *new(error)
			var taskKey = ""
			dataMaps := teamTaskEvent.Data.(map[string]interface{})
			var userKey = dataMaps["user"]
			if userKey != nil {
				err = startCreateTeam(dataMaps, fb)
				taskKey = userKey.(string)
			} else {
				for taskKey := range dataMaps {
					teamTask := dataMaps[taskKey].(map[string]interface{})
					err = startCreateTeam(teamTask, fb)
				}
			}
			if err != nil {
				log.Println("Error, could not create team")
				log.Fatal(err)
			}
			taskFB := queuesTeamsFB.Child(taskKey)
			if err := taskFB.Remove(); err != nil {
				log.Fatal(err)
			}
		}

		if teamTaskEvent.Type == firego.EventTypeError {
			log.Print("Error occurred, loop ending")
		}
	}

	fmt.Printf("Tasks have stopped")
}

func startCreateTeam(data map[string]interface{}, fb *firego.Firebase) error {
	var userID = data["user"].(string)
	newTeam, err := createTeam(userID, data["name"].(string), fb)
	if err != nil {
		log.Println("Error, could not create team")
		log.Fatal(err)
		return err
	}

	teamsFB := fb.Child("teams")
	if err := teamsFB.Update(newTeam); err != nil {
		log.Println("Error, could not set new team in FB")
		log.Fatal(err)
	}

	assignNewTeamPlayers(userID, fb)
	return err

}
