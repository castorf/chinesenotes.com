/* 
Command line utility to mark up HTML files with Chinese notes.
 */
package main

import (
	"flag"
	"fmt"
	analysis "cnreader/analysis"
	config "cnreader/config"
)

//Entry point for the chinesenotes command line tool.
func main() {
	// Command line flags
	var infile = flag.String("infile", "testdata/test.html", "Input file")
	var outfile = flag.String("outfile", "testoutput/test-gloss.html",
		"Output file")
	var all = flag.Bool("all", false, "Convert all the files listed in " +
		"data/corpus/html-conversion.csv")
	flag.Parse()

	// Set project home relative to the command line tool directory
	projectHome := "../../.."
	config.SetProjectHome(projectHome)

	// Read in dictionary
	analysis.ReadDict("../../../data/words.txt")

	if !*all {
		fmt.Printf("main: input file: %s, output file: %s\n", *infile, *outfile)


		// Read text and perform vocabulary analysis
		text := analysis.ReadText(*infile)
		tokens, vocab := analysis.ParseText(text)
		analysis.WriteDoc(tokens, vocab, *outfile)
	} else {
		fmt.Printf("main: Converting all HTML files\n")
		webDir := projectHome + "/web"
		conversions := config.GetHTMLConversions()
		for _, conversion := range conversions {
			src := webDir + "/" + conversion.SrcFile
			dest := webDir + "/" + conversion.DestFile
			fmt.Printf("main: input file: %s, output file: %s\n", src, dest)
			text := analysis.ReadText(src)
			tokens, vocab := analysis.ParseText(text)
			analysis.WriteDoc(tokens, vocab, dest)
		}
	}
}