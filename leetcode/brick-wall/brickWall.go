package main

func leastBricks(wall [][]int) int {

	m := make(map[int]int)

	for _, brickRow := range wall {

		edgeLocation := 0

		for i := 0; i < len(brickRow)-1; i++ {

			brick := brickRow[i]
			edgeLocation = edgeLocation + brick
			m[edgeLocation] = m[edgeLocation] + 1

		}

	}

	maxEdgeNO := 0

	for _, v := range m {

		if v > maxEdgeNO {
			maxEdgeNO = v
		}
	}

	return len(wall) - maxEdgeNO

}
