package persistor

// Persistor provides persistence for named data.
type Persistor interface {
	// Persist persists named data to the specified path.
	Persist(name string, data []byte, path string) error
}
