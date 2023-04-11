package lib

import (
	"bufio"
	"container/list"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func itemInList(item Item, list list.List) bool {
	for e := list.Front(); e != nil; e = e.Next() {
		//println("lis", e.Value.(*Item).Value)
		if item.Value == e.Value.(*Item).Value {
			return true
		}
	}
	return false
}

func euclideanDistance(x1, y1, x2, y2 float32) float32 {
	return float32(math.Sqrt(float64((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))))
}

func ReadFiletoGraph(dir string) (*Graph, []float32, []float32, error) {
	f, err := os.Open(dir)
	if err != nil {
		//log.Fatal(err)
		return nil, nil, nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	nodeCount, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		//log.Fatal(err)
		return nil, nil, nil, err
	}
	//println(parseInt)
	names := make([]string, int(nodeCount))
	xarr := make([]float32, int(nodeCount))
	yarr := make([]float32, int(nodeCount))
	if nodeCount < 8 {
		return nil, nil, nil, errors.New("invalid node count")
	}
	for i := int64(0); i < nodeCount; i++ {
		scanner.Scan()
		tmp := strings.Split(scanner.Text(), " ")
		if len(tmp) != 3 {
			return nil, nil, nil, errors.New("invalid argument number")
		}
		names[i] = tmp[0]
		x, err := strconv.ParseFloat(tmp[1], 64)
		if err != nil {
			//log.Fatal(err)
			return nil, nil, nil, err
		}
		xarr[i] = float32(x)
		y, err := strconv.ParseFloat(tmp[2], 64)
		if err != nil {
			//log.Fatal(err)
			return nil, nil, nil, err
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
		if len(matrix[j]) != len(tmp) {
			return nil, nil, nil, errors.New("out of bound")
		}
		for i := 0; i < len(tmp); i++ {
			parseF, err := strconv.ParseFloat(tmp[i], 64)
			if err != nil {
				//log.Fatal(err)
				return nil, nil, nil, err
			}
			if parseF < 0 {
				return nil, nil, nil, errors.New("negative weight")
			}
			matrix[j][i] = float32(parseF)
		}
	}
	return NewGraphNamed(matrix, names), xarr, yarr, nil
}

func NameToIndex(g Graph, name string) int {
	for i := 0; i < len(g.names); i++ {
		if g.names[i] == name {
			return i
		}
	}
	return -1
}

func PathToName(g Graph, i Item) []string {
	path := make([]string, 0)
	for e := i.PassedNode.Front(); e != nil; e = e.Next() {
		path = append(path, g.names[e.Value.(int)])
	}
	return path
}

func FileNameParse() (Graph, []float32, []float32) {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("invalid argument number")
	}
	g, x, y, err := ReadFiletoGraph(args[0])
	if err != nil {
		log.Fatal(err)
	}
	return *g, x, y
}

func RangedInput(min, max int) int {
	var read string
	var input int
	for {
		_, err := fmt.Scan(&read)
		input, err = strconv.Atoi(read)
		if err != nil {
			fmt.Println("invalid input")
			continue
		}
		if input >= min && input <= max {
			break
		}
		fmt.Println("invalid input")
	}
	return input
}

func PrintPath(g Graph, i Item) {
	if i.Priority == -1 {
		fmt.Println("No path found")
		return
	}
	fmt.Println("Path: ", PathToName(g, i))
	fmt.Println("Cost: ", i.Priority)
}
