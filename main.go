package main

// [START import_libraries]
import (
	"context"
	"errors"
	"fmt"

	dialogflow "cloud.google.com/go/dialogflow/apiv2"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
)

// [END import_libraries]

// [START dialogflow_detect_intent_text]
func DetectIntentText(projectID, sessionID, text, languageCode string) (string, error) {
	ctx := context.Background()

	sessionClient, err := dialogflow.NewSessionsClient(ctx)
	if err != nil {
		return "", err
	}
	defer sessionClient.Close()

	if projectID == "" || sessionID == "" {
		return "", errors.New(fmt.Sprintf("Received empty project (%s) or session (%s)", projectID, sessionID))
	}

	sessionPath := fmt.Sprintf("projects/%s/agent/sessions/%s", projectID, sessionID)
	textInput := dialogflowpb.TextInput{Text: text, LanguageCode: languageCode}
	queryTextInput := dialogflowpb.QueryInput_Text{Text: &textInput}
	queryInput := dialogflowpb.QueryInput{Input: &queryTextInput}
	request := dialogflowpb.DetectIntentRequest{Session: sessionPath, QueryInput: &queryInput}

	response, err := sessionClient.DetectIntent(ctx, &request)
	if err != nil {
		return "", err
	}

	queryResult := response.GetQueryResult()
	fulfillmentText := queryResult.GetFulfillmentText()
	return fulfillmentText, nil
}

func main() {
	var projectID = "xxxxxxxxx"
	var sessionID = "123456789"
	var text = "I know English"
	var languageCode = "en-US"
	fulfillmentText, err := DetectIntentText(projectID, sessionID, text, languageCode)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(fulfillmentText)
}
