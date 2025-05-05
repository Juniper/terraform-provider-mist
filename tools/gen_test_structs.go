package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	inRes     = "device_gateway"
	pvdFolder = "/Users/kdejong/go/src/github.com/terraform-provider-mist"
	inFile    = pvdFolder + "/internal/resource_" + inRes + "/" + inRes + "_resource_gen.go"
	outFile   = pvdFolder + "/internal/provider/" + inRes + "_test_structs.go"
	matchFile = "./test/" + inRes + "/matching.yaml"
)

var customMatches map[string]string

func loadCustomMatches() error {
	file, err := os.Open(matchFile)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	return decoder.Decode(&customMatches)
}

func main() {
	err := loadCustomMatches()
	if err != nil {
		fmt.Printf("Error loading custom matches: %v\n", err)
	}

	in, err := os.Open(inFile)
	if err != nil {
		fmt.Printf("Error opening input file: %v\n", err)
		return
	}
	defer in.Close()

	out, err := os.Create(outFile)
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		return
	}
	defer out.Close()

	writer := bufio.NewWriter(out)
	defer writer.Flush()

	writer.WriteString("package provider\n\nimport ()\n\n")

	isStruct := false
	nested := false

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "Model struct {") {
			isStruct = true
		}
		if strings.Contains(line, "Value struct {") {
			isStruct = true
			nested = true
		}

		if isStruct {
			tag := "hcl"
			if nested {
				tag = "cty"
			}

			if !strings.Contains(line, "attr.ValueState") {
				splittedLine := strings.Fields(line)
				if len(splittedLine) == 3 {
					varName := splittedLine[0]
					varType := splittedLine[1]
					varTag := strings.Trim(splittedLine[2], "`")
					varTag = strings.Split(varTag, ":")[1]
					varTag = strings.Trim(varTag, "\"")

					if customType, ok := customMatches[varName]; ok {
						line = fmt.Sprintf("\t%s %s `%s:\"%s\"`\n", varName, customType, tag, varTag)
					} else if strings.Contains(varType, "types.String") {
						line = fmt.Sprintf("\t%s string `%s:\"%s\"`\n", varName, tag, varTag)
					} else if strings.Contains(varType, "types.Bool") {
						line = fmt.Sprintf("\t%s bool `%s:\"%s\"`\n", varName, tag, varTag)
					} else if strings.Contains(varType, "types.Float64") {
						line = fmt.Sprintf("\t%s float64 `%s:\"%s\"`\n", varName, tag, varTag)
					} else if strings.Contains(varType, "types.Int64") {
						line = fmt.Sprintf("\t%s int64 `%s:\"%s\"`\n", varName, tag, varTag)
					} else if strings.Contains(varType, "types.List") {
						line = fmt.Sprintf("\t%s []%sValue `%s:\"%s\"`\n", varName, varName, tag, varTag)
					} else if strings.Contains(varType, "types.Object") {
						line = fmt.Sprintf("\t%s %sValue `%s:\"%s\"`\n", varName, varName, tag, varTag)
					} else if strings.Contains(varType, "types.Map") {
						line = fmt.Sprintf("\t%s map[string]%sValue `%s:\"%s\"`\n", varName, varName, tag, varTag)
					} else {
						line = fmt.Sprintf("\t%s %s `%s:\"%s\"`\n", varName, varType, tag, varTag)
					}
				} else {
					line = line + "\n"
				}

				if strings.HasPrefix(line, "}") {
					line = line + "\n"
					isStruct = false
				}

				writer.WriteString(line)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
	}
}
