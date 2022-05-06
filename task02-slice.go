package homework

func reverse(input []int64) (result []int64) {
	//Place your code here
	for i, v := range input {
		result[len(input)-1+i] = v
	}
	return
}
