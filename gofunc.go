package gofunc

func Filter[T any](s []T, f func(T) bool) []T {
	if len(s) == 0 {
		return s
	}

	filteredRecords := make([]T, len(s))
	for _, record := range s {
		if f(record) {
			filteredRecords = append(filteredRecords, record)
		}
	}
	return filteredRecords
}

// FlatMap takes a slice and a function and returns a single slice with all the elements
func FlatMap[T, U any](s [][]T, f func(T) U) []U {
	flattenedRecords := make([]U, 0)

	for _, record := range s {
		for _, element := range record {
			flattenedRecords = append(flattenedRecords, f(element))
		}
	}
	return flattenedRecords
}

// Flatten takes a slice of slices and returns a single slice with all the elements
func Flatten[T any](s [][]T) []T {
	flattenedRecords := make([]T, 0)

	for _, record := range s {
		flattenedRecords = append(flattenedRecords, record...)
	}
	return flattenedRecords
}

// Map takes a slice and a function and returns a new slice with the function applied to each element
func Map[T, U any](s []T, f func(T) U) []U {
	mappedRecords := make([]U, len(s))

	for i, record := range s {
		mappedRecords[i] = f(record)
	}
	return mappedRecords
}

// Unique takes a slice and returns a new slice with only the unique elements
func Unique[T comparable](s []T) []T {
	if len(s) == 0 {
		return s
	}

	uniqueRecords := make([]T, 0)
	uniqueMap := make(map[T]bool)

	for _, record := range s {
		if _, ok := uniqueMap[record]; !ok {
			uniqueMap[record] = true
			uniqueRecords = append(uniqueRecords, record)
		}
	}
	return uniqueRecords
}
