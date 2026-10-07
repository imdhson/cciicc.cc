package main

import (
	"html/template"
	"testing"
)

var tmplCache *template.Template
var tmplErr error

func init() {
	tmplCache, tmplErr = template.ParseFiles("wwwfiles/main.html")
}

func BenchmarkParseFiles(b *testing.B) {
	for i := 0; i < b.N; i++ {
		template.ParseFiles("wwwfiles/main.html")
	}
}

func BenchmarkCache(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = tmplCache
	}
}
