package main

type bill struct {
	nama string
	item map[string]float64
	tip  float64
}

// buat bill baru

func newBill(nama string) bill {
	b := bill{
		nama: nama,
		item: map[string]float64{},
		tip:  0,
	}
	return b
}
