package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
)

func main() {
	diceGame(3, 4)
}

func diceGame(players, totalDice int) {
	fmt.Println("Pemain = ", players, ", Dadu = ", totalDice)
	dice := []int{1, 2, 3, 4, 5, 6}
	var playerPoint = map[int]int{}
	var playerPicked = map[int][]int{}
	throw := 1

	for i := 0; i < players; i++ {
		playerPoint[i] = 0
	}

	fmt.Println("======================")
	fmt.Printf("Giliran %v lempar dadu:	\n", throw)

	// first throw
	for i := 0; i < len(playerPoint); i++ {
		var picked = []int{}

		for j := 0; j < totalDice; j++ {
			randomIndex := rand.Intn(len(dice))
			pick := dice[randomIndex]
			picked = append(picked, pick)

			if pick == 6 {
				playerPoint[i] = playerPoint[i] + 1
			}
		}

		playerPicked[i] = picked

		// Format value picked array
		valuesText := []string{}
		for key := range playerPicked[i] {
			text := strconv.Itoa(playerPicked[i][key])
			valuesText = append(valuesText, text)
		}
		textPicked := strings.Join(valuesText, ",")
		fmt.Printf("\t Pemain #%v (%v): %v \n", i+1, playerPoint[i], textPicked)

	}

	// Evaluasi
	fmt.Println("Setelah evaluasi:")
	for i := 0; i < len(playerPoint); i++ {
		nextIndex := i
		if nextIndex+1 != len(playerPicked) {
			nextIndex = nextIndex + 1
		}

		// cek value 1
		if cekValueInArray(1, playerPicked[i]) && (!cekValueInArray(1, playerPicked[nextIndex])) {
			playerPicked[nextIndex] = append(playerPicked[nextIndex], 1)
		}

		// delete value 6
		playerPicked[i] = deleteInArray(6, playerPicked[i])

		// sort descending
		sort.Slice(playerPicked[i], func(x, y int) bool {
			return playerPicked[i][x] > playerPicked[i][y]
		})

		// Format value picked array
		valuesText := []string{}
		for key := range playerPicked[i] {
			text := strconv.Itoa(playerPicked[i][key])
			valuesText = append(valuesText, text)
		}
		textPicked := strings.Join(valuesText, ",")
		fmt.Printf("\t Pemain #%v (%v): %v \n", i+1, playerPoint[i], textPicked)
	}

	fmt.Println("======================")

	// 2nd throw and so on
	for {
		throw++
		fmt.Printf("Giliran %v lempar dadu:	\n", throw)

		// Picked
		for i := 0; i < len(playerPoint); i++ {
			var picked = []int{}

			for j := 0; j < len(playerPicked[i]); j++ {
				randomIndex := rand.Intn(len(dice))
				pick := dice[randomIndex]
				picked = append(picked, pick)

				if pick == 6 {
					playerPoint[i] = playerPoint[i] + 1
				}
			}

			playerPicked[i] = picked

			// Format value picked array
			valuesText := []string{}
			for key := range playerPicked[i] {
				text := strconv.Itoa(playerPicked[i][key])
				valuesText = append(valuesText, text)
			}
			textPicked := strings.Join(valuesText, ",")

			fmt.Printf("\t Pemain #%v (%v): %v \n", i+1, playerPoint[i], textPicked)

		}

		// Evaluasi
		fmt.Println("Setelah evaluasi:")
		for i := 0; i < len(playerPoint); i++ {
			nextIndex := i
			if nextIndex+1 != len(playerPicked) {
				nextIndex = nextIndex + 1
			}

			// cek value 1
			if cekValueInArray(1, playerPicked[i]) && (!cekValueInArray(1, playerPicked[nextIndex])) {
				playerPicked[nextIndex] = append(playerPicked[nextIndex], 1)
			}

			// delete value 6
			playerPicked[i] = deleteInArray(6, playerPicked[i])

			// sort descending
			sort.Slice(playerPicked[i], func(x, y int) bool {
				return playerPicked[i][x] > playerPicked[i][y]
			})

			// Format value picked array
			valuesText := []string{}
			for key := range playerPicked[i] {
				text := strconv.Itoa(playerPicked[i][key])
				valuesText = append(valuesText, text)
			}
			textPicked := strings.Join(valuesText, ",")

			fmt.Printf("\t Pemain #%v (%v): %v \n", i+1, playerPoint[i], textPicked)
		}

		fmt.Println("======================")
		var players = []int{}
		for i, v := range playerPicked {
			if len(v) > 0 {
				players = append(players, i)
			}
		}

		if len(players) < 2 {
			fmt.Printf("Game berakhir karena hanya pemain #%v yang memiliki dadu. \n", players[0]+1)
			break
		}
	}

	// get player win
	max, playerWin := 0, 0
	for index, value := range playerPoint {
		if value > max {
			max = value
			playerWin = index + 1
		}
	}

	fmt.Printf("Game dimenangkan oleh pemain #%v karena memiliki poin lebih banyak dari pemain lainnya. \n", playerWin)

}

func deleteInArray(param int, arry []int) (res []int) {
	for _, value := range arry {
		if value != param {
			res = append(res, value)
		}
	}
	return res
}

func cekValueInArray(param int, arry []int) bool {
	for _, value := range arry {
		if value == param {
			return true
		}
	}
	return false
}
