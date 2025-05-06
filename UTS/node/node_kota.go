package node

type Kota struct {
	NamaKota string
	Stasiun  *ListStasiun
}

type ListKota struct {
	Data Kota
	Link *ListKota
}
