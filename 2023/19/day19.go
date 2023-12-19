package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"AdventOfCode/utils"
)

type Rule struct {
	Category     byte
	Operator     byte
	Num          int
	WorkflowName string
}

type Workflows map[string][]Rule

type Part struct {
	X int
	M int
	A int
	S int
}

func parseInput(input []string) (Workflows, []Part) {
	workflows := Workflows{}
	parts := []Part{}

	isWorkflow := true
	for _, line := range input {
		if line == "" {
			isWorkflow = false
		} else if isWorkflow {
			workflowName, rules := parseWorkflow(line)
			workflows[workflowName] = rules
		} else {
			part := parsePart(line)
			parts = append(parts, part)
		}
	}

	return workflows, parts
}

func parseWorkflow(line string) (string, []Rule) {
	idx := strings.Index(line, "{")

	workflowName := line[0:idx]
	rules := []Rule{}

	rulesStr := strings.Split(line[idx+1:len(line)-1], ",")
	for _, ruleStr := range rulesStr {
		rule := Rule{}
		idx = strings.Index(ruleStr, ":")
		if idx == -1 {
			rule.WorkflowName = ruleStr
		} else {
			rule.Category = byte(ruleStr[0])
			rule.Operator = byte(ruleStr[1])
			rule.Num, _ = strconv.Atoi(ruleStr[2:idx])
			rule.WorkflowName = ruleStr[idx+1:]
		}

		rules = append(rules, rule)
	}

	return workflowName, rules
}

func parsePart(line string) Part {
	var part Part
	_, err := fmt.Sscanf(line, "{x=%d,m=%d,a=%d,s=%d}", &part.X, &part.M, &part.A, &part.S)
	if err != nil {
		panic(err)
	}
	return part
}

func applyWorkflow(part Part, workflows Workflows, workflowName string) bool {
	if workflowName == "A" {
		return true
	}
	if workflowName == "R" {
		return false
	}

	for _, rule := range workflows[workflowName] {
		var rating int
		switch rule.Category {
		case 'x':
			rating = part.X
		case 'm':
			rating = part.M
		case 'a':
			rating = part.A
		case 's':
			rating = part.S
		default:
			rating = 0
		}

		var isValid = true
		switch rule.Operator {
		case '>':
			isValid = rating > rule.Num
		case '<':
			isValid = rating < rule.Num
		default:
			isValid = true
		}

		if isValid {
			return applyWorkflow(part, workflows, rule.WorkflowName)
		}
	}

	return false
}

func Part1(input []string) int {
	startWorflow := "in"

	workflows, parts := parseInput(input)

	res := 0
	for _, part := range parts {
		isValid := applyWorkflow(part, workflows, startWorflow)
		if isValid {
			res += part.X + part.M + part.A + part.S
		}
	}

	return res
}

func Part2(input []string) int {
	res := 0
	return res
}

func main() {
	fileName := "input.txt"
	input := utils.ReadFile(fileName)

	start1 := time.Now()
	fmt.Println("Answer 1 : ", Part1(input))
	fmt.Println(time.Since(start1))

	start2 := time.Now()
	fmt.Println("Answer 2 : ", Part2(input))
	fmt.Println(time.Since(start2))
}
