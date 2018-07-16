// Unit tests for find functions
package find

import (
	"log"
	"testing"
)

// Test package initialization, which requires a database connection
func TestInit(t *testing.T) {
	log.Printf("TestInit: Begin unit tests\n")
}

func TestFindDocuments1(t *testing.T) {
	dict := map[string]Word{}
	parser := DictQueryParser{dict}
	qr, err := FindDocuments(parser, "Assembly")
	if err != nil {
		t.Error("TestFindDocuments1: got error, ", err)
	}
	if len(qr.Terms) != 1 {
		t.Error("TestFindDocuments1: len(qr.Terms) != 1, ", qr)
	}
}

func TestFindDocuments2(t *testing.T) {
	dict := map[string]Word{}
	parser := DictQueryParser{dict}
	_, err := FindDocuments(parser, "")
	if err == nil {
		t.Error("TestFindDocuments2: expected error for empty string")
	}
}

func TestFindDocuments3(t *testing.T) {
	dict := map[string]Word{}
	parser := DictQueryParser{dict}
	qr, err := FindDocuments(parser, "hello")
	if err != nil {
		t.Error("TestFindDocuments3: got error, ", err)
	}
	if len(qr.Terms) != 1 {
		t.Error("TestFindDocuments3: len(qr.Terms) != 1, ", qr)
	}
	if len(qr.Terms[0].Senses) == 0 {
		t.Error("TestFindDocuments3: len(qr.Terms.Senses) == 0, ", qr)
	}
}

func TestFindInBody0(t *testing.T) {
	terms := []string{}
	docSimilarity, err := findInBody(terms)
	if err == nil {
		t.Error("TestFindInBody0: expected an error, ", err)
		return
	}
	log.Printf("TestFindInBody0, len(docSimilarity) = %d", len(docSimilarity))
}

func TestFindInBody1(t *testing.T) {
	terms := []string{"后妃"}
	docSimilarity, err := findInBody(terms)
	if err == nil {
		t.Error("TestFindInBody0: expected an error, ", err)
		return
	}
	log.Printf("TestFindInBody1, len(docSimilarity) = %d", len(docSimilarity))
}

func TestFindInBody2(t *testing.T) {
	terms := []string{"后妃", "之"}
	docSimilarity, err := findInBody(terms)
	if err != nil {
		t.Error("TestFindInBody: got error, ", err)
	}
	log.Printf("TestFindInBody2, len(docSimilarity) = %d", len(docSimilarity))
}

func TestFindInBody3(t *testing.T) {
	terms := []string{"后妃", "之", "德"}
	docSimilarity, err := findInBody(terms)
	if err != nil {
		t.Error("TestFindInBody: got error, ", err)
	}
	log.Printf("TestFindInBody3, len(docSimilarity) = %d", len(docSimilarity))
}

func TestFindInBody4(t *testing.T) {
	terms := []string{"后妃", "之", "德", "也"}
	docSimilarity, err := findInBody(terms)
	if err != nil {
		t.Error("TestFindInBody: got error, ", err)
	}
	log.Printf("TestFindInBody4, len(docSimilarity) = %d", len(docSimilarity))
}

func TestFindWords1(t *testing.T) {
	words, err := findWords("Assembly")
	if err != nil {
		t.Error("TestFindWords1: got error, ", err)
	}
	if len(words) != 0 {
		t.Error("TestFindWords1: len(words) != 0, ", len(words))
	}
}

func TestFindWords2(t *testing.T) {
	words, err := findWords("金剛")
	if err != nil {
		t.Error("TestFindWords2: got error, ", err)
	}
	if len(words) != 1 {
		t.Error("TestFindWords2: len(words) != 1, ", len(words))
	}
}

