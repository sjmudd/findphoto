package main 

import (
	"log"
	"os"
)

func main () {
	// collect filenames from X
	// collect search directory from Y

	if len(os.Args) != 3 {
		log.Fatalf("Usage: %s <file_with_missing_filenames> <directory_to_search>", os.Args[0])
	}
	fileList := os.Args[1]
	log.Printf("Checking for files in : %q\n", fileList)
	searchPath := os.Args[2]
	log.Printf("Search path: %q\n", searchPath)

	

//	walk the tree at Y looking for files in X
//	if found
//		have seen file before?
//			yes, continue and ignore
//		check the exifinfo for the camera
//		if canon eos 70d then make a copy with checksum

}
