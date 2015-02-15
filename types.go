package main

import (
	"time"
)

type Language string

type Format string

const (
	English Language = "en_GB"
	Turkish Language = "tr_TR"
)

const (
	HTML     Format = "html"
	Markdown Format = "md"
	RST      Format = "rst"
)

type Entry struct {
	Title        string
	Slug         string
	Content      string
	CreationDate time.Time `bson:"creationDate"`
	UpdateDate   time.Time `bson:"updateDate"`
	Tags         []string
	Language     Language
	Format       Format
}

func fromLanguageCode(code string) Language {
	switch code {
	case "tr_TR":
		return Turkish
	case "en_US":
		return English
	case "en_GB":
		return English
	default:
		return English
	}
}

func (l Language) String() string {
	switch l {
	case Turkish:
		return "tr_TR"
	case English:
		return "en_GB"
	default:
		return ""
	}
}

func fromContentType(t string) Format {
	switch t {
	case "text/html":
		return HTML
	case "text/markdown":
		return Markdown
	case "text/x-markdown":
		return Markdown
	case "text/x-rst":
		return RST
	default:
		return HTML
	}
}

func (f Format) String() string {
	switch f {
	case HTML:
		return "text/html"
	case Markdown:
		return "text/markdown"
	case RST:
		return "text/x-rst"
	default:
		return ""
	}
}
