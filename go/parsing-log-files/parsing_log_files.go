package parsinglogfiles

import (
	"strings"
	"regexp"
	"fmt"
)

func IsValidLine(text string) bool {
	if strings.HasPrefix(text,"[TRC]") || strings.HasPrefix(text,"[DBG]") || strings.HasPrefix(text,"[INF]") || strings.HasPrefix(text,"[WRN]") || strings.HasPrefix(text,"[ERR]") || strings.HasPrefix(text,"[FTL]"){
		return true
	}
	return false
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<(\**|~*|=*|-*)*>`)
	return re.Split(text,-1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	re := regexp.MustCompile(`(?i)".*password.*"`)
	for _, v := range lines {
		if re.MatchString(v) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile("end-of-line[0-9]*")
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	
	result := []string{}
	re := regexp.MustCompile(`User\s+(?P<User>[A-Za-z]+\d*)`)
	
	for _, v := range lines {
		if(re.MatchString(v)) {
			matches := re.FindStringSubmatch(v)
			userIndex := re.SubexpIndex("User")
			result = append(result, fmt.Sprintf("[USR] %s %s", matches[userIndex], v))
		} else {
			result = append(result, v)
		}
	}
	
	return result
}
