package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/alexeyco/simpletable"
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
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "ID"},
			{Align: simpletable.AlignCenter, Text: "タスク"},
			{Align: simpletable.AlignCenter, Text: "完了"},
			{Align: simpletable.AlignRight, Text: "作成時刻"},
			{Align: simpletable.AlignRight, Text: "完了時刻"},
		},
	}

	var cells [][]*simpletable.Cell

	for i, item := range *t {
		i++
		task := blue(item.Task)
		done := blue("no")
		completedAt := item.CompletedAt.Format("2006/01/02 03:04:05")
		if item.Done {
			task = green(item.Task)
			done = green("yes")
		} else {
			completedAt = ""
		}
		cells = append(cells, []*simpletable.Cell{
			{Text: fmt.Sprintf("%d", i)},
			{Text: task},
			{Text: done},
			{Text: item.CreatedAt.Format("2006/01/02 03:04:05")},
			{Text: completedAt},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Span: 5, Text: red(fmt.Sprintf("残りタスク数 %d", t.CountPending()))},
	}}

	table.SetStyle(simpletable.StyleUnicode)

	table.Println()
}

func (t *Todos) CountPending() int {
	total := 0
	for _, item := range *t {
		if !item.Done {
			total++
		}
	}

	return total
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
