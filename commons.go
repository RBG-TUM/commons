package commons

// Unique filters duplicates of a slice l of elements that are reducible to a uint by the function toId, e.g. gorm models.
func Unique[T any](l []T, toId func(T) uint) []T {
	var unique map[uint]bool
	var result []T
	for _, e := range l {
		id := toId(e)
		if _, ok := unique[id]; !ok {
			unique[id] = true
			result = append(result, e)
		}
	}
	return result
}
