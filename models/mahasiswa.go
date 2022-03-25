package models	

type Mahasiswa struct {
	ID int `json:"id"`
	Nama string `json:"nama"`
	Prodi string `json:"prodi"`
	Fakultas string `json:"fakultas"`
	NIM int `json:"nim"`
	Angkatan int `json:"angkatan"`
}
