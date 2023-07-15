package singlefile

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Table string

type SingleFileDB interface {
	SaveTable(table Table, rows []any)
	GetTable(table Table) []any
}

func (jsonfiledb *JsonFileDB) writeToFile() error {
	data, err := json.Marshal(jsonfiledb.Data)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(jsonfiledb.filepath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func ConvertToTypedSlide[T any](rows []any, t T) []*T {
	out := make([]*T, len(rows))
	for i, v := range rows {
		result, _ := json.Marshal(v)
		var tmp *T
		json.Unmarshal(result, &tmp)
		out[i] = tmp
	}
	return out
}

func (jsonfiledb *JsonFileDB) GetTable(table Table) []any {
	jsonData, err := readFile(jsonfiledb.filepath)
	if err != nil {
		log.Fatal(err)
	}
	var result map[string][]any
	json.Unmarshal([]byte(jsonData), &result)
	return result[string(table)]
}

func (jsonfiledb *JsonFileDB) SaveTable(table Table, rows []any) {
	jsonfiledb.Data[table] = rows
	jsonfiledb.writeToFile()
}

type Data map[Table][]any
type JsonFileDB struct {
	filepath string
	Data
}

func NewJsonFileDB(filepath string) SingleFileDB {
	return &JsonFileDB{
		filepath: filepath,
		Data:     Data{},
	}
}

func readFile(filename string) ([]byte, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}
