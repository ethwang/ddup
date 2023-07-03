package misccode

func CalcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	record := map[string]int{}
	for _, equation := range equations {
		v1, v2 := equation[0], equation[1]
		if _, isExist := record[v1]; !isExist {
			record[v1] = len(record)
		}
		if _, isExist := record[v2]; !isExist {
			record[v2] = len(record)
		}
	}

	graph := make([][]float64, len(record))
	for i := 0; i < len(graph); i++ {
		graph[i] = make([]float64, len(record))
	}
	for index, equation := range equations {
		graph[record[equation[0]]][record[equation[1]]] = values[index]
		graph[record[equation[1]]][record[equation[0]]] = 1 / values[index]
	}

	res := []float64{}
	for _, query := range queries {
		v1, v2 := query[0], query[1]
		result := BFS(graph, record[v1], record[v2])
		res = append(res, result)
	}
	return res
}
func BFS(graph [][]float64, start, end int) float64 {
	r := make(map[int]float64)
	r[start] = 1.0
	queue := []int{start}
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		if v == end {
			return r[v]
		}

		for i := 0; i < len(graph[v]); i++ {
			if graph[v][i] > 0 && r[i] == 0 {
				r[i] = r[v] * graph[v][i]
				queue = append(queue, i)
			}
		}
	}
	return -1.0
}
