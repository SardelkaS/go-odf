package meta

import (
	"fmt"
	"time"
)

type Meta struct {
	generator      string
	title          string
	description    string
	subject        string
	initialCreator string
	creator        string
	creationDate   time.Time
	date           time.Time
}

func New() Meta {
	return Meta{
		generator:      "MicrosoftOffice/15.0 MicrosoftWord",
		title:          "",
		description:    "",
		subject:        "",
		initialCreator: "go-odt",
		creator:        "go-odt",
		creationDate:   time.Now(),
		date:           time.Now(),
	}
}

// SetGenerator set generator metadata field
func (m Meta) SetGenerator(g string) {
	m.generator = g
}

// SetTitle set title metadata field
func (m Meta) SetTitle(t string) {
	m.title = t
}

// SetDescription set description metadata field
func (m Meta) SetDescription(d string) {
	m.description = d
}

// SetSubject set subject metadata field
func (m Meta) SetSubject(s string) {
	m.subject = s
}

// SetInitialCreator set initial creator metadata field
func (m Meta) SetInitialCreator(i string) {
	m.initialCreator = i
}

// SetCreator set creator metadata field
func (m Meta) SetCreator(c string) {
	m.creator = c
}

// SetCreationDate set creation date metadata field
func (m Meta) SetCreationDate(c time.Time) {
	m.creationDate = c
}

// SetDate set date metadata field
func (m Meta) SetDate(d time.Time) {
	m.date = d
}

// Generate generates xml code
func (m Meta) Generate() string {
	return fmt.Sprintf(
		`<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<office:document-meta xmlns:office="urn:oasis:names:tc:opendocument:xmlns:office:1.0"
    xmlns:ooo="http://openoffice.org/2004/office" xmlns:xlink="http://www.w3.org/1999/xlink"
    xmlns:dc="http://purl.org/dc/elements/1.1/"
    xmlns:meta="urn:oasis:names:tc:opendocument:xmlns:meta:1.0"
    xmlns:grddl="http://www.w3.org/2003/g/data-view#" office:version="1.4">
	<office:meta>
		<meta:generator>%s</meta:generator>
		<title>%s</title>
		<description>%s</description>
		<subject>%s</subject>
		<initial-creator>%s</initial-creator>
		<creator>%s</creator>
		<meta:creation-date>%s</meta:creation-date>
		<dc:date>%s</dc:date>
		<template xlink:href="Normal.dotm" xlink:type="simple" />
		<editing-cycles>
			1
		</editing-cycles>
		<editing-duration>
			PT60S
		</editing-duration>
	</office:meta>
</office:document-meta>`,
		m.generator, m.title, m.subject, m.subject, m.initialCreator, m.creator, m.creationDate.Format("2006-01-02T15:04:05")+"Z", m.date.Format("2006-01-02T15:04:05")+"Z")
}
