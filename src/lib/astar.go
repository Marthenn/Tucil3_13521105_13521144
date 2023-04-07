package lib

func expandedWeight(dist []float32, euclid []float32) []float32 {
	expanded := make([]float32, len(dist))
	for i := 0; i < len(dist); i++ {
		expanded[i] = dist[i] + euclid[i]
	}
	return expanded
}
