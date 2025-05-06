package node

type Stasiun struct {
	Nama   string
	Alamat string
	Notelp string
}

type ListStasiun struct {
	Data Stasiun
	Link *ListStasiun
}
