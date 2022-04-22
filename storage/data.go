package storage

// Cluster 集群
type Cluster struct {
}

// Master 控制节点
type Master struct {
	Name     string `json:"name"`
	TaskSlot int    `json:"task_slot"`
	Ip       string
}

// Task 任务
type Task struct {
	Slot   int          `json:"slot"`
	Name   string       `json:"name"`
	Status []TaskStatus `json:"status"`
}

// TaskStatus 任务状态
type TaskStatus struct {
}

// Worker 工作节点
type Worker struct {
	Name     string          `json:"name"`
	Ip       string          `json:"ip"`
	TaskSlot int             `json:"task_slot"`
	Tasks    map[string]Task `json:"tasks"`
}
