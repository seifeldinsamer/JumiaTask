package files

import (
	"backend/entity"
	"encoding/csv"
	"io"
	"log"
	"os"
)

const (
	idIndex           = 0
	emailIndex        = 1
	phoneNumberIndex  = 2
	parcelWeightIndex = 3
)

func ReadCSVFile(filePath string) (users []entity.User) {
	isFirstRow := true
	// Load a csv file.
	f, err := os.Open(filePath)
	checkError("Some other error occurred", err)

	defer f.Close()
	// Create a new reader.
	r := csv.NewReader(f)
	for {
		// Read row
		record, err := r.Read()

		// Stop at EOF.
		if err == io.EOF {
			break
		}

		checkError("Some other error occurred", err)

		if isFirstRow {
			isFirstRow = false
			continue
		}
		// Create new person and add to persons array
		users = append(users, entity.User{
			Id:           record[idIndex],
			Email:        record[emailIndex],
			PhoneNumber:  record[phoneNumberIndex],
			ParcelWeight: record[parcelWeightIndex],
		})
	}
	return
}

func checkError(message string, err error) {
	// Error Logging
	if err != nil {
		log.Fatal(message, err)
	}
}
