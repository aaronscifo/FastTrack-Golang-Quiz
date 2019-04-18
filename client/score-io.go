package client

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const scoreFilePath = "scores.csv"

func SaveScoreToFile(newScore float32) error {
	var fileContent string

	fileExists, fileExistsErr := exists(scoreFilePath)
	if fileExistsErr == nil {
		if fileExists {
			existingFileContent, err := getScoreListStr()
			if err == nil && len(existingFileContent) > 0 {
				fileContent += (existingFileContent + ",")
			}
		}
		fileContent += floatToString(float64(newScore))

		writeErr := ioutil.WriteFile(scoreFilePath, []byte(fileContent), 0644)
		if writeErr != nil {
			return writeErr
		}
		return nil
	}
	return fileExistsErr
}

func GetScoreRanked(orderedScores []float32, userScore float32) int {
	for index, score := range orderedScores {
		if userScore == score {
			return index + 1
		}
	}
	return 1
}

func GetScoresList() []float32 {
	// sort.Ints(arr[:])
	fileContent, err := getScoreListStr()
	if err != nil {
		return []float32{}
	}
	fileContentStrSplit := strings.Split(fileContent, ",")

	numberOfScores := len(fileContentStrSplit)
	fileContentCSVResult := make([]float32, numberOfScores)

	for scoreIndex, scoreStr := range fileContentStrSplit {
		scoreInt, err := strconv.ParseFloat(scoreStr, 32)
		if err == nil {
			fileContentCSVResult[scoreIndex] = float32(scoreInt)
		}
	}

	return bubbleSortScoreList(fileContentCSVResult)

}

//Bubble sort implemented from https://en.wikipedia.org/wiki/Bubble_sort
func bubbleSortScoreList(scores []float32) []float32 {
	swapped := true
	for swapped {
		swapped = false
		n := len(scores)
		for i := 1; i < n; i++ {
			if scores[i-1] > scores[i] {
				originalIValue := scores[i]
				scores[i] = scores[i-1]
				scores[i-1] = originalIValue

				swapped = true
			}
		}
		n--
	}
	return scores
}

func getScoreListStr() (string, error) {
	fileContent, err := ioutil.ReadFile(scoreFilePath)
	if err != nil {
		return "", err
	}
	return string(fileContent), nil
}

func floatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

// exists returns whether the given file or directory exists
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
