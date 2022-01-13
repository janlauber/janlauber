package main

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	Version = "0.0.1"
)

type Person struct {
	Name     string
	Age      int
	Gender   string
	Location string
}

type Dev struct {
	Person     Person
	Code       []string
	Tools      []string
	Frameworks []string
	Design     []string
	Hobbies    []string
}

func main() {
	janlauber := &Dev{
		Person: Person{
			Name:     "Jan Lauber",
			Age:      22,
			Gender:   "male",
			Location: "Switzerland",
		},
		Code:       []string{"go", "java", "javascript", "html5", "css3", "bash"},
		Tools:      []string{"helm", "git", "zshell", "docker", "kubernetes", "rancher", "fluxcd", "longhorn"},
		Frameworks: []string{"react", "nextjs", "nodejs"},
		Design:     []string{"adobe illustrator", "figma", "sketch"},
		Hobbies:    []string{"gym", "volleyball", "swimming", "music production", "nerdy stuff"},
	}

	// convert to json
	json, err := json.Marshal(janlauber)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("request received")
		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	})

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
