package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	inRes     = "device_gateway"
	pvdFolder = "/Users/kdejong/go/src/github.com/terraform-provider-mist"
	inFile    = pvdFolder + "/internal/resource_" + inRes + "/" + inRes + "_resource_gen.go"
	outFile   = pvdFolder + "/internal/provider/" + inRes + "_test_structs.go"
	matchFile = "./test/" + inRes + "/matching.yaml"
)

type attrParameters struct {
	Required   bool
	IsListType bool
	Computed   bool
	Optional   bool
	ElemType   string
}

func main() {
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
	isSchema := false
	attrLookup := make(map[string][]attrParameters)
	attrStack := stack{}

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

		if strings.Contains(line, "return schema.Schema{") {
			fmt.Println("Start schema definition")
			isSchema = true
		}

		// Parse Schema and derermine attribute parameters
		if isSchema {
			splittedLine := strings.Fields(line)
			if strings.HasPrefix(splittedLine[0], "\"") && strings.HasSuffix(splittedLine[0], ":") {
				attrName := strings.Trim(splittedLine[0], ":")
				attrName = strings.Trim(attrName, "\"")
				if attrName != "" {
					attrStack = attrStack.Push(attrName)
					if _, exists := attrLookup[attrName]; !exists {
						attrLookup[attrName] = make([]attrParameters, 0)
					}

					attrLookup[attrName] = append(attrLookup[attrName], attrParameters{})
					// fmt.Printf("Pushed attribute %d: %s\n", len(attrLookup[attrName]), attrName)
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

			//var attrName string
			if strings.HasPrefix(strings.TrimSpace(line), "},") {
				attrStack, _ = attrStack.Pop()
				// fmt.Printf("Popped attribute: %s\n", attrName)
			}

			if strings.Contains(line, "Required:") {
				attrLookup[attrStack.Peek()][len(attrLookup[attrStack.Peek()])-1].Required = true
			} else if strings.Contains(line, "Optional:") {
				attrLookup[attrStack.Peek()][len(attrLookup[attrStack.Peek()])-1].Optional = true
			} else if strings.Contains(line, "Computed:") {
				attrLookup[attrStack.Peek()][len(attrLookup[attrStack.Peek()])-1].Computed = true
			} else if strings.Contains(line, "ElementType:") {
				splittedLine := strings.Fields(line)
				if len(splittedLine) > 1 {
					fmt.Printf("Setting ElementType for %s to %s\n", attrStack.Peek(), splittedLine[1])
					attrLookup[attrStack.Peek()][len(attrLookup[attrStack.Peek()])-1].IsListType = true
					attrLookup[attrStack.Peek()][len(attrLookup[attrStack.Peek()])-1].ElemType = strings.Trim(splittedLine[1], ",")
				}
			}

			if line == "}" {
				fmt.Println("End schema definition")
				isSchema = false
			}
		}

		// parse go struct and populate tags
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

					attrParam := attrParameters{}
					if attrParameters, ok := attrLookup[varTag]; ok && len(attrParameters) > 0 {
						attrParam = attrParameters[0]
						if len(attrParameters) > 1 {
							attrLookup[varTag] = attrParameters[1:]
						}
					} else {
						fmt.Printf("Warning: No parameters found for attribute %s\n", varTag)
						continue
					}

					if attrParam.Computed && !attrParam.Optional {
						fmt.Printf("Computed attribute %s skipped\n", varTag)
					}

					if strings.Contains(varType, "types.List") {
						fmt.Println("Found a list type:", varName)
						fmt.Println("Element Type:", attrParam.ElemType)
						switch {
						case attrParam.ElemType == "types.StringType":
							line = fmt.Sprintf("\t%s []string `%s:\"%s\"`\n", varName, tag, varTag)
						case attrParam.ElemType == "types.Bool":
							line = fmt.Sprintf("\t%s []bool `%s:\"%s\"`\n", varName, tag, varTag)
						case attrParam.ElemType == "types.Float64":
							line = fmt.Sprintf("\t%s []float64 `%s:\"%s\"`\n", varName, tag, varTag)
						case attrParam.ElemType == "types.Int64Type":
							line = fmt.Sprintf("\t%s []int64 `%s:\"%s\"`\n", varName, tag, varTag)
						default:
							line = fmt.Sprintf("\t%s []%sValue `%s:\"%s\"`\n", varName, varName, tag, varTag)
						}
					} else if strings.Contains(varType, "types.String") {
						if attrParam.Optional && !attrParam.Computed {
							line = fmt.Sprintf("\t%s *string `%s:\"%s\"`\n", varName, tag, varTag)
						} else {
							line = fmt.Sprintf("\t%s string `%s:\"%s\"`\n", varName, tag, varTag)
						}
					} else if strings.Contains(varType, "types.Bool") {
						line = fmt.Sprintf("\t%s bool `%s:\"%s\"`\n", varName, tag, varTag)
					} else if strings.Contains(varType, "types.Float64") {
						line = fmt.Sprintf("\t%s float64 `%s:\"%s\"`\n", varName, tag, varTag)
					} else if strings.Contains(varType, "types.Int64") {
						line = fmt.Sprintf("\t%s int64 `%s:\"%s\"`\n", varName, tag, varTag)
					} else if strings.Contains(varType, "types.Object") {
						line = fmt.Sprintf("\t%s %sValue `%s:\"%s\"`\n", varName, varName, tag, varTag)
					} else if strings.Contains(varType, "types.Map") {
						switch {
						case attrParam.ElemType == "types.StringType":
							line = fmt.Sprintf("\t%s map[string]string `%s:\"%s\"`\n", varName, tag, varTag)
						case attrParam.ElemType == "types.Bool":
							line = fmt.Sprintf("\t%s map[string]bool `%s:\"%s\"`\n", varName, tag, varTag)
						case attrParam.ElemType == "types.Float64":
							line = fmt.Sprintf("\t%s map[string]float64 `%s:\"%s\"`\n", varName, tag, varTag)
						case attrParam.ElemType == "types.Int64Type":
							line = fmt.Sprintf("\t%s map[string]int64 `%s:\"%s\"`\n", varName, tag, varTag)
						default:
							line = fmt.Sprintf("\t%s map[string]%sValue `%s:\"%s\"`\n", varName, varName, tag, varTag)
						}
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
