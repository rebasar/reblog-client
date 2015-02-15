package main

import (
	"io"
	"io/ioutil"
	"log"
	"mime"
	"net/mail"
	"regexp"
	"strings"
	"time"
)

func slugify(s string) string {
	base := strings.Replace(strings.ToLower(s), " ", "-", -1)
	re, err := regexp.Compile("[^a-z0-9-]")
	if err != nil {
		log.Println(err)
		return base
	}
	return strings.Trim(re.ReplaceAllString(base, "-"), "-")
}

func parseDate(date string) time.Time {
	result, err := time.Parse("2006-01-02T15:04:05-0700", date)
	if err != nil {
		log.Printf("Got error while parsing date: %s, Error: %s\n", date, err)
		log.Println("Using now instead")
		return time.Now()
	} else {
		return result
	}
}

func parseTags(tags string) []string {
	if len(tags) == 0 {
		return make([]string, 0)
	}
	result := strings.Split(tags, ",")
	for i := 0; i < len(result); i++ {
		result[i] = strings.TrimSpace(result[i])
	}
	return result
}

func readBody(body io.Reader) string {
	result, err := ioutil.ReadAll(body)
	checkError(err)
	return string(result)
}

func readEntry(in io.Reader) Entry {
	eml, err := mail.ReadMessage(in)
	checkError(err)
	mediaType, _, err := mime.ParseMediaType(eml.Header.Get("Content-Type"))
	checkError(err)
	title := eml.Header.Get("Subject")
	slug := slugify(title)
	cType := fromContentType(mediaType)
	lang := fromLanguageCode(eml.Header.Get("Language"))
	date := parseDate(eml.Header.Get("Date"))
	update := time.Now()
	tags := parseTags(eml.Header.Get("Tags"))
	body := readBody(eml.Body)
	return Entry{Title: title, Slug: slug, Content: body, CreationDate: date, UpdateDate: update, Tags: tags, Language: lang, Format: cType}
}
