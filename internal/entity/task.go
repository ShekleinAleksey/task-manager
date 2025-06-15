package entity

import (
	"sync"
	"time"
)

type TaskStatus string

const (
	StatusPending   TaskStatus = "pending"
	StatusRunning   TaskStatus = "running"
	StatusCompleted TaskStatus = "completed"
	StatusFailed    TaskStatus = "failed"
)

type Task struct {
	ID        string     `json:"id"`
	Status    TaskStatus `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	Result    string     `json:"result,omitempty"`
	Duration  string     `json:"duration,omitempty"`
}

type TaskManager struct {
	sync.Mutex
	tasks map[string]*Task
}
