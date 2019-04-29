package gofunc

func Insorting(array []int) []int{
	for i := 1; i < len(array); i++ {
		key := array[i]
		j := i - 1
		for ; (j >= 0) && (array[j] > key); j-- {
			array[j + 1] = array[j]
		}
		array[j + 1] = key
	}
	return array
}

