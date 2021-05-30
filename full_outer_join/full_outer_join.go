package full_outer_join

import (
	"os"
	"sort"
	"strings"
	"sync"
)

func insertDataToMap(filePath string, hashMap *map[string]bool, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Open(filePath)
	if err != nil {
		panic("Cannot open file in " + filePath)
	}
	defer file.Close()

	stat, _ := file.Stat()
	bdata := make([]byte, stat.Size())
	file.Read(bdata)
	sdata := strings.Split(string(bdata), "\n")
	for _, line := range sdata {
		(*hashMap)[line] = true
	}

}
func FullOuterJoin(f1Path, f2Path, resultPath string) {
	var wg sync.WaitGroup
	
	hashMapFirst := make(map[string]bool)
	hashMapSecond := make(map[string]bool)

	wg.Add(2)
	go insertDataToMap(f1Path, &hashMapFirst, &wg)
	go insertDataToMap(f2Path, &hashMapSecond, &wg)
	wg.Wait()

	for key, _ := range hashMapFirst {
		if hashMapSecond[key] {
			// if our map1 and map2 has general key then we just delete
			// as a result we will have two map with unique elements and we just need to join
			delete(hashMapFirst, key)
			delete(hashMapSecond, key)
		}
	}

	// add unique elements from map2 to map1, as a result map1 it is our final output
	for key, _ := range hashMapSecond {
		hashMapFirst[key] = true
	}

	file, err := os.Create(resultPath)
	if err != nil {
		panic("Couldn't create file to Write our result")
	}
	defer file.Close()

	// unfortunately map isn't unordered, we have to add elements to slice for sort
	lines := make([]string, 0, 100)
	for line, _ := range hashMapFirst {
		lines = append(lines, line)
	}
	sort.Strings(lines)

	result := ""
	for index, line := range lines {
		result += line
		// we dont need add new line after last word
		if index != len(lines)-1 {
			result += "\n"
		}
	}

	_, errWrite := file.WriteString(result)
	if errWrite != nil {
		panic("couldnt write data to " + resultPath)
	}

}

// func main() {
// 	FullOuterJoin("text1.txt", "text2.txt", "output.txt")
// }
