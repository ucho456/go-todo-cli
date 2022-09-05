package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Todos []item

func (t *Todos) Add(task string) {
	todo := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*t = append(*t, todo)
}

func (t *Todos) Comp(id int) error {
	ls := *t
	if id <= 0 || id > len(ls) {
		return errors.New("不正なIDです。")
	}

	ls[id-1].CompletedAt = time.Now()
	ls[id-1].Done = true

	return nil
}

func (t *Todos) Del(id int) error {
	ls := *t
	if id <= 0 || id > len(ls) {
		return errors.New("不正なIDです。")
	}

	*t = append(ls[:id-1], ls[id:]...)

	return nil
}

func (t *Todos) Print() {
	for i, item := range *t {
		i++
		fmt.Printf("%d - %s\n", i, item.Task)
	}
}

func (t *Todos) Load(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return err
	}
	if err := json.Unmarshal(file, t); err != nil {
		return err
	}
	return nil
}

func (t *Todos) Store(filename string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}
