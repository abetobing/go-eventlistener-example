package ext

import (
	"fmt"
	"time"
)

// SubmitElastic an example function which submits log to elasticsearch
func SubmitElastic(interface{}) {
	time.Sleep(2 * time.Second)
	fmt.Println("Successfully submitting log to Elasticsearch")
}
