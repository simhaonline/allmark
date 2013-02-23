package parser

import (
	"andyk/docs/date"
	"andyk/docs/util"
	"strings"
	"time"
)

type MetaData struct {
	Language string
	Date     time.Time
	Tags     []string
	ItemType string
}

func (metaData MetaData) String() string {
	s := "Language: " + metaData.Language
	s += "\nDate: " + metaData.Date.String()
	s += "\nTags: " + strings.Join(metaData.Tags, ", ")
	s += "\nType: " + metaData.ItemType

	return s
}

func GetMetaData(lines []string, structure DocumentStructure) (MetaData, Match) {
	metaDataLocation := locateMetaData(lines, structure)
	if !metaDataLocation.Found {
		return MetaData{}, NotFound()
	}

	// assemble meta data
	var metaData MetaData

	for _, line := range metaDataLocation.Matches {
		isKeyValuePair, matches := util.IsMatch(line, documentStructure.MetaData)

		// skip if line is not a key-value pair
		if !isKeyValuePair {
			continue
		}

		key := strings.ToLower(strings.TrimSpace(matches[1]))
		value := strings.TrimSpace(matches[2])

		switch strings.ToLower(key) {

		case "language":
			{
				metaData.Language = value
				break
			}

		case "date":
			{
				date, err := date.ParseIso8601Date(value)
				if err == nil {
					metaData.Date = date
				}
				break
			}

		case "tags":
			{
				metaData.Tags = getTagsFromValue(value)
				break
			}

		case "type":
			{
				metaData.ItemType = value
				break
			}

		}
	}

	return metaData, metaDataLocation
}

// locateMetaData checks if the current Document
// contains meta data.
func locateMetaData(lines []string, structure DocumentStructure) Match {

	// Find the last horizontal rule in the document
	lastFoundHorizontalRulePosition := -1
	for lineNumber, line := range lines {

		lineMatchesHorizontalRulePattern := structure.HorizontalRule.MatchString(line)
		if lineMatchesHorizontalRulePattern {
			lastFoundHorizontalRulePosition = lineNumber
		}

	}

	// If there is no horizontal rule there is no meta data
	if lastFoundHorizontalRulePosition == -1 {
		return NotFound()
	}

	// If the document has no more lines than
	// the last found horizontal rule there is no
	// room for meta data
	metaDataStartLine := lastFoundHorizontalRulePosition + 1
	if len(lines) <= metaDataStartLine {
		return NotFound()
	}

	// Check if the last horizontal rule is followed
	// either by white space or be meta data
	for _, line := range lines[metaDataStartLine:] {

		lineMatchesMetaDataPattern := structure.MetaData.MatchString(line)
		if lineMatchesMetaDataPattern {

			endLine := len(lines)
			return Found(metaDataStartLine, endLine, lines[metaDataStartLine:endLine])

		}

		lineIsEmpty := structure.EmptyLine.MatchString(line)
		if !lineIsEmpty {
			return NotFound()
		}

	}

	return NotFound()
}

func getTagsFromValue(value string) []string {
	rawTags := strings.Split(value, ",")
	tags := make([]string, 0, 1)

	for _, tag := range rawTags {
		trimmedTag := strings.TrimSpace(tag)
		if trimmedTag != "" {
			tags = append(tags, trimmedTag)
		}

	}

	return tags
}