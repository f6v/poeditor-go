package client

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const apiUrl = "https://poeditor.com/api/"

type Client struct {
	Config *AppConfig
}

func NewClient(config *AppConfig) *Client {
	return &Client{Config: config}
}

func (c *Client) Update() error {
	for _, langConfig := range c.Config.Languages {
		err := c.updateLanguage(langConfig)
		if err != nil {
			return err
		} else {
			fmt.Printf("Update language: %s\n", langConfig.LangCode)
		}
	}
	return nil
}

func (c *Client) updateLanguage(langConfig LanguageConfig) error {
	postParams := c.getPostParams(langConfig.LangCode)
	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(postParams.Encode()))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var dat map[string]interface{}
	if err := json.Unmarshal(body, &dat); err != nil {
		return err
	}
	downloadLink := dat["item"]
	c.dowloadFile(downloadLink.(string), langConfig.ExportPath)

	return nil
}

func (c *Client) getPostParams(languageCode string) *url.Values {
	data := url.Values{}
	data.Add("action", "export")
	data.Add("api_token", c.Config.ApiToken)
	data.Add("id", c.Config.ProjectId)
	data.Add("type", c.Config.Type)
	data.Add("language", languageCode)
	tags, _ := json.Marshal(c.Config.Tags)
	data.Add("tags", string(tags))
	return &data
}

func (c *Client) dowloadFile(downloadLink string, destination string) error {
	out, err := os.Create(destination)
	defer out.Close()

	resp, err := http.Get(downloadLink)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("Can't get file")
		return err
	}
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println("Can't save file")
		return err
	}
	fmt.Println("Downlaoded file: ", downloadLink)

	return nil
}
