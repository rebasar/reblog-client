package main

import (
	"os"
	"strings"
	"testing"
	"time"
)

func TestSlugifyDoesNotLeaveDashesInTheFrontOrBack(t *testing.T) {
	s := slugify("asdf-")
	if strings.HasSuffix(s, "-") {
		t.Errorf("String \"%s\" contains a dash in the end", s)
	}
}

func TestSlugifiedStringDoesNotContainInvalidCharacters(t *testing.T) {
	s := slugify("asdf xkcd şğıç")
	if strings.ContainsAny(s, " şğıç") {
		t.Errorf("Slugified string \"%s\" contains invalid characters", s)
	}
}

func assertInt(t *testing.T, name string, real int, expected int) {
	if real != expected {
		t.Errorf("Wrong %s! Expected: %d, Got: %d", name, expected, real)
	}
}

func assertDateTime(t *testing.T, real time.Time, expected time.Time) {
	assertInt(t, "year", real.Year(), expected.Year())
	if real.Month() != expected.Month() {
		t.Errorf("Wrong month! Expected: %s, Got: %s", expected.Month(), real.Month())
	}
	assertInt(t, "day", real.Day(), expected.Day())
	assertInt(t, "hour", real.Hour(), expected.Hour())
	assertInt(t, "minute", real.Minute(), expected.Minute())
	assertInt(t, "second", real.Second(), expected.Second())
	_, realZone := real.Zone()
	_, expectedZone := expected.Zone()
	assertInt(t, "zone", realZone, expectedZone)
}

func TestValidDatesAreParsedCorrectly(t *testing.T) {
	real := parseDate("2015-02-14T23:30:14+0200")
	expected := time.Date(2015, time.February, 14, 23, 30, 14, 0, time.FixedZone("EET", 3600*2))
	assertDateTime(t, real, expected)
}

// This test might fail if you run it around midnight
func TestInvalidDateResultsInNow(t *testing.T) {
	faulDate := "2014/03/01 22:21:14 CET"
	now := time.Now()
	d := parseDate(faulDate)
	assertDate(t, d, now)
}

func TestParsedTagsShouldNotHaveWhiteSpaceAround(t *testing.T) {
	s := "tag, one, two ,three "
	tags := parseTags(s)
	for _, tag := range tags {
		if strings.HasPrefix(tag, " ") || strings.HasSuffix(tag, " ") {
			t.Errorf("Tag \"%s\" is not stripped off spaces", tag)
		}
	}
}

func TestParseTagsProducesEmptySliceIfInputIsEmpty(t *testing.T) {
	tags := parseTags("")
	if len(tags) != 0 {
		t.Errorf("Parsing empty string produced wrong result: %s, length: %d", tags, len(tags))
	}
}

func assertEquals(t *testing.T, real string, expected string) {
	if real != expected {
		t.Errorf("Expected: %s, Got: %s", expected, real)
	}
}

func TestLoadingValidMarkdownEntry(t *testing.T) {
	f, err := os.Open("testdata/complete_markdown_entry.txt")
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	entry := readEntry(f)
	assertEquals(t, entry.Title, "Test 3.14")
	assertEquals(t, entry.Slug, "test-3-14")
	assertEquals(t, entry.Language.String(), English.String())
	assertEquals(t, entry.Format.String(), Markdown.String())
}

func assertDate(t *testing.T, real time.Time, expected time.Time) {
	if real.Year() != expected.Year() {
		t.Errorf("Wrong year! Expected: %d, Got: %d", expected.Year(), real.Year())
	}
	if real.Month() != expected.Month() {
		t.Errorf("Wrong month! Expected: %s, Got: %s", expected.Month(), real.Month())
	}
	if real.Day() != expected.Day() {
		t.Errorf("Wrong day! Expected: %d, Got: %d", expected.Day(), real.Day())
	}
}

func TestLoadingWithMissingHeadersShouldProduceDefaults(t *testing.T) {
	f, err := os.Open("testdata/minimal_html_entry.txt")
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	now := time.Now()
	entry := readEntry(f)
	assertEquals(t, entry.Title, "Test 3.1415")
	assertEquals(t, entry.Slug, "test-3-1415")
	assertEquals(t, entry.Language.String(), English.String())
	assertEquals(t, entry.Format.String(), HTML.String())
	if len(entry.Tags) != 0 {
		t.Errorf("Entry contains tags: %s", entry.Tags)
	}
	assertDate(t, entry.CreationDate, now)
}
