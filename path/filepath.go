package main

import (
	"path/filepath"
	"fmt"
)

func main() {
	//matches, err := filepath.Glob("/es_home/local_storage/log/(es.log)|(es_deprecation.log)|(es_index_indexing_slowlog.log)|(es_index_search_slowlog.log)|(^gc.*.current$)")
	//if err != nil {
	//	fmt.Println("Error: ", err)
	//}
	//
	//for index, matchPath := range matches {
	//	fmt.Println(index, ": ", matchPath)
	//}

	// pattern := "/es_home/local_storage/log/(es.log)|(es_deprecation.log)|(es_index_indexing_slowlog.log)|(es_index_search_slowlog.log)|(^gc.*.current$)"
	pattern1 := "['/es_home/local_storage/log/es.log', 'es_deprecation.log']"
	name := "/es_home/local_storage/log/es.log"
	matched, err := filepath.Match(pattern1, name)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	if matched {
		fmt.Println("Match...")
	} else {
		fmt.Println("Not Match...")
	}
}

