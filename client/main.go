package client

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"quiz-cli/api"
	"strconv"
	"strings"
)

var currentQuestionIndex int = 1
var TotalNumberOfQuestions int = 1
var currentScore int = 0

func ShowNextQuestion(currentQuestionIndex int, currentScore int) (int, error) {
	userAnswer, questionErr := showQuestion(currentQuestionIndex)
	if questionErr != nil {
		fmt.Println("An error has occurred", questionErr)
		return currentScore, questionErr
	} else {
		ansResult, ansErr := GetAnswer(currentQuestionIndex, userAnswer)
		if ansErr != nil {
			fmt.Println("An error has occurred", ansErr)
			return currentScore, ansErr
		}
		if ansResult.AnsweredCorrectly {
			fmt.Println("--- You Answered correctly ----")
			currentScore++
		} else {
			fmt.Println("--- You Answered incorrectly, correct answer is " + ansResult.CorrectAnswer + " ----")
		}

		if currentQuestionIndex < TotalNumberOfQuestions {
			currentQuestionIndex++
			return ShowNextQuestion(currentQuestionIndex, currentScore)
		}
		return currentScore, nil

	}

}

func showQuestion(currentQuestionIndex int) (int, error) {
	question, err := GetQuestion(currentQuestionIndex)
	if err != nil {
		fmt.Println("An error has occurred", err)
		return 0, errors.New("Internet error")
	}

	if question.Question.TotalQuestions != TotalNumberOfQuestions {
		TotalNumberOfQuestions = question.Question.TotalQuestions
	}

	questionWithAnswers := questionAndAnswers(question, currentQuestionIndex, TotalNumberOfQuestions)

	userStrAnswer := askQuestion(questionWithAnswers)
	userAnswer, err := strconv.Atoi(userStrAnswer)

	if err != nil || userAnswer <= 0 || userAnswer > len(question.Question.Answers) {
		fmt.Println(" ==== Invalid Input (" + userStrAnswer + "), TRY AGAIN ====")
		return showQuestion(currentQuestionIndex)
	}

	return userAnswer, nil

}

func askQuestion(question string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(question)
	text, _ := reader.ReadString('\n')

	replacer := strings.NewReplacer("\n", "",
		" ", "")
	result := replacer.Replace(text)
	return result
}

func questionAndAnswers(questionResponse api.JSONQuestionResponse, questionIndex int, totalNumberOfQuestions int) string {
	result := "Question " + strconv.Itoa(questionIndex) + "/" + strconv.Itoa(totalNumberOfQuestions) + ":   " + questionResponse.Question.Question + ":"
	for index, ans := range questionResponse.Question.Answers {
		result += ("\n  " + strconv.Itoa(index+1) + ") " + ans)
	}
	return result + "\nYour Answer: "
}
