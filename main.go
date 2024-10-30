package main

import "fmt"

func GetTicketPrice(VIP, regular, student, day int) float32 {
	// Harga tiket
	const (
		vipPrice    = 30.0
		regularPrice = 20.0
		studentPrice = 10.0
	)

	// Menghitung total harga tiket
	totalPrice := float32(VIP)*vipPrice + float32(regular)*regularPrice + float32(student)*studentPrice

	// Jika total harga kurang dari $100, tidak ada diskon
	if totalPrice < 100 {
		return totalPrice
	}

	// Menghitung jumlah tiket yang dibeli
	totalTickets := VIP + regular + student
	var discount float32

	// Menentukan diskon berdasarkan tanggal dan jumlah tiket
	if day%2 == 0 { // Genap
		if totalTickets < 5 {
			discount = 0.10 // 10%
		} else {
			discount = 0.20 // 20%
		}
	} else { // Ganjil
		if totalTickets < 5 {
			discount = 0.15 // 15%
		} else {
			discount = 0.25 // 25%
		}
	}

	// Menghitung total harga setelah diskon
	totalPrice -= totalPrice * discount
	return totalPrice
}

// debug
func main() {
	fmt.Println(GetTicketPrice(1, 1, 1, 20))  // Output: 60
	fmt.Println(GetTicketPrice(3, 3, 3, 20))  // Output: 144
	fmt.Println(GetTicketPrice(4, 0, 0, 21))   // Output: 102
	fmt.Println(GetTicketPrice(0, 5, 5, 1))    // Output: 100
}
