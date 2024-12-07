# MSDS431Week9_Repo2
# Trimmed Mean Go Package

This repository provides a Go package that implements a trimmed mean function. The package is designed to compute a trimmed mean from a slice of integers or floats, where a specified proportion of elements from both ends of the sorted slice are removed. The trimmed mean is computed from the remaining elements.

## Installation

```bash
go get github.com/mvellucci100/MSDS431Wee9_TrimMedMean/trimmedmean
```
## Run the program
```bash
go run main.go
```

## Results Comparison

The trimmed mean function in this Go package is designed to be comparable to the `mean(x, trim = 0.05)` function in R, which computes the mean after trimming 5% of the observations from both ends of the sorted data.

### Go Trimmed Mean Float: 487.3469
### R Trimmed Mean Float: 527.82777
-both with 5% trimming

## Results Comparison
The trimmed mean values computed in Go and R are significantly different for the floating-point sample. The difference between the two values is approximately **40.48**. Here are a few possible reasons for the discrepancy:

### Possible Explanations:

1. **Different Data Sorting:** Both Go and R should be sorting the data before trimming, but the sorting mechanisms or optimizations in each language might result in slight variations in the sorted order, especially when working with floating-point numbers.
   
2. **Floating-Point Precision:** Go and R may handle floating-point precision slightly differently. The way each language rounds or approximates floating-point numbers during calculations might lead to discrepancies in the final results.



