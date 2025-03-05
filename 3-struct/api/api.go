package api

import (
	"3-struct/config"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type ApiCreateBinReq struct {
	Name string `json:"name"`
}

type MetadataResp struct {
	Id string `json:"id"`
	CreatedAt string `json:"createdAt"`
	Private bool `json:"private"`
}

type ApiBinResp struct {
	Record ApiCreateBinReq `json:"record"`
	Metadata MetadataResp `json:"metadata"`
}

func makeHttpRequest(method, url string, body interface{}, config config.Config) (*ApiBinResp, error) {
	var jsonData []byte
	var err error

	if body != nil {
		jsonData, err = json.Marshal(body)
		if err != nil {
			return errorAnswer("error while encoding JSON")
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return errorAnswer("error while creating request")
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("X-Master-Key", config.Key)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return errorAnswer("error while executing request")
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return errorAnswer("error reading response")
	}

	if resp.StatusCode != 200 {
		return errorAnswer(fmt.Sprintf("error: answer code: %d", resp.StatusCode))
	}

	var bin ApiBinResp
	err = json.Unmarshal(bodyBytes, &bin)
	if err != nil {
		return errorAnswer("JSON decoding error")
	}

	return &bin, nil
}

func GetBin(config config.Config, idBin string) (*ApiBinResp, error) {
	baseUrl, err := getBaseUrlForBin(idBin)
	if err != nil {
		return nil, err
	}

	return makeHttpRequest("GET", baseUrl, nil, config)
}

func CreateBin(config config.Config, name string) (*ApiBinResp, error) {
	body := ApiCreateBinReq{
		Name: name,
	}

	return makeHttpRequest("POST", "https://api.jsonbin.io/v3/b/", body, config)
}

func UpdateBin(config config.Config, idBin string, name string) (*ApiBinResp, error) {
	baseUrl, err := getBaseUrlForBin(idBin)
	if err != nil {
		return nil, err
	}

	body := ApiCreateBinReq{Name: name}
	return makeHttpRequest("PUT", baseUrl, body, config)
}

func DeleteBin(config config.Config, idBin string) (error) {
	baseUrl, err := getBaseUrlForBin(idBin)
	if err != nil {
		return err
	}

	_, err = makeHttpRequest("DELETE", baseUrl, nil, config)

	return err
}

func getBaseUrlForBin(idBins string) (string, error) {
	baseUrl, err := url.Parse("https://api.jsonbin.io/v3/b/" + idBins)
	return baseUrl.String(), err
}

func errorAnswer(text string) (*ApiBinResp, error){
	return &ApiBinResp{}, fmt.Errorf("%s", text)
}
