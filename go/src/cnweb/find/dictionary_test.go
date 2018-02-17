// Unit tests for query parsing functions
package find

import (
	"log"
	"testing"
)

// Test trivial query with empty dictionary
func TestLoadDict(t *testing.T) {
	log.Printf("TestLoadDict: Begin unit tests\n")
	wdict, err := LoadDict()
	if err != nil {
		t.Error("TestLoadDict: encountered error: ", err)
		return
	}
	if len(wdict) == 0 {
		t.Error("TestLoadDict: len(wdict) == 0")
	}
	w := wdict["猴"]
	if w.HeadwordId == 0 {
		t.Error("TestLoadDict: w.HeadwordId == 0")
	}
	if w.Pinyin != "hóu" {
		t.Error("TestLoadDict: w.Pinyin != hóu", w.Pinyin)
	}
}