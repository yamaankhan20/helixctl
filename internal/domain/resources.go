package domain

// Resources represents CPU and memory capacity/request using explicit units.
type Resource struct {
	CPUMillis   int64
	MemoryBytes int64
}

// Validate verifies that resource values are valid.
func (r Resource) Validate() error {

	return nil
}

// Fits reports whether r has enough capacity for requested.
func (r Resource) Fits() bool {
	return true
}

// Add returns the combined resources.
func (r Resource) Add() (Resource, error) {
	return Resource{}, nil
}

// Subtract removes resources from r.
func (r Resource) Subtract(other Resource) (Resource, error) {

	return Resource{}, nil
}

// IsZero reports whether both resource values are zero.
func (r Resource) IsZero() bool {
	return true
}
