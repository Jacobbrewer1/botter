package helper

import (
	"errors"
	"log"
	"regexp"
	"strings"
)

func FormatIdFromMention(id string) (string, error) {
	id = RemoveMultiSpaces(id)
	if strings.Contains(id, " ") {
		return "", errors.New("id input contains string")
	}
	if strings.Count(id, "@") > 1 {
		return "", errors.New("multiple id's passed")
	}
	if strings.Contains(id, "<") {
		id = strings.TrimLeft(id, "<")
	}
	if strings.Contains(id, "@") {
		id = strings.TrimLeft(id, "@")
	}
	if strings.Contains(id, "!") {
		id = strings.TrimLeft(id, "!")
	}
	if strings.Contains(id, "&") {
		id = strings.TrimLeft(id, "&")
	}
	if strings.Contains(id, ">") {
		id = strings.TrimRight(id, ">")
	}
	return strings.TrimSpace(id), nil
}

func FormatMessage(input, prefix string) (string, string) {
	var returnQuery string
	input = RemoveMultiSpaces(strings.TrimSpace(input))
	command := strings.TrimLeft(input, prefix)
	x := strings.Split(command, " ")
	command = strings.ToLower(x[0])
	if len(x) > 1 {
		var y []string
		for pos, elm := range x {
			if pos == 0 {
				continue
			}
			y = append(y, elm)
		}
		returnQuery = strings.TrimSpace(strings.Join(y, " "))
	} else {
		returnQuery = ""
	}
	return command, returnQuery
}

func RemoveMultiSpaces(input string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(input)), " ")
}

func RemoveNonAlphaChars(text string) string {
	// Regex to say we only want letters and numbers
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Println(err)
		return ""
	}
	// Removing double spaces in the string after the conversion. This will prevent weird formatting
	return RemoveMultiSpaces(reg.ReplaceAllString(text, " "))
}

func FormatStrings(text string) string {
	// Regex to say we only want letters and numbers
	reg, err := regexp.Compile("[^a-zA-Z0-9,]+")
	if err != nil {
		log.Println(err)
		return ""
	}
	return strings.ToLower(RemoveMultiSpaces(reg.ReplaceAllString(RemoveNewLines(RemoveTab(text)), "")))
}

func RemoveNewLines(text string) string {
	return RemoveMultiSpaces(strings.ReplaceAll(text, "\n", " "))
}

func RemoveTab(text string) string {
	return RemoveMultiSpaces(strings.ReplaceAll(text, "\t", " "))
}

func RemoveBoldness(text string) string {
	return strings.ReplaceAll(text, "*", "")
}
