package es

type Config struct {
	SnapshotFrequency int64 `json:"snapshot_frequency" validate:"required,gte=0"`
}
