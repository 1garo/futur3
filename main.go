package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	errCouldNotUnmarshal = errors.New("Could not unmarshal the content of the file.")
)

type Yaml struct {
	Years        int `yaml:"years"`
	Monthly      float32 `yaml:"monthly"`

	InitialInvestment float32 `yaml:"initial_investment"`

	IsInterestCompound bool `yaml:"is_interest_compound"`
	Tax          float32 `yaml:"tax_percentage"`
}

func compoundInterest(data *Yaml) string {
	monthly := float32(data.Monthly)
	total := data.InitialInvestment

	taxPerMonth := (data.Tax / 12) / 100
	months := data.Years * 12

	for i := 1; i <= months; i++ {
		received := total * taxPerMonth
		total += received + monthly
	}

	// TODO: refactor this
	return fmt.Sprintf("Initial investment: $%.2f\nTotal in %d years: $%.2f\nEarned $%.2f in %d years with this investment.",
		data.InitialInvestment,
		data.Years,
		total,
		total - data.InitialInvestment,
		data.Years,
	)
}

func simpleInterest(data *Yaml) string {
	tax := data.Tax / 100
	gains := data.InitialInvestment * tax
	total := data.InitialInvestment + gains * float32(data.Years)

	return fmt.Sprintf("Initial investment: $%.2f\nTotal in %d years: $%.2f\nEarned $%.2f in %d years with this investment.",
		data.InitialInvestment,
		data.Years,
		total,
		total - data.InitialInvestment,
		data.Years,
	)
}

func calculateInterest(content []byte) (string, error) {
	var data Yaml

	if err := yaml.Unmarshal(content, &data); err != nil {
		return "", errCouldNotUnmarshal
	}

	if !data.IsInterestCompound {
		return simpleInterest(&data), nil
	}

	return compoundInterest(&data), nil
}

func main() {
	yfile, err := os.ReadFile("example.yml")
	if err != nil {
		log.Fatal(err)
	}

	print, err := calculateInterest(yfile);
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(print)
}
