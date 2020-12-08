package goquerypossiblebug

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"os"
	"testing"
)

func Test_test(t *testing.T) {
	processFile(t, "./test-fixed-line.rdf.xml")
	processFile(t, "./test-removed-line.rdf.xml")
	processFile(t, "./test.rdf.xml")
}

func processFile(t *testing.T, fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	reader, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		log.Fatal(err)
	}
	result := ""
	reader.Find("pgterms\\:agent").
		ChildrenFiltered("pgterms\\:name").
		Each(func(i int, selection *goquery.Selection) {
			if result != "" {
				result += "\n"
			}
			result += selection.Text()
		})

	want := "Furniss, Harry\nFurniss, Dorothy"
	if result != want {
		t.Errorf("%v:\ngot:\n%v\nwant:\n%v", fileName, result, want)
	}
}
