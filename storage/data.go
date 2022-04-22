package storage

// Cluster 集群
type Cluster struct {
}

// Master 控制节点
type Master struct {
	Name string `json:"name"`
}

// Task 任务
type Task struct {
}

// Worker 工作节点
type Worker struct {
}
