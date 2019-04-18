package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type JSONQuestionResponse struct {
	Status   int          `json:"status"`
	Question JSONQuestion `json:"question"`
}

type JSONQuestion struct {
	Question       string   `json:"question"`
	Answers        []string `json:"answer"`
	TotalQuestions int      `json:"totalQuestions"`
	QuestionIndex  int      `json:"questionIndex"`
}

type JSONError struct {
	Status   int    `json:"status"`
	Messsage string `json:"message"`
}

type JSONAnswer struct {
	CorrectAnswer     string `json:"correctAnswer"`
	AnsweredCorrectly bool   `json:"answeredCorrectly"`
}

func StartServer() {
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/question/{questionIndex}", questionRouteHandler)
	myRouter.HandleFunc("/answer/{questionIndex}/{questionAnswer}", answerRouteHandler)

	myRouter.Host("localhost")
	http.Handle("/", myRouter)

	// HTTP - port 3001
	err := http.ListenAndServe(":3001", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		fmt.Printf("ListenAndServe:%s\n", err.Error())
	}
}

func questionRouteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	questionIndex, err := strconv.Atoi(vars["questionIndex"])
	if err != nil {
		jsonError := returnJSONError("Given index is not a valid integer")
		w.Write(jsonError)
	} else {
		question, indexError := getQuestion(questionIndex)
		if indexError != nil {
			jsonError := returnJSONError(indexError.Error())
			w.Write(jsonError)
		} else {
			questionJSON := returnJSONQuestion(question, questionIndex, getTotalNumberOfQuestions())
			w.Write(questionJSON)
		}
	}
}

func answerRouteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	questionAnswer, errQuestionAns := strconv.Atoi(vars["questionAnswer"])
	questionIndex, errQuestionIndex := strconv.Atoi(vars["questionIndex"])
	if errQuestionAns != nil {
		jsonError := returnJSONError("Given question answer index is not a valid integer")
		w.Write(jsonError)
	} else if errQuestionIndex != nil {
		jsonError := returnJSONError("Given question index is not a valid integer")
		w.Write(jsonError)
	}

	question, indexError := getQuestion(questionIndex)
	if indexError != nil {
		jsonError := returnJSONError(indexError.Error())
		w.Write(jsonError)
	} else {
		questionJSON := returnJSONAnswer(question, questionAnswer)
		w.Write(questionJSON)
	}
}

func returnJSONQuestion(question Question, questionIndex int, totalQuestions int) []byte {
	jsonQuestion := JSONQuestion{Question: question.Question, Answers: question.Answers, QuestionIndex: questionIndex, TotalQuestions: totalQuestions}
	jsonResponse := JSONQuestionResponse{Status: 1, Question: jsonQuestion}

	encodedJSONResponse, err := json.Marshal(jsonResponse)
	if err == nil {
		return encodedJSONResponse
	}
	return nil
}

func returnJSONAnswer(question Question, questionAnswer int) []byte {
	var answeredCorrectly bool
	if (questionAnswer - 1) == question.CorrectAnswer { //user submits question answers starting from 1 not zero
		answeredCorrectly = true
	}

	jsonAnswer := JSONAnswer{CorrectAnswer: getQuestionCorrectAnswer(question), AnsweredCorrectly: answeredCorrectly}
	jsonData, err := json.Marshal(jsonAnswer)

	if err == nil {
		return jsonData
	}
	return nil
}

func getQuestionCorrectAnswer(question Question) string {
	result := ""
	if question.CorrectAnswer < len(question.Answers) {
		result = question.Answers[question.CorrectAnswer]
	}
	return result
}

func returnJSONError(errorStr string) []byte {
	jsonError := JSONError{Status: 0, Messsage: errorStr}
	jsonData, err := json.Marshal(jsonError)
	if err == nil {
		return jsonData
	}
	return nil
}

func main() {
	StartServer()
}
