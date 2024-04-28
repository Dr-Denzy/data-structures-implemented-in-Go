package containers

// Container is the root type in the hierarchy of containers
// To satisfy the Container interface, these methods or operations must be implemented.
type Container interface {
	Size() int     // return the number of elements in the container
	IsEmpty() bool // returns true if and only if the container is empty
	Clear()        // clears out the elements in a container
}
