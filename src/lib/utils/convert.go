package convert

import "strconv"

// FIXME: This seems incredibly lame.
func StringtoFloat64(stringArray []string) []float64 {
	var float64Array []float64

	for i := range stringArray {
		str := stringArray[i]
		num, err := strconv.ParseFloat(str, 64)

		// FIXME: Very, very, very ugly.
		if err != nil {
			return nil
		}

		float64Array = append(float64Array, num)
	}

	return float64Array
}
