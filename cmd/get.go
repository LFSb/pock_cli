/*
Copyright © 2023 Lee Beenen leebeenen@gmail.com

*/
package cmd

import (
	"fmt"

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

func auth(consumerKey string) string {
	fmt.Println("This is where auth will happen, someday.")

	result := "auth_token"

	return result

	// Here, I should implement Oauth2.

	// Step1: load the platform consumer key.

	// Step 2: Obtain a request token

	// Step 3: Redirect user to Pocket to continue authorization

	// Step 4: Receive the callback from Pocket

	// Step 5: Convert a request token into a Pocket access token

	// Step 6: Make authenticated requests to pocket. For this we'll need the consumer_key from step 1, and an access token that we get from step 5. 
}

func init() {
	rootCmd.AddCommand(getCmd)

	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/pock_cli/")   // path to look for the config file in
	viper.AddConfigPath("$HOME/.pock_cli/")  // call multiple times to add many search paths
	viper.AddConfigPath(".")               // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	consumer_key := viper.Get("consumer_key")
	fmt.Printf("loaded consumer_key: '%s'\n", consumer_key)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
