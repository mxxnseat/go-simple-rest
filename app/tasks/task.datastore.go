package tasks

import (
	"sync"
)

type Task struct{
    Id int `json:"id"`
    Name string `json:"name"`
}

type TaskStore struct{
	sync.Mutex

	tasks map[int]Task
	nextId int
}

func New() *TaskStore{
	ts := &TaskStore{}
	ts.tasks = make(map[int]Task)
	ts.nextId = 0
	return ts;
}

func (ts *TaskStore) Create(name string) Task{
	task := Task{
		Id: ts.nextId,
		Name: name,
	}
	ts.tasks[ts.nextId] = task
	ts.nextId++
	return task
}

func (ts *TaskStore) List() []Task{
	tasks := make([]Task,0,len(ts.tasks))
	for _,task := range ts.tasks{
		tasks = append(tasks, task)
	}
	return tasks
}