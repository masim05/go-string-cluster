package scluster

import (
	"github.com/texttheater/golang-levenshtein/levenshtein"
)

func weightedDistanceForStrings(a, b string) float32 {
	if len(a) == 0 && len(b) == 0 {
		return 0
	}

	distance := levenshtein.DistanceForStrings([]rune(a), []rune(b), levenshtein.DefaultOptions)
	return float32(distance) / float32(len(a)+len(b))
}

func ClusterStringSet(strings []string, threshold float32) [][]string {
	clusters := [][]string{}

	for _, str := range strings {
		foundCluster := false

		for i := range clusters {
			dist := weightedDistanceForStrings(str, clusters[i][0])
			if dist <= threshold {
				clusters[i] = append(clusters[i], str)
				foundCluster = true
				break
			}
		}

		if !foundCluster {
			clusters = append(clusters, []string{str})
		}
	}

	return clusters
}
