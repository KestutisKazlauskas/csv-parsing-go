package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

type category struct {
	ID          string
	Name        string
	Description string
	Items       []item
}

func (cat *category) AdItem(it item) {
	cat.Items = append(cat.Items, it)
}

type item struct {
	ID   string
	Name string
}

func main() {

	categoriesRaw, err := readCsvFile("input/categories.csv")
	if err != nil {
		log.Fatalf("Error reading categories.csv %v", err)
	}

	itemsRaw, err := readCsvFile("input/items.csv")
	if err != nil {
		log.Fatalf("Error reading categories.csv %v", err)
	}

	categories := parseCategories(categoriesRaw)
	items := parseItmes(itemsRaw)

	for _, it := range items {
		for key, cat := range categories {
			if strings.Contains(cat.Name, it.Name) {
				categories[key].AdItem(it)
				break
			}
		}
	}

	err = writeCsvFile("output/results.csv", categories)
	if err != nil {
		log.Fatalf("Error in writng result.csv: %v", err)
	}

}

func parseCategories(categoriesRaw [][]string) []category {

	cats := []category{}

	for key, categoryRaw := range categoriesRaw {
		if key == 0 {
			continue
		}
		cat := category{
			ID:          categoryRaw[0],
			Name:        categoryRaw[1],
			Description: categoryRaw[2],
		}

		cats = append(cats, cat)
	}

	return cats
}

func parseItmes(itemsRaw [][]string) []item {
	items := []item{}

	for key, itemRaw := range itemsRaw {
		if key == 0 {
			continue
		}
		it := item{
			ID:   itemRaw[0],
			Name: itemRaw[1],
		}

		items = append(items, it)
	}

	return items
}

func readCsvFile(filepath string) ([][]string, error) {
	csvFile, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	return csv.NewReader(csvFile).ReadAll()
}

func writeCsvFile(filepath string, data []category) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	//Write header
	err = writer.Write([]string{"Category", "Items"})
	if err != nil {
		return err
	}

	//Write content
	for _, cat := range data {
		for key, it := range cat.Items {
			var value []string
			if key == 0 {
				value = []string{cat.Name, it.Name}
			} else {
				value = []string{"", it.Name}
			}

			err := writer.Write(value)
			if err != nil {
				return err
			}
		}

	}
	return nil
}
