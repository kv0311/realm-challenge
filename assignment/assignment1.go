package assignment

import (
	"realm-challenge/utils"
	"strconv"
)

func SolveProblem1(inputRecords [][]string, partnersRecords [][]string) [][]string {
	var output [][]string
	for _, inputRow := range inputRecords {
		var partnerID string
		var mincost = 0

		for _, partnersRow := range partnersRecords[1:] {

			if utils.CheckText(inputRow[2], partnersRow[0]) && utils.CheckSizeSlab(partnersRow[1], inputRow[1]) { /* check mapping partner and theater and size slab*/
				sizeGb := utils.ToInt(inputRow[1])
				partnerMinCost := utils.ToInt(partnersRow[2])
				costPerGb := utils.ToInt(partnersRow[3])
				partnerID = partnersRow[4]
				cost := sizeGb * costPerGb /* calculate total cost */
				if cost < partnerMinCost {
					cost = partnerMinCost /*  if total cost comes less than minimum cost minimum cost will be charged. */
				}
				if mincost == 0 || cost < mincost {
					mincost = cost
				}
			}
		}

		result := make([]string, 4)
		result[0] = inputRow[0]
		if mincost == 0 { /* delivery is not possible */
			result[1] = "false"
			result[2] = ""
			result[3] = ""
		} else {
			result[1] = "true"
			result[2] = partnerID
			result[3] = strconv.Itoa(mincost)
		}

		output = append(output, result)
	}
	return output
}
