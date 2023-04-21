package main

import (
	"math"
	"sort"
)

type byRespect []Validator

func (a byRespect) Len() int           { return len(a) }
func (a byRespect) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byRespect) Less(i, j int) bool { return a[i].Respect < a[j].Respect }

func findNearestValidator(validators []Validator, target int) Validator {
	sort.Sort(byRespect(validators))

	left := 0
	right := len(validators) - 1
	var nearest Validator
	minDiff := math.MaxInt32

	for left <= right {
		mid := (left + right) / 2
		diff := int(math.Abs(float64(validators[mid].Respect - target)))
		if diff < minDiff {
			minDiff = diff
			nearest = validators[mid]
		}
		if validators[mid].Respect < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if target < nearest.Respect && left > 0 {
		return validators[left-1]
	}

	return nearest
}
