package usecase

import (
	"encoding/json"
	"io"
	"os"
	"p06/domain"
)

const DefaultName = "data.txt"

type FileData struct {
	FileName string
}

func NewFileData(fileName string) FileData {
	if fileName == "" {
		fileName = DefaultName
	}

	return FileData{fileName}
}

func (f FileData) SaveProduct(product domain.Product) error {
	file, err := os.OpenFile(f.FileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(product)
}

func (f FileData) ReadProducts() ([]domain.Product, error) {
	file, err := os.Open(f.FileName)
	if err != nil {
		return []domain.Product{}, err
	}
	defer file.Close()

	products := []domain.Product{}
	decoder := json.NewDecoder(file)
	for {
		var p domain.Product
		if err := decoder.Decode(&p); err != nil {
			if err == io.EOF {
				break
			}

			return []domain.Product{}, err
		}

		products = append(products, p)
	}

	return products, nil
}
