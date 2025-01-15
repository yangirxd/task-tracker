package todo

import (
	"encoding/json"
	"os"
	"time"
)

type Item struct {
	Id          string
	Description string
	Status      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (i *Item) SetPriority(pri int) {
	switch pri {
	case 1:
		i.Status = 1
	case 3:
		i.Status = 3
	default:
		i.Status = 2
	}
}

func SaveItems(filename string, items []Item) error {
	b, err := json.Marshal(items)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}

	return nil
}

func ReadItems(filename string) ([]Item, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return []Item{}, err
	}

	var items []Item
	if err := json.Unmarshal(b, &items); err != nil {
		return []Item{}, err
	}
	return items, nil
}

func (i *Item) PrettyP() string {
	switch i.Status {
	case 1:
		return "in-progress"
	case 3:
		return "done"
	default:
		return "todo"
	}
}

type ByPri []Item

func (s ByPri) Len() int      { return len(s) }
func (s ByPri) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByPri) Less(i, j int) bool {
	if s[i].Status != s[j].Status {
		return s[i].Status < s[j].Status
	} else {
		return s[i].CreatedAt.After(s[j].CreatedAt)
	}
}
