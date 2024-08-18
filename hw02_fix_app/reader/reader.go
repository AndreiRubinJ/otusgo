package reader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"hw02_fix_app/employee"
)

func ReadJSON(filePath string) ([]employee.Employee, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}

	bytes, err := io.ReadAll(file)

	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}

	var data []employee.Employee
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling json: %w", err)
	}
	return data, nil
}
