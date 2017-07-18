package tools

import (
	"encoding/json"
	"errors"
)


func GetTaskFromFile(task string, fileName string) (data []byte, err error){
	allTask := make(map[string]json.RawMessage)
	file, er := ReadFromFile(fileName)
	if er != nil{
		return nil, er
	}
	err = json.Unmarshal(file, &allTask)
	if err != nil{
		return nil, er
	}
	data, ok := allTask[task]
	if !ok {
		return nil, errors.New("Can't find task in file")
	}
	return data, nil
}

func GetAllTaskFromFile(fileName string) (data []byte, err error){
	data, err = ReadFromFile(fileName)
	if err != nil{
		return nil, err
	}
	return data, nil
}