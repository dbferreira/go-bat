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
	queueFB := fb.Child("queues")
	teamsFB := fb.Child("teams")

	teamTaskFB := queueFB.Child("teams")
	teamsEvent := make(chan firego.Event)
	if err := teamTaskFB.Watch(teamsEvent); err != nil {
		log.Fatal(err)
	}

	defer teamTaskFB.StopWatching()

	for teamTaskEvent := range teamsEvent {
		data := teamTaskEvent.Data
		log.Println("Event Received")
		log.Printf("Type: %s\n", teamTaskEvent.Type)
		log.Printf("Path: %s\n", teamTaskEvent.Path)
		log.Printf("Data: %v\n", teamTaskEvent.Data)
		if data != nil {
			dataMaps := teamTaskEvent.Data.(map[string]interface{})
			log.Println("DataMaps = ", dataMaps)
			for taskKey := range dataMaps {
				teamTask := dataMaps[taskKey].(map[string]interface{})

				log.Println("teamTask = ", teamTask)
				newTeam, err := createTeam(taskKey)
				if err != nil {
					log.Println("Error, could not create team")
					log.Fatal(err)
				}
				log.Println("Task:", teamTask)
				log.Println("User:", teamTask["user"])
				log.Println("Created:", teamTask["created"])

				userTeamFB := teamsFB.Child(taskKey)

				if err := userTeamFB.Set(newTeam); err != nil {
					log.Fatal(err)
				}

				taskFB := teamTaskFB.Child(taskKey)
				if err := taskFB.Remove(); err != nil {
					log.Fatal(err)
				}
			}
		}
		// // log.Printf("My value %v\n", v)

		if teamTaskEvent.Type == firego.EventTypeError {
			log.Print("Error occurred, loop ending")
		}
	}
	fmt.Printf("Tasks have stopped")
}
