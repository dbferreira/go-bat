package gobat

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/zabawaba99/firego"
)

const teamSize = 15

var countries = map[string]string{
	"ZA": "Netherlands",
	"ZW": "Canada",
	"AU": "Australia",
	"NZ": "New+Zealand",
	"GB": "England",
	"IN": "India",
	"LK": "India",
	"PK": "Pakistan",
	// "NL": "Netherlands",
	// "US": "United+States",
	// "Albania",
	// "Argentina",
	// "Armenia",
	// "Australia",
	// "Austria",
	// "Azerbaijan",
	// "Bangladesh",
	// "Belgium",
	// "Bosnia and Herzegovina",
	// "Brazil",
	// "Canada",
	// "China",
	// "Colombia",
	// "Denmark",
	// "Egypt",
	// "England",
	// "Estonia",
	// "Finland",
	// "France",
	// "Georgia",
	// "Germany",
	// "Greece",
	// "Hungary",
	// "India",
	// "Iran",
	// "Israel",
	// "Italy",
	// "Japan",
	// "Korea",
	// "Mexico",
	// "Morocco",
	// "Netherlands",
	// "New Zealand",
	// "Nigeria",
	// "Norway",
	// "Pakistan",
	// "Poland",
	// "Portugal",
	// "Romania",
	// "Russia",
	// "Slovakia",
	// "Slovenia",
	// "Spain",
	// "Sweden",
	// "Switzerland",
	// "Turkey",
	// "Ukraine",
	// "United States",
	// "Vietnam",
}

type uiNames struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Gender  string `json:"gender"`
	Region  string `json:"region"`
}

func newPlayer(userKey string, age int, nationality string) map[string]interface{} {
	fmt.Println("Creating random player for team...")
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	now := time.Now().UnixNano() / 1000000 // Convert to milliseconds
	countryCode, country := randomCountry()
	newPlayer := map[string]interface{}{
		"age":         age,
		"name":        getName(country),
		"created":     now,
		"nationality": strings.ToLower(countryCode),
		"team":        userKey,
		"batting":     random.Intn(8),
		"bowling":     random.Intn(8),
		"stamina":     random.Intn(6),
		"fitness":     5,
	}
	return newPlayer
}

func getName(Country string) string {
	var uiNamesURL = "http://uinames.com/api/?gender=male&region=" + strings.ToLower(Country)

	n := new(uiNames)
	getJSON(uiNamesURL, n)
	return fmt.Sprintf("%v %v", n.Name, n.Surname)
}

func getJSON(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func createPlayer(userKey string, p map[string]interface{}, fb *firego.Firebase) {
	playersFB := fb.Child("players/" + userKey)
	_, err := playersFB.Push(p)
	if err != nil {
		log.Println("Error, could not set new player in FB")
		log.Fatal(err)
	}
}

func assignNewTeamPlayers(userKey string, fb *firego.Firebase) {
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	for x := 0; x < teamSize; x++ {
		r := random.Intn(15)
		p := newPlayer(userKey, 16+r, "za") // Todo: bug here, get country code from user?
		createPlayer(userKey, p, fb)
	}
}

func randomCountry() (cc string, c string) {
	i := rand.Intn(len(countries))
	for countryCode := range countries {
		if i == 0 {
			return countryCode, countries[countryCode]
		}
		i--
	}
	panic("never")
}
