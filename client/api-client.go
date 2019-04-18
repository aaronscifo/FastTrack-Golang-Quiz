package client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"quiz-cli/api"
	"strconv"
)

type API_ENDPOINT int

const API_URL = "http://localhost:3001"

const ( // iota is reset to 0
	getQuestionsAPIEnpoint API_ENDPOINT = iota // c0 == 0
	getAnswerAPIEnpoint                        // c1 == 1
)

func GetQuestion(questionIndex int) (api.JSONQuestionResponse, error) {
	// url := getAPIEnpointUrl(getQuestionsAPIEnpoint) + strconv.Itoa(questionIndex)
	// result, err := sendServerRequest(url)
	// if err != nil {
	// 	return JSONQuestion{}, err
	// }
	// var jsonResult JSONQuestion
	// err = json.Unmarshal(result, &jsonResult)
	// if err != nil {
	// 	return JSONQuestion{}, err
	// }
	// return jsonResult, nil

	url := getAPIEnpointURL(getQuestionsAPIEnpoint) + strconv.Itoa(questionIndex)
	result, err := sendServerRequest(url)
	if err != nil {
		return api.JSONQuestionResponse{}, err
	} else {
		var jsonResult api.JSONQuestionResponse
		err = json.Unmarshal(result, &jsonResult)
		if err != nil {
			return jsonResult, err
		}
		return jsonResult, nil
	}

}

func GetAnswer(questionIndex int, ansIndex int) (api.JSONAnswer, error) {
	url := getAPIEnpointURL(getAnswerAPIEnpoint) + strconv.Itoa(questionIndex) + "/" + strconv.Itoa(ansIndex)
	result, err := sendServerRequest(url)
	if err != nil {
		return api.JSONAnswer{}, err
	} else {
		var jsonResult api.JSONAnswer
		err = json.Unmarshal(result, &jsonResult)
		if err != nil {
			return api.JSONAnswer{}, err
		}
		return jsonResult, nil
	}
}

func sendServerRequest(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		// resultStr := string(data)
		// fmt.Println(resultStr)
		return data, nil
	}
}

func getAPIEnpointURL(endpoint API_ENDPOINT) string {
	url := API_URL
	switch endpoint {
	case getQuestionsAPIEnpoint:
		url += "/question/"
		break
	case getAnswerAPIEnpoint:
		url += "/answer/"
		break
	}
	return url
}
