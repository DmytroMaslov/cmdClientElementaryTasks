package tools

import (
	"net/http"
	"bytes"
	"io/ioutil"
)


func SendRequest (URL string, data []byte) ([]byte, error){
	
	body := bytes.NewBuffer(data)
	resp, err :=http.Post(URL, "raw", body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err!= nil{
		return nil, err
	}
	return respBody, err

}
