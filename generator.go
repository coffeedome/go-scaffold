package main

import (
	"log"
	"os"
)

func GenerateApiScaffold() error {

	log.Print("Generating API Scaffold...")

	err := CreateDir("scaffolding", 0755)
	if err != nil {
		return err
	}

	domains := map[string][]string{
		"animals": {"name:string", "type:string", "habitat:string", "age:string"},
	}

	err = DomainScaffolder(domains)
	if err != nil {
		return err
	}

	return nil

}

func DomainScaffolder(domains map[string][]string) error {

	for domain_name, domainAttributes := range domains {

		file, err := CreateFile(domain_name)
		if err != nil {
			return err
		}
		defer file.Close()

		for _, domainAttributes := range domainAttributes {
			err = WriteControllers(*file, domainAttributes)
			if err != nil {
				return err
			}
		}

	}

	return nil
}

func WriteControllers(file os.File, domainAttribute string) error {

	sampleController, err := os.ReadFile("templates/controller_create.go")
	if err != nil {
		return err
	}

	_, err = file.Write(sampleController)
	if err != nil {
		return err
	}

	return nil
}
