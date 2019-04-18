package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type Question struct {
	Question      string   `json:"question"`
	Answers       []string `json:"answers"`
	CorrectAnswer int      `json:"correctAnswer"`
}

var lazyLoadedQuestions []Question

func fetchQuestionFromJSON() ([]Question, error) {
	var questionsResult []Question

	jsonFile, err := os.Open("api/questions.json")
	if err != nil {
		return questionsResult, err
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &questionsResult)

	if err != nil {
		fmt.Printf("There was an error decoding the json. err = %s", err)
		return questionsResult, err
	}

	return questionsResult, nil
}

func loadQuestions() ([]Question, error) {
	if lazyLoadedQuestions != nil {
		return lazyLoadedQuestions, nil
	} else {
		questionsResult, err := fetchQuestionFromJSON()
		if err == nil {
			lazyLoadedQuestions = questionsResult
			return lazyLoadedQuestions, nil
		} else {
			return lazyLoadedQuestions, err
		}
	}
}

func getQuestion(questionIndex int) (Question, error) {
	var questionToReturn Question
	var errorToReturn error
	lazyLoadedQuestions, errorToReturn = loadQuestions()

	if errorToReturn == nil {
		if (questionIndex-1) >= 0 && questionIndex <= len(lazyLoadedQuestions) {
			questionToReturn = lazyLoadedQuestions[questionIndex-1]
		} else {
			errorToReturn = errors.New("Invalid Index")
		}
	}

	return questionToReturn, errorToReturn
}

func getTotalNumberOfQuestions() int {
	loadQuestions()
	if lazyLoadedQuestions != nil {
		return len(lazyLoadedQuestions)
	}
	return 0
}
