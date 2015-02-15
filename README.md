# reblog Client

This is a simple client for adding entries to [reblog](https://github.com/rebasar/reblog) by directly writing the entries to mongodb. It is mainly written for learning a bit of [Go](https://golang.org) though it lacks most "fancy" features of the language (for example: it does not use any channels)

## Install

Assuming that your go environment is [setup correctly](https://golang.org/doc/code.html), just run:

```
go get github.com/rebasar/reblog-client
```

This will download the project and it's dependencies and install them in your `GOPATH`

## Input format

The input file is basically an e-mail in a RFC2822-like format. The difference is, it is encoded in UTF-8 and it does not support quoted printable encoding etc... You can use the following headers in the document to provide metadata:

- Subject: The title of the post
- Language: The language code for the document currently supported languages: `tr_TR`, `en_GB`, `en_US`
- Content-type: This is mandatory and specifies the format of the document. Valid values are: `text/html`, `text/makdown`, `text/x-markdown`, `text/x-rst`
- Tags: A comma separated set of strings
- Date: Instead of an RFC2822 date, this field is designed to be formatted as: 2015-02-15T11:54:39+0200

See the [sample file](testdata/complete_markdown_entry.txt) for a valid example.

For more information on how to use, write `reblog-client -h` in the command line.

Have fun!
