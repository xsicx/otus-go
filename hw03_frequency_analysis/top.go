package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

type wordWithFrequency struct {
	word      string
	frequency int
}

func Top10(str string) []string {
	if len(str) == 0 {
		return make([]string, 0)
	}
	str = clearSpaces(str)
	str = prepareByExtraRules(str)
	wordsWithFrequency := frequencyAnalyze(str)

	return getSortedWordsByLimit(wordsWithFrequency, 10)
}

func clearSpaces(str string) string {
	space := regexp.MustCompile(`\s+`)
	return space.ReplaceAllString(str, " ")
}

// Преобразования строки для задания со звездочкой.
func prepareByExtraRules(str string) string {
	str = strings.ToLower(str)
	str = strings.ReplaceAll(str, " - ", " ")
	reg := regexp.MustCompile(`[.?!)(,:]+`)
	return reg.ReplaceAllString(str, "")
}

func frequencyAnalyze(str string) []wordWithFrequency {
	uniqueWords := make(map[string]int)
	for _, word := range strings.Fields(str) {
		if _, exists := uniqueWords[word]; exists {
			uniqueWords[word]++
			continue
		}
		uniqueWords[word] = 1
	}

	wordsWithFrequency := make([]wordWithFrequency, 0, len(uniqueWords))

	for word, frequency := range uniqueWords {
		wordsWithFrequency = append(wordsWithFrequency, wordWithFrequency{word, frequency})
	}

	sort.Slice(wordsWithFrequency, func(a, b int) bool {
		if wordsWithFrequency[a].frequency == wordsWithFrequency[b].frequency {
			return wordsWithFrequency[a].word < wordsWithFrequency[b].word
		}
		return wordsWithFrequency[a].frequency > wordsWithFrequency[b].frequency
	})

	return wordsWithFrequency
}

func getSortedWordsByLimit(wordsWithFrequency []wordWithFrequency, limit int) []string {
	res := make([]string, 0)

	if limit < 1 {
		return res
	}

	for i, data := range wordsWithFrequency {
		if i > limit-1 {
			break
		}
		res = append(res, data.word)
	}

	return res
}
