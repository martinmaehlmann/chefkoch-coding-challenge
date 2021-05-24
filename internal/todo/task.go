package todo

import (
	"encoding/json"
	"gorm.io/gorm"
)

// Task represents one task.
type Task struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	TodoID      int
}

func (t *Task) PrettyString(indentation string) (string, error) {
	jsonString, err := json.MarshalIndent(t, "", indentation)
	if err != nil {
		return "", err
	}

	return string(jsonString), nil
}

// Valid checks, if the Name has been set.
func (t *Task) Valid() bool {
	if len(t.Name) == 0 {
		return false
	}

	return true
}
