package storage

// Service storage service
type Service interface {

	// Store write
	Store(cluster Cluster) error

	// Read read
	Read() (Cluster, error)
}
