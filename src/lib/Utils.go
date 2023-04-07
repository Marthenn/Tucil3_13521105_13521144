package lib

import (
	"bufio"
	"container/list"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func itemInList(item Item, list list.List) bool {
	for e := list.Front(); e != nil; e = e.Next() {
		if item.Value == e.Value {
			return true
		}
	}
	return false
}

func euclideanDistance(x1, y1, x2, y2 float32) float32 {
	return float32(math.Sqrt(float64((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))))
}

func ReadFiletoGraph(dir string) (*Graph, []float32, []float32) {
	f, err := os.Open(dir)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	nodeCount, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	//println(parseInt)
	names := make([]string, int(nodeCount))
	xarr := make([]float32, int(nodeCount))
	yarr := make([]float32, int(nodeCount))
	for i := int64(0); i < nodeCount; i++ {
		scanner.Scan()
		tmp := strings.Split(scanner.Text(), " ")
		names[i] = tmp[0]
		x, err := strconv.ParseFloat(tmp[1], 64)
		if err != nil {
			log.Fatal(err)
		}
		xarr[i] = float32(x)
		y, err := strconv.ParseFloat(tmp[2], 64)
		if err != nil {
			log.Fatal(err)
		}
		yarr[i] = float32(y)
	}
	matrix := make([][]float32, nodeCount)
	for i := range matrix {
		matrix[i] = make([]float32, nodeCount)
	}
	for j := 0; j < int(nodeCount); j++ {
		scanner.Scan()
		tmp := strings.Split(scanner.Text(), " ")
		for i := 0; i < len(tmp); i++ {
			parseF, err := strconv.ParseFloat(tmp[i], 64)
			if err != nil {
				log.Fatal(err)
			}
			matrix[j][i] = float32(parseF)
		}
	}
	return NewGraphNamed(matrix, names), xarr, yarr
}
