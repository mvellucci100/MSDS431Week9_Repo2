package main

import (
	"fmt"
	"math/rand"
	"github.com/mvellucci100/MSDS431Week9_TrimMedMean/trimmedmean"
	"time"
	"log"
)

// Function to generate a sample of integers and floats
func generateSample(size int, isInt bool) interface{} {
	rand.Seed(time.Now().UnixNano())
	if isInt {
		// Generate a slice of integers
		sample := make([]int, size)
		for i := 0; i < size; i++ {
			sample[i] = rand.Intn(1000) // Random integer between 0 and 999
		}
		return sample
	} else {
		// Generate a slice of floats
		sample := make([]float64, size)
		for i := 0; i < size; i++ {
			sample[i] = rand.Float64() * 1000 // Random float between 0 and 1000
		}
		return sample
	}
}

func main() {
	// Generate samples
	//intSample := generateSample(100, true).([]int)
	floatSample := generateSample(100, false).([]float64)

// Call TrimmedMeanInt with 5% trimming
//trimmedMeanInt, err := trimmedmean.TrimmedMean(intSample, 0.05)
//if err != nil {
//	log.Fatalf("Error calculating trimmed mean for integers: %v", err)
//}
//fmt.Printf("Trimmed Mean (Integers): %v\n", trimmedMeanInt)

// Call TrimmedMean with 5% trimming
trimmedMeanFloat, err := trimmedmean.TrimmedMean(floatSample, 0.05)
if err != nil {
	log.Fatalf("Error calculating trimmed mean for floats: %v", err)
}
fmt.Printf("Trimmed Mean (Floats): %v\n", trimmedMeanFloat)
}
