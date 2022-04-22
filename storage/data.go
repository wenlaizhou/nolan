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
}

// Worker 工作节点
type Worker struct {
	Name string
	Ip   string
}
