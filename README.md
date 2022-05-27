- 👋 Hi, I’m @janlauber
- 👀 I’m interested in the following CNCF projects:
    - [Kubernetes](https://github.com/kubernetes/kubernetes)
    - [RKE2](https://github.com/rancher/rke2)
    - [FluxCD](https://github.com/fluxcd/flux)
    - [Longhorn](https://github.com/longhorn/longhorn)
    - [Cilium](https://github.com/cilium/cilium)
- 🛠️ Also I'm learning the following languages/frameworks to code some awesome web applications:
    - [NextJS](https://nextjs.org)
    - [GoLang](https://go.dev)
- 📖 In 2020, I started my bachelors degree at the [Bern University of Applied Sciences](https://bfh.ch)
- 🌱 I’m currently learning programming and other concepts to contribute to some projects listed above or other ones.
- 💞️ I’m looking to collaborate on building reliable PaaS and Cloud Native tools/applications.
- 📫 You can reach me on my private email address: jan.lauber+github@protonmail.ch
- 📇 [LinkedIn](https://www.linkedin.com/in/jan-lauber/), [Instagram](https://www.instagram.com/jaenu.lauber/)
<hr>
<details>
    <summary>
    <i><b> Nerd section<b></i> 👨‍💻
    </summary>

#### Development:

```go
// REST API for information about janlauber
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
```
</details>
<hr>

#### Statistics

<a href="https://youtu.be/dQw4w9WgXcQ">
  <img align="" src="https://github-readme-stats.vercel.app/api?username=janlauber&count_private=true&theme=vue&show_icons=true&hide_title=false&count_private=true&include_all_commits=true" />
</a>

