package client

import (
	"bufio"
	"log"
	"os"
	"regexp"
)

func readInputs(a *admin) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		switch {
		case matchWithPattern(commandRegex, input):
			command := classifyCommand(input)
			command.handleCommand(a)
		default:
			log.Println("not command")
		}
	}
}

func classifyCommand(input string) command {
	switch {
	case matchWithPattern(commandJoinRegex, input):
		result := getParsedText(commandJoinRegex, input)
		return commandJoin{
			addr:     result["addr"],
			id:       result["id"],
			password: result["password"],
			name:     result["name"],
		}
	case matchWithPattern(commandLeaveRegex, input):
		result := getParsedText(commandLeaveRegex, input)
		return commandLeave{
			id: result["id"],
		}
	case matchWithPattern(commandFocusRegex, input):
		result := getParsedText(commandFocusRegex, input)
		return commandFocus{
			id: result["id"],
		}
	case matchWithPattern(commandSendRegex, input):
		result := getParsedText(commandSendRegex, input)
		return commandSend{
			message: result["chat"],
		}
	case matchWithPattern(commandListRegex, input):
		return commandList{}
	default:
		return commandUnknown{}
	}
}

func matchWithPattern(pattern string, text string) (match bool) {
	match, err := regexp.MatchString(pattern, text)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func getParsedText(pattern string, text string) (result map[string]string) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatal(err)
	}
	match := re.FindStringSubmatch(text)
	result = make(map[string]string)
	for idx, name := range re.SubexpNames() {
		if name != "" && idx != 0 {
			result[name] = match[idx]
		}
	}
	return
}
