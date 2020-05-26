package Week_01

import "math"

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	result := 0
	for l < r {
		v := calcV(l, r, height)
		result = int(math.Max(float64(v), float64(result)))

		if leight[l] > leight[r] {
			r -= 1
		} else {
			l += 1
		}
	}

	return result
}

func calcV(i int, j int, height []int) int {
	return (j - i) * int(math.Min(float64(height[i]), float64(height[j])))
}
