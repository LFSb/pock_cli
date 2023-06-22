/*
Copyright © 2023 Lee Beenen <leebeenen@gmail.com>

*/
package cmd

import (
	"io"
	"fmt"
	"net/http"
	"net/url"
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

		fmt.Println(auth_token)
	},
}

type OauthResponse struct {
	Code string
}

func auth(consumerKey string) string {
	consumer_key := viper.Get("consumer_key").(string)

	oauth_url := "https://getpocket.com/v3/oauth/request"	

	params := url.Values{}
	params.Add("consumer_key", consumer_key)
	params.Add("redirect_uri", "pock_cli:authorizationFinished")

	response, err := http.PostForm(oauth_url, params)

	if err != nil { 
		panic(fmt.Errorf("Error when retrieving oauth token: %w", err))
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		panic(fmt.Errorf("Error when reading oauth reponse body: %w", err))
	}

	str := string(body)
	fmt.Printf("%s", str)

	// Here, I should implement Oauth2.

	// DONE:
	// 1: Retrieve the platform consumer key
	// 2: Obtain a request token: Sorta, I have a response body with a request token in there.

	//TODO:
	// 2b: Parse the response body and retrieve the "code" from there, as that's our request token that we need in further steps.
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
