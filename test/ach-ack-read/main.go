package main

import (
	"fmt"
	"log"
	"os"

	"github.com/moov-io/ach"
)

func main() {
	// open a file for reading. Any io.Reader Can be used
	f, err := os.Open("ack-read.ach")
	if err != nil {
		log.Fatal(err)
	}
	r := ach.NewReader(f)
	achFile, err := r.Read()
	if err != nil {
		fmt.Printf("Issue reading file: %+v \n", err)
	}
	// ensure we have a validated file structure
	if achFile.Validate(); err != nil {
		fmt.Printf("Could not validate entire read file: %v", err)
	}
	// If you trust the file but it's formatting is off building will probably resolve the malformed file.
	if achFile.Create(); err != nil {
		fmt.Printf("Could not build file with read properties: %v", err)
	}

	fmt.Printf("Credit Total Amount: %v \n", achFile.Control.TotalCreditEntryDollarAmountInFile)
	fmt.Printf("Batch Credit Total Amount: %v \n", achFile.Batches[0].GetControl().TotalCreditEntryDollarAmount)
	fmt.Printf("Total Amount: %v \n", achFile.Batches[0].GetEntries()[0].Amount)
	fmt.Printf("SEC Code: %v \n", achFile.Batches[0].GetHeader().StandardEntryClassCode)
	fmt.Printf("Original Trace Number: %v \n", achFile.Batches[0].GetEntries()[0].OriginalTraceNumberField())
}
