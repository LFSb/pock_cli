/*
Copyright © 2023 Lee Beenen <leebeenen@gmail.com>

*/
package cmd

import (
	"fmt"
	"bytes"
	"encoding/json"
	"net/http"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve your pocket data.",
	Long: `Retrieve your pocket data. By default you will retrieve your unread items, but eventually there's a bunch of other stuff that'll pop up here.`,
	Run: func(cmd *cobra.Command, args []string) {
		auth_token := auth("consumerKey")

		fmt.Printf("DEBUG: %s", auth_token)
	},
}

type OauthResponse struct {
	Code string
}

func get_request_token(consumerKey string) string {

	oauth_url := "https://getpocket.com/v3/oauth/request"	

	client := &http.Client{}

	bodyData := map[string]interface{}{
		"consumer_key": consumerKey,
		"redirect_uri": "pock_cli:authorizationFinished",
	}

	jsonBody, _ := json.Marshal(bodyData)	

	req, err := http.NewRequest("POST", oauth_url, bytes.NewBuffer(jsonBody)) 

	if err != nil {
		panic(fmt.Errorf("Error creating request: %w", err))
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Accept", "application/json")

	response, err := client.Do(req)

	if err != nil { 
		panic(fmt.Errorf("Error when retrieving oauth token: %w", err))
	}

	defer response.Body.Close()

	var oauth_code OauthResponse
	err = json.NewDecoder(response.Body).Decode(&oauth_code)

	if err != nil {
		panic(fmt.Errorf("Error when decoding oauth reponse body: %w", err))
	}

	request_token := oauth_code.Code

	return request_token
}

func auth(consumerKey string) string {
	consumer_key := viper.Get("consumer_key").(string)

	request_token := get_request_token(consumer_key)

	fmt.Printf("request_token: %s\n", request_token)

	// Here, I should implement Oauth2.

	// DONE:
	// 1: Retrieve the platform consumer key
	// 2: Obtain a request token: Sorta, I have a response body with a request token in there.
	// 2b: Parse the response body and retrieve the "code" from there, as that's our request token that we need in further steps.

	//TODO:
	// 3: Redirect user to Pocket to continue authorization
	// 4: Receive the callback from Pocket
	// 5: Convert a request token into a Pocket access token
	// 6: Make authenticated requests to pocket. For this we'll need the consumer_key from step 1, and an access token that we get from step 5. 

	result := "auth_token"

	return result
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
