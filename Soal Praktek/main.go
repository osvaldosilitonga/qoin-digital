package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var pemain, dadu int

	fmt.Print("Masukkan Jumlah Pemain: ")
	fmt.Scan(&pemain)

	fmt.Print("Masukkan Jumlah Dadu: ")
	fmt.Scan(&dadu)

	playGame(pemain, dadu)
}

func playGame(pemain, dadu int) {
	data := [][]int{}
	point := make([]int, pemain)

	// for i := range data {
	// 	data[i] = []int{}
	// }

	for i := 0; i < pemain; i++ {
		data = append(data, []int{0})
		for x := 0; x < dadu; x++ {
			data[i] = append(data[i], 0)
		}
	}

	status := false
	for !status {
		data = lemparDadu(data)
		fmt.Println(data, "<------ lempar dadu")

		e, p, s := evaluasi(data, point)
		data = e

		fmt.Println(e, "<------ evaluasi")
		fmt.Println(p, "<------ point")
		fmt.Println(s, "<------ status")

		status = s
	}

}

func lemparDadu(data [][]int) [][]int {
	for i := range data {
		for x := range data[i] {
			data[i][x] = randDadu()
		}
	}

	return data
}

func evaluasi(data [][]int, point []int) ([][]int, []int, bool) {
	tmp := [][]int{}
	tmp = append(tmp, data...)

	for i := range data {
		tmp = append(tmp, []int{})
		for x, vx := range data[i] {
			if vx == 1 {
				tmp[i] = append(tmp[i][:x], tmp[i][x+1:]...) // hapus data dari slice

				if i != len(data)-1 {
					idx := i + 1
					tmp[idx] = append(tmp[idx], 1)
				} else {
					tmp[0] = append(tmp[0], 1)
				}

			}

			if vx == 6 {
				tmp[i] = append(tmp[i][:x], tmp[i][x+1:]...) // hapus data dari slice
				point[i] += 1
				if len(tmp[i]) == 0 {
					return tmp, point, true
				}
			}
		}
	}

	return tmp, point, false
}

func randDadu() int {
	time.Sleep(3 * time.Millisecond)
	randomizer := rand.New(rand.NewSource(time.Now().UnixNano()))

	ranNum := randomizer.Intn(6) + 1
	return ranNum
}
