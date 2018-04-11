package disjointSetTree

const offlineminimumExtract = -1

//input sequence: -1 means extract, minimum insert value should be 1
func offLineMinimum(seq []int) []int {
	//init sets
	if len(seq) == 0 {
		panic("seq must not be empty!")
	}
	if seq[0] == offlineminimumExtract {
		panic("first item must be insert!")
	}
	//insert set, value is index by m
	insertSets := make([]*disjointSet, 0, 0)
	//value set, value is insert value
	values := make([]*disjointSet, len(seq), cap(seq))
	n, m := 1, 0
	insertSets = append(insertSets, makeSet(0))
	for i := range seq {
		if seq[i] == offlineminimumExtract {
			m++
			insertSets = append(insertSets, makeSet(m))
		} else {
			if seq[i] < 1 {
				panic("minimum insert value must >= 1!")
			}
			values[seq[i]] = makeSet(seq[i])
			insertSets[m] = union(insertSets[m], values[seq[i]])
			n++
		}
	}

	//get minimum sequence
	extractSeq := make([]int, m, m)
	for i := 1; i < n; i++ {
		j := findSet(values[i]).Value.(int)
		if j != m {
			extractSeq[j] = i
			for l := j + 1; l <= m; l++ {
				if insertSets[l] != nil {
					insertSets[l] = union(insertSets[l], insertSets[j])
					insertSets[l].Value = l
					insertSets[j] = nil
					break
				}
			}
		}
	}
	return extractSeq
}
