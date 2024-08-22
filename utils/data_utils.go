package utils

import "strings"

// NormalizeLocation adapts location strings by replacing dashes and underscores with spaces.
func NormalizeLocation(location string) string {
	location = strings.ReplaceAll(location, "-", " ")
	location = strings.ReplaceAll(location, "_", " ")
	return location
}

// NormalizeLocations processes a slice of locations.
func NormalizeLocations(locations []string) []string {
	for i, location := range locations {
		locations[i] = NormalizeLocation(location)
	}
	return locations
}

// NormalizeDatesLocations processes a map of date locations.
func NormalizeDatesLocations(datesLocations map[string][]string) map[string][]string {
	for key, locations := range datesLocations {
		newKey := NormalizeLocation(key)
		datesLocations[newKey] = NormalizeLocations(locations)
		if newKey != key {
			delete(datesLocations, key)
		}
	}
	return datesLocations
}
