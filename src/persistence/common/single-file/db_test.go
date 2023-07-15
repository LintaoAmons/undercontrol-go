package singlefile_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	singlefile "github.com/LintaoAmons/undercontrol/src/persistence/common/single-file"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func Test_GetTable(t *testing.T) {
	result := singlefile.NewJsonFileDB("test.json").GetTable("Test")
	for _, v := range result {
		result, _ := json.Marshal(v)
		fmt.Println(string(result))
		tmp2 := &struct {
			Name string `json:"name"`
		}{}
		json.Unmarshal(result, tmp2)
		fmt.Println(tmp2.Name)
		spew.Dump(tmp2)
		// for key, value := range v.(map[string]any) {
		// 	fmt.Println(key)
		// 	fmt.Println(value.(string))
		// }
	}
}

func Test_ConvertToTypedSlide(t *testing.T) {
	data := singlefile.NewJsonFileDB("test.json").GetTable("Test")
	spew.Dump(data)
	result := singlefile.ConvertToTypedSlide(data, &struct {
		Name string `json:"name"`
	}{})
	spew.Dump(result)
}

func TestSaveTable(t *testing.T) {
	// Create a temporary file for testing
	tmpfile, err := ioutil.TempFile("", "testdb.json")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	// Create a new JsonFileDB instance
	db := singlefile.NewJsonFileDB(tmpfile.Name())

	// Define the table and rows to save
	table := singlefile.Table("users")
	rows := []any{
		map[string]interface{}{
			"id":   1,
			"name": "John Doe",
		},
		map[string]interface{}{
			"id":   2,
			"name": "Jane Smith",
		},
	}

	// Save the table
	db.SaveTable(table, rows)

	// Read the saved data from the file
	data, err := ioutil.ReadFile(tmpfile.Name())
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	// Verify that the saved data is correct
	expected := `{"users":[{"id":1,"name":"John Doe"},{"id":2,"name":"Jane Smith"}]}`
	if string(data) != expected {
		t.Errorf("Unexpected data in file. Expected: %s, Got: %s", expected, string(data))
	}
}

type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func TestGetTable(t *testing.T) {
	// Create a temporary file for testing
	file, err := ioutil.TempFile("", "testdb.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	// Create a sample data for the test
	data := map[string][]Person{
		"table1": {
			{ID: 1, Name: "John"},
			{ID: 2, Name: "Jane"},
		},
		"table2": {
			{ID: 1, Name: "Alice"},
			{ID: 2, Name: "Bob"},
		},
	}

	// Write the sample data to the temporary file
	jsonData, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}
	err = ioutil.WriteFile(file.Name(), jsonData, 0644)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new instance of JsonFileDB
	db := singlefile.NewJsonFileDB(file.Name())

	// Call the GetTable method to retrieve the data from the "table1"
	rows := db.GetTable("table1")

	result := singlefile.ConvertToTypedSlide(rows, Person{})
	// Assert that the retrieved rows match the expected data
	expectedRows := []*Person{
		{ID: 1, Name: "John"},
		{ID: 2, Name: "Jane"},
	}
	assert.Equal(t, expectedRows, result)
}
