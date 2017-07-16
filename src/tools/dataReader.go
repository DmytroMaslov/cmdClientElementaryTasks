package tools

import (
	"errors"
	"os"
)

func ReadFromFile(fileName string) ([]byte, error){

	file, err := os.Open(fileName)
	if err !=nil{
		return nil,  errors.New("Can't find file")
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		return nil, errors.New("Stat Error")
	}

	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil{
		return nil, errors.New("read error")
	}
	return bs, nil
}
