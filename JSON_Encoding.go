package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Manager struct with JSON field names and omitempty
type Manager struct {
	FullName       string `json:"full_name,omitempty"`
	Position       string `json:"position,omitempty"`
	Age            int32  `json:"age,omitempty"`
	YearsInCompany int32  `json:"years_in_company,omitempty"`
}

// EncodeManager encodes a Manager to JSON and returns it as an io.Reader
func EncodeManager(manager *Manager) (io.Reader, error) {
	data, err := json.Marshal(manager)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(data), nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Read inputs from stdin (full name, position, age, years)
	fullName, _ := reader.ReadString('\n')
	position, _ := reader.ReadString('\n')
	ageStr, _ := reader.ReadString('\n')
	yearsStr, _ := reader.ReadString('\n')

	// Trim whitespace/newlines
	fullName = strings.TrimSpace(fullName)
	position = strings.TrimSpace(position)
	ageStr = strings.TrimSpace(ageStr)
	yearsStr = strings.TrimSpace(yearsStr)

	// Convert age and years to int32
	ageInt, _ := strconv.Atoi(ageStr)
	yearsInt, _ := strconv.Atoi(yearsStr)

	// Create the Manager struct
	manager := &Manager{
		FullName:       fullName,
		Position:       position,
		Age:            int32(ageInt),
		YearsInCompany: int32(yearsInt),
	}

	// Encode to JSON and print
	jsonReader, err := EncodeManager(manager)
	if err != nil {
		fmt.Println("Error encoding:", err)
		return
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(jsonReader)
	fmt.Println(buf.String())
}
