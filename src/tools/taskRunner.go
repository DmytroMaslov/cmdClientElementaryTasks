package tools

import (
	"encoding/json"
	"fmt"
)

const URL string = "http://localhost:8080"
const ALL_TASK string = "/tasks"
const ONE_TASK string = "/task"
const SEPARATOR  = "/"

type Respons struct {
	Task string `json: task`
	Resp string `json:"resp"`
	Reason string `json:"reason"`
}
func RunTask (task string, file string) (map[string][]string, error){

	data, er := ReadFromFile(file)
	if er != nil{
		return nil, er
	}
	var resp = new(Respons)
	fUrl := URL + ONE_TASK + SEPARATOR + task
	rs, err := SendRequest(fUrl, data)
	if err != nil{
		return nil, er
	}
	err = json.Unmarshal(rs, &resp)
	if err != nil{
		return nil, er
	}
	m:=make(map[string][]string)
	m [resp.Task]= []string{
		resp.Task,
		resp.Resp,
		resp.Reason,
	}
	return m, nil
}

func RunAllTasks(file string) (map[string][]string, error) {
	data, err := ReadFromFile(file)
	if err != nil{
		return nil, err
	}
	fUrl := URL + ALL_TASK
	rs, err := SendRequest(fUrl, data)
	if err != nil{
		return nil, err
	}
	res := make([]Respons,0)
	er:=json.Unmarshal(rs, &res)
	if er!=nil{
		fmt.Println(er.Error())
	}
	m:=make(map[string][]string)
	for _, v := range res{
		m[v.Task] = []string{
			v.Task,
			v.Resp,
			v.Reason,
		}
	}
	return m, nil


}




