package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type User struct {
	Name          string    `json:"name"`
	Password      string    `json:"-"`                       // "-" means this field will be omitted in the JSON output
	PreferredFish []string  `json:"preferredFish,omitempty"` // omitempty means this field will be omitted if it's empty
	CreatedAt     time.Time `json:"createdAt"`
}

func main() {
	user := &User{
		Name:      "Sammy the Shark",
		Password:  "fisharegreat",
		CreatedAt: time.Now(),
	}

	out, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println(string(out))
}
