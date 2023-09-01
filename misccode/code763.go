package misccode

func PartitionLabels(s string) []int {
	ans := []int{}
	l := map[byte]int{}
	for i := 0; i < len(s); i++ {
		l[s[i]] = i
	}

	for i := 0; i < len(s); {
		position := l[s[i]]
		j := i
		for ; j <= position; j++ {
			if l[s[j]] > position {
				position = l[s[j]]
			}
		}
		ans = append(ans, j-i)
		i = j
	}
	return ans
}

func partitionLabels(s string) (partition []int) {
	lastPos := [26]int{}
	for i, c := range s {
		lastPos[c-'a'] = i
	}
	start, end := 0, 0
	for i, c := range s {
		if lastPos[c-'a'] > end {
			end = lastPos[c-'a']
		}
		if i == end {
			partition = append(partition, end-start+1)
			start = end + 1
		}
	}
	return
}
