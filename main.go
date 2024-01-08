package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Yaml struct {
	Years        int `yaml:"years"`
	Tax          float32 `yaml:"tax_percentage"`
	InitialValue float32 `yaml:"initial_value"`
	Monthly      float32 `yaml:"monthly"`
}

func main() {
	yfile, err := os.ReadFile("params.yml")
	if err != nil {
		log.Fatal(err)
	}

	var data Yaml

	err2 := yaml.Unmarshal(yfile, &data)

	if err2 != nil {
		log.Fatal(err2)
	}

	monthly := data.Monthly
	total := data.InitialValue
	taxPerMonth := (data.Tax / 12) / 100
	months := data.Years * 12

	for i := 0; i <= months; i++ {
		receive := total * taxPerMonth
		total += receive + float32(monthly)
	}

	fmt.Printf("Initial value: $%.2f\nTotal in %d years: $%.2f\nGained $%.2f in %d years with this investment.\n",
		data.InitialValue,
		data.Years,
		total,
		total - data.InitialValue,
		data.Years,
	)

	// initial_value = 10_000
	// total = initial_value
	//
	// years = 8
	// months = years * 12
	// monthly = 2000
	//
	// # investiment tax
	// tax = 14.57
	// tax_per_month = (tax / 12) / 100
	//
	// for i in range(1, months + 1):
	//
	//	gain = initial_value * tax_per_month
	//	total += int(gain) + monthly
	//
	// print(f'total in {years} year: {int(total)} \nlucra {int(total) - initial_value} em {years} anos')
}
