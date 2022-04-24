package storage

// Cluster 集群
type Cluster struct {
	master    Master
	candidate Master
	workers   []Worker
}

// Master 控制节点
type Master struct {
	Name          string `json:"name"`
	TotalTaskSlot int    `json:"totalTaskSlot"`
	IdleTaskSlot  int    `json:"idleTaskSlot"`
	Ip            string
}

// Task 任务
type Task struct {
	Slot          int          `json:"slot"`
	Name          string       `json:"name"`
	CurrentStatue TaskStatus   `json:"currentStatue"`
	Status        []TaskStatus `json:"status"`
}

// TaskStatus 任务状态
type TaskStatus struct {
	Log        []string `json:"log"`
	CreateTime int      `json:"createTime"`
	Statue     string   `json:"statue"`
}

// Worker 工作节点
type Worker struct {
	Name          string          `json:"name"`
	Ip            string          `json:"ip"`
	TotalTaskSlot int             `json:"totalTaskSlot"`
	IdleTaskSlot  int             `json:"idleTaskSlot"`
	Tasks         map[string]Task `json:"tasks"`
}
