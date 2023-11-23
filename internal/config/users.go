package config

import (
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// Get users from users.yaml
// Users is map[username][]filename
type Users struct {
	Users map[string][]string `yaml:"users"`
}

var users Users

func GetUsers() Users {
	return users
}

func loadUsers() {
	// Load users.yaml
	data, err := os.ReadFile("users.yaml")
	if err != nil {
		// Try to create users.yaml
		_, err := os.Create("users.yaml")
		if err != nil {
			logrus.Fatal("Failed to read and/or create users.yaml")
		}
		logrus.Info("Created users.yaml")
		return
	}

	err = yaml.Unmarshal(data, &users)
	if err != nil {
		return
	}
}

func init() {

	loadUsers()

	// Listen for changes in users.yaml
	go func() {
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			logrus.Fatal("Failed to create file watcher:", err)
		}
		defer watcher.Close()

		err = watcher.Add("users.yaml")
		if err != nil {
			logrus.Fatal("Failed to add users.yaml to file watcher:", err)
		}

		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					logrus.Info("users.yaml modified. Reloading...")
					loadUsers()
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				logrus.Error("File watcher error:", err)
			}
		}
	}()

	// users.yaml example:
	// users:
	//   user1:
	//     - file1
	//     - file2
	//   user2:
	//     - file3
	//     - file4

	logrus.Info("Loaded users.yaml")
}
