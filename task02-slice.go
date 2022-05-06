package homework

func reverse(input []int64) (result []int64) {
	//Place your code here
	for i, _ := range input {
		result = append(result, input[len(input)-1-i])
	}
	return
}
