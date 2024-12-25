package utils

import (
	"log"
	"strconv"
	"strings"
)

// Utility function to parse float from string, handling special cases like %, M, B
func parseFloat(input string) float64 {
	input = strings.TrimSpace(input)

	// Handle percentage values
	if strings.HasSuffix(input, "%") {
		value, err := strconv.ParseFloat(strings.TrimSuffix(input, "%"), 64)
		if err != nil {
			log.Printf("Error parsing percentage float: %v, input: %s", err, input)
			return 0.0
		}
		return value / 100.0 // Convert percentage to decimal
	}

	// Handle values with suffix M (millions)
	if strings.HasSuffix(input, "M") {
		value, err := strconv.ParseFloat(strings.TrimSuffix(input, "M"), 64)
		if err != nil {
			log.Printf("Error parsing millions float: %v, input: %s", err, input)
			return 0.0
		}
		return value * 1e6 // Convert to numeric value
	}

	// Handle values with suffix B (billions)
	if strings.HasSuffix(input, "B") {
		value, err := strconv.ParseFloat(strings.TrimSuffix(input, "B"), 64)
		if err != nil {
			log.Printf("Error parsing billions float: %v, input: %s", err, input)
			return 0.0
		}
		return value * 1e9 // Convert to numeric value
	}

	// Fallback to standard parsing for plain numbers
	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Printf("Error parsing float: %v, input: %s", err, input)
		return 0.0
	}
	return value
}
