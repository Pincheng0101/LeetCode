package p2222

func numberOfWays(s string) int64 {
	// index: string cnts
	// 0: "0", 1: "01", 2: "010"
	cnts010 := [3]int64{}
	// 0: "1", 1: "10", 2: "101"
	cnts101 := [3]int64{}

	for _, c := range s {
		if c == '0' {
			cnts010[0] += 1
			cnts101[1] += cnts101[0]
			cnts010[2] += cnts010[1]
		} else {
			cnts101[0] += 1
			cnts010[1] += cnts010[0]
			cnts101[2] += cnts101[1]
		}
	}
	return cnts010[2] + cnts101[2]
}
