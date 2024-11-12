package main

import "fmt"

var LevelNodes = [18]uint64{0, 5, 15, 25, 35, 45, 55, 65, 75, 85, 95, 110, 125, 140, 160, 180, 200, 240}

func getNodesFromLevel(level uint64) uint64 {
	if level < 18 {
		return LevelNodes[level-1]
	}
	if level < 38 {
		return 200 + ((level - 17) * 40)
	}
	return 1000 + ((level - 37) * 50)
}

func getMinNodes(nodes uint64, level uint64) uint64 {
	if level == 1 {
		return 0
	} else if level == 2 {
		return 4
	}
	eightyPercent := uint64(float64(nodes) * 0.8)
	if eightyPercent < getNodesFromLevel(level-2) {
		return getNodesFromLevel(level - 2)
	} else {
		return eightyPercent
	}
}

func getMaxNodes(level uint64) uint64 {
	var maxNodes uint64
	if level < 18 {
		maxNodes = LevelNodes[level-1]
	} else if level < 37 {
		maxNodes = 200 + ((level - 17) * 40)
	} else {
		maxNodes = 950 + ((level - 36) * 50) //Als we 950 nemen start level 37 op 1000 wat de bedoeling is
	}
	return maxNodes
}

func getMinAndMaxNodes(level uint64) (uint64, uint64) {
	var maxNodes, minNodes uint64
	maxNodes = getMaxNodes(level)
	minNodes = getMinNodes(maxNodes, level)
	return minNodes, maxNodes
}

func main() {
	var startLevel, endLevel, minNodes, maxNodes uint64
	fmt.Print("\nStart Level: ")
	fmt.Scanln(&startLevel)
	fmt.Print("End Level: ")
	fmt.Scanln(&endLevel)
	for i := startLevel; i < endLevel+1; i++ {
		minNodes, maxNodes = getMinAndMaxNodes(i)
		fmt.Println("Level:", i, " MinNodes:", minNodes, "  MaxNodes:", maxNodes)
	}
}
