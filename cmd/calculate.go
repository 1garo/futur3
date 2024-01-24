package cmd

import (
	"errors"
	"fmt"

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

func (y *Yaml) compoundInterest() {
	monthly := float32(y.Monthly)
	total := y.InitialInvestment

	taxPerMonth := (y.Tax / 12) / 100
	months := y.Years * 12

	for i := 1; i <= months; i++ {
		received := total * taxPerMonth
		total += received + monthly
	}

	fmt.Printf("[COMPOUND] total: %0.2f\n", total)
}

func (y *Yaml) simpleInterest() {
	tax := y.Tax / 100
	gains := y.InitialInvestment * tax
	total := y.InitialInvestment + gains * float32(y.Years)

	fmt.Printf("[SIMPLE] total: %0.2f\n", total)
}

func CalculateInterest(content []byte) error {
	var y Yaml

	if err := yaml.Unmarshal(content, &y); err != nil {
		return errCouldNotUnmarshal
	}

	if !y.IsInterestCompound {
		y.simpleInterest()
	} else {
		y.compoundInterest()
	}

	return nil
}
