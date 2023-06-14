package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client := &http.Client{}

	message := map[string]interface{}{
		"recipients": []string{
			"RECIPIENT ID HERE",
			"RECIPIENT ID HERE",
			"RECIPIENT ID HERE",
		},
	}

	messageJSON, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Error encoding message:", err)
		return
	}

	req, err := http.NewRequest("POST", "https://discord.com/api/v9/users/@me/channels", bytes.NewBuffer(messageJSON))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Authorization", "PUT YOUR TOKEN HERE")
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Connection to Discord API failed:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Discord API returned non-200 status code:", resp.StatusCode)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println(string(body))
}
