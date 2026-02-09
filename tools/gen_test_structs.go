package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	pvdFolder = "<path/to/your/terraform-provider-mist>" // Adjust this path accordingly
)

type attrParameters struct {
	Required   bool
	IsListType bool
	Computed   bool
	Optional   bool
	ElemType   string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run gen_test_structs.go <resource_name>")
		fmt.Println("Example: go run gen_test_structs.go device_gateway")
		os.Exit(1)
	}

	inRes := os.Args[1]
	inFile := pvdFolder + "/internal/resource_" + inRes + "/" + inRes + "_resource_gen.go"
	outFile := pvdFolder + "/internal/provider/" + inRes + "_test_structs.go"

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

	_, err = writer.WriteString("package provider\n\nimport ()\n\n")
	if err != nil {
		fmt.Println("Error writing string:", err)
		return
	}

	isStruct := false
	nested := false
	isSchema := false
	attrLookup := make(map[string][]attrParameters)
	attrStack := stack{}
	resourceModel := "UnknownModel"

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "Model struct {") {
			isStruct = true
			modelName := strings.Split(line, " ")[1]
			resourceModel = modelName[:len(modelName)-5] // Remove last 5 letters ("Model")
		}
		if strings.Contains(line, "Value struct {") {
			isStruct = true
			nested = true
			structName := strings.Split(line, " ")[1]
			line = strings.Replace(line, structName, resourceModel+structName, 1)
		}

		if strings.Contains(line, "return schema.Schema{") {
			fmt.Println("Start schema definition")
			isSchema = true
		}

		// Parse Schema and derermine attribute parameters
		if isSchema {
			if strings.HasSuffix(strings.TrimSpace(line), "map[string]attr.Value{") {
				attrStack = attrStack.Push("ignoreMap")
			}

			if strings.HasSuffix(strings.TrimSpace(line), "[]attr.Value{") {
				attrStack = attrStack.Push("ignoreAttrList")
			}

			if strings.HasPrefix(strings.TrimSpace(line), "},") || strings.HasPrefix(strings.TrimSpace(line), "}),") || strings.HasPrefix(strings.TrimSpace(line), "})),") {
				attrStack, _ = attrStack.Pop()

				continue
			}

			if attrStack.Peek() == "ignoreMap" || attrStack.Peek() == "validators" || attrStack.Peek() == "ignoreAttrList" {
				continue
			}

			splittedLine := strings.Fields(line)
			if strings.HasPrefix(splittedLine[0], "\"") && strings.HasSuffix(splittedLine[0], ":") {
				attrName := strings.Trim(splittedLine[0], ":")
				attrName = strings.Trim(attrName, "\"")
				if attrName != "" {
					attrStack = attrStack.Push(attrName)
					if _, exists := attrLookup[attrName]; !exists {
						attrLookup[attrName] = make([]attrParameters, 0)
						attrLookup[attrName] = append(attrLookup[attrName], attrParameters{})
					}

					attrLookup[attrName] = append(attrLookup[attrName], attrParameters{})
					continue
				}
			}

			if strings.Contains(line, "Validators:") {
				attrStack = attrStack.Push("validators")
			} else if strings.Contains(line, "Attributes:") {
				attrStack = attrStack.Push("attributes")
			} else if strings.Contains(line, "NestedObject:") {
				attrStack = attrStack.Push("nestedobject")
			} else if strings.Contains(line, "CustomType:") {
				attrStack = attrStack.Push("customtype")
			} else if strings.Contains(line, "ObjectType:") {
				attrStack = attrStack.Push("objecttype")
			} else if strings.Contains(line, "PlanModifiers:") {
				attrStack = attrStack.Push("planmodifiers")
			}

			index := 0
			if attrStack.Len() > 2 {
				index = len(attrLookup[attrStack.Peek()]) - 1
			}

			if strings.Contains(line, "Required:") {
				attrLookup[attrStack.Peek()][index].Required = true
			} else if strings.Contains(line, "Optional:") {
				attrLookup[attrStack.Peek()][index].Optional = true
			} else if strings.Contains(line, "Computed:") {
				attrLookup[attrStack.Peek()][index].Computed = true
			} else if strings.Contains(line, "ElementType:") {
				splittedLine := strings.Fields(line)
				if len(splittedLine) > 1 {
					fmt.Printf("Setting ElementType for %s to %s\n", attrStack.Peek(), splittedLine[1])
					attrLookup[attrStack.Peek()][index].IsListType = true
					attrLookup[attrStack.Peek()][index].ElemType = strings.Trim(splittedLine[1], ",")
				}
			}

			if line == "}" {
				fmt.Println("End schema definition")
				isSchema = false
			}
		}

		// parse go struct and populate tags
		if isStruct {
			if !strings.Contains(line, "attr.ValueState") {
				splittedLine := strings.Fields(line)
				if len(splittedLine) == 3 {
					varName := splittedLine[0]
					varType := splittedLine[1]
					varTag := strings.Trim(splittedLine[2], "`")
					varTag = strings.Split(varTag, ":")[1]
					varTag = strings.Trim(varTag, "\"")

					attrParam := attrParameters{}
					if attrParameters, ok := attrLookup[varTag]; ok && len(attrParameters) > 0 {
						attrParam = attrParameters[0]
						if !attrParam.Required && !attrParam.Optional && !attrParam.Computed {
							attrLookup[varTag] = attrParameters[1:]
							attrParam = attrParameters[1]
						}

						if len(attrLookup[varTag]) > 1 {
							attrLookup[varTag] = attrLookup[varTag][1:]
						}
					} else {
						fmt.Printf("Warning: No parameters found for attribute %s\n", varTag)
						continue
					}

					if attrParam.Computed && !attrParam.Optional {
						fmt.Printf("Computed attribute %s skipped\n", varTag)
						continue
					}

					fullTag := fmt.Sprintf("`hcl:\"%s\"`", varTag)
					if nested {
						fullTag = fmt.Sprintf("`cty:\"%s\" hcl:\"%s\"`", varTag, varTag)
					}

					if strings.Contains(varType, "types.List") || strings.Contains(varType, "types.Set") {
						fmt.Println("Found a list type:", varName)
						fmt.Println("Element Type:", attrParam.ElemType)

						switch {
						case attrParam.ElemType == "types.StringType":
							line = fmt.Sprintf("\t%s []string %s\n", varName, fullTag)
						case attrParam.ElemType == "types.Bool":
							line = fmt.Sprintf("\t%s []bool %s\n", varName, fullTag)
						case attrParam.ElemType == "types.Float64":
							line = fmt.Sprintf("\t%s []float64 %s\n", varName, fullTag)
						case attrParam.ElemType == "types.Int64Type":
							line = fmt.Sprintf("\t%s []int64 %s\n", varName, fullTag)
						default:
							line = fmt.Sprintf("\t%s []%s%sValue %s\n", varName, resourceModel, varName, fullTag)
						}
					} else if strings.Contains(varType, "types.String") {
						if attrParam.Optional {
							line = fmt.Sprintf("\t%s *string %s\n", varName, fullTag)
						} else {
							line = fmt.Sprintf("\t%s string %s\n", varName, fullTag)
						}
					} else if strings.Contains(varType, "types.Bool") {
						if attrParam.Optional {
							line = fmt.Sprintf("\t%s *bool %s\n", varName, fullTag)
						} else {
							line = fmt.Sprintf("\t%s bool %s\n", varName, fullTag)
						}
					} else if strings.Contains(varType, "types.Float64") {
						if attrParam.Optional {
							line = fmt.Sprintf("\t%s *float64 %s\n", varName, fullTag)
						} else {
							line = fmt.Sprintf("\t%s float64 %s\n", varName, fullTag)
						}
					} else if strings.Contains(varType, "types.Int64") {
						if attrParam.Optional {
							line = fmt.Sprintf("\t%s *int64 %s\n", varName, fullTag)
						} else {
							line = fmt.Sprintf("\t%s int64 %s\n", varName, fullTag)
						}
					} else if strings.Contains(varType, "types.Object") {
						if attrParam.Optional {
							line = fmt.Sprintf("\t%s *%s%sValue %s\n", varName, resourceModel, varName, fullTag)
						} else {
							line = fmt.Sprintf("\t%s %s%sValue %s\n", varName, resourceModel, varName, fullTag)
						}
					} else if strings.Contains(varType, "types.Map") {
						switch {
						case attrParam.ElemType == "types.StringType":
							line = fmt.Sprintf("\t%s map[string]string %s\n", varName, fullTag)
						case attrParam.ElemType == "types.Bool":
							line = fmt.Sprintf("\t%s map[string]bool %s\n", varName, fullTag)
						case attrParam.ElemType == "types.Float64":
							line = fmt.Sprintf("\t%s map[string]float64 %s\n", varName, fullTag)
						case attrParam.ElemType == "types.Int64Type":
							line = fmt.Sprintf("\t%s map[string]int64 %s\n", varName, fullTag)
						default:
							line = fmt.Sprintf("\t%s map[string]%s%sValue %s\n", varName, resourceModel, varName, fullTag)
						}
					} else {
						if attrParam.Optional {
							line = fmt.Sprintf("\t%s *%s%s %s\n", varName, resourceModel, varType, fullTag)
						} else {
							line = fmt.Sprintf("\t%s %s%s %s\n", varName, resourceModel, varType, fullTag)
						}
					}
				} else {
					line = line + "\n"
				}

				if strings.HasPrefix(line, "}") {
					line = line + "\n"
					isStruct = false
				}

				_, err := writer.WriteString(line)
				if err != nil {
					fmt.Println("Error writing string:", err)
					return
				}
			}
		}
	}

	err = scanner.Err()
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
	}
}

type stack []string

func (s stack) Push(v string) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, string) {
	if len(s) == 0 {
		return s, ""
	}

	l := len(s)
	return s[:l-1], s[l-1]
}

func (s stack) Peek() string {
	if len(s) == 0 {
		return ""
	}
	return s[len(s)-1]
}

func (s stack) Len() int {
	return len(s)
}

func (s stack) IsEmpty() bool {
	return len(s) == 0
}
