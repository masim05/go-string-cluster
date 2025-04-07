package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/masim05/go-string-cluster/pkg/scluster"
)

func main() {
	thresholdPtr := flag.Float64("threshold", 0.2, "threshold for weighted Levenstein distance")
	flag.Parse()

	input := []string{}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error while reading: ", err)
	}

	threshold := float32(*thresholdPtr)
	clusters := scluster.ClusterStringSet(input, threshold)

	fmt.Println("Line clusters:")
	for _, cluster := range clusters {
		fmt.Printf("(%d)\t: %s\n", len(cluster), cluster[0])
	}
}
