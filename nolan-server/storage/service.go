package storage

// Service storage service
type Service interface {

	// Store write
	Store(cluster Cluster)

	// Read read
	Read() Cluster
}
