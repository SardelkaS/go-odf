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
		generator:      "SardelkaS/go-odf",
		title:          "",
		description:    "",
		subject:        "",
		initialCreator: "go-odt",
		creator:        "go-odt",
		creationDate:   time.Now(),
		date:           time.Now(),
	}
}

func (m Meta) SetGenerator(g string) {
	m.generator = g
}

func (m Meta) SetTitle(t string) {
	m.title = t
}

func (m Meta) SetDescription(d string) {
	m.description = d
}

func (m Meta) SetSubject(s string) {
	m.subject = s
}

func (m Meta) SetInitialCreator(i string) {
	m.initialCreator = i
}

func (m Meta) SetCreator(c string) {
	m.creator = c
}

func (m Meta) SetCreationDate(c time.Time) {
	m.creationDate = c
}

func (m Meta) SetDate(d time.Time) {
	m.date = d
}

func (m Meta) Generate() string {
	return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<document-meta xmlns:office="urn:oasis:names:tc:opendocument:xmlns:office:1.0" xmlns:meta="urn:oasis:names:tc:opendocument:xmlns:meta:1.0" xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:xlink="http://www.w3.org/1999/xlink" office:version="1.3">
	<meta>
		<generator>%s</generator>
		<title>%s</title>
		<description>%s</description>
		<subject>%s</subject>
		<initial-creator>%s</initial-creator>
		<creator>%s</creator>
		<creation-date>%s</creation-date>
		<date>%s</date>
		<template xlink:href="Normal.dotm" xlink:type="simple" />
		<editing-cycles>
			1
		</editing-cycles>
		<editing-duration>
			PT60S
		</editing-duration>
	</meta>
</document-meta>`,
		m.generator, m.title, m.subject, m.subject, m.initialCreator, m.creator, m.creationDate.Format(time.RFC3339), m.date.Format(time.RFC3339))
}
