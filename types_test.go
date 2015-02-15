package main

import (
	"testing"
)

func buildLanguageFromString(t *testing.T, name string, lang Language) {
	l := fromLanguageCode(name)
	if l != lang {
		t.Errorf("Built value \"%s\" is not equal to %s", l, lang)
	}
}

func buildFormatFromString(t *testing.T, name string, ct Format) {
	f := fromContentType(name)
	if f != ct {
		t.Errorf("Built value \"%s\" is not equal to %s", f, ct)
	}
}

func TestBuildingLanguageFromString(t *testing.T) {
	buildLanguageFromString(t, "tr_TR", Turkish)
	buildLanguageFromString(t, "en_GB", English)
	buildLanguageFromString(t, "en_US", English)
}

func TestBuildingUnsupportedLanguageResultsInDefault(t *testing.T) {
	buildLanguageFromString(t, "na_NE", English)
}

func TestBuildingContentTypeFromString(t *testing.T) {
	buildFormatFromString(t, "text/html", HTML)
	buildFormatFromString(t, "text/markdown", Markdown)
	buildFormatFromString(t, "text/x-markdown", Markdown)
	buildFormatFromString(t, "text/x-rst", RST)
}

func TestBuildingUnsupportedFomatResultsInDefault(t *testing.T) {
	buildFormatFromString(t, "application/octet-stream", HTML)
}

func TestLanguagesAreCorrectlyConvertedToString(t *testing.T) {
	assertEquals(t, Turkish.String(), "tr_TR")
	assertEquals(t, English.String(), "en_GB")
}

func TestFormatsAreCorrectlyConvertedToString(t *testing.T) {
	assertEquals(t, HTML.String(), "text/html")
	assertEquals(t, Markdown.String(), "text/markdown")
	assertEquals(t, RST.String(), "text/x-rst")
}
