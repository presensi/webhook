package pdm

import "github.com/whatsauth/itmodel"

// Pdm generates a reply string based on the provided IteungMessage.
func Pdm(Pesan itmodel.IteungMessage) (reply string) {
	return `Nomor Induk Mahasiswa: Kartu Tanda Mahasiswa, Ijazah dan Transkrip Nilai (Apabila Mahasiswa sudah lulus), Kartu Hasil Studi

Nama Mahasiswa: Akte Kelahiran atau Surat Kenal Lahir atau Kartu Keluarga atau Ijazah, Kartu Tanda Mahasiswa, Ijazah dan Transkrip Nilai (Apabila Mahasiswa sudah lulus)

Nama Ibu Kandung: Akte Kelahiran atau Surat Kenal Lahir atau Kartu Keluarga

Tempat Lahir: Akte Kelahiran atau Surat Kenal Lahir atau Kartu Keluarga atau Ijazah, Kartu Tanda Mahasiswa, Ijazah dan Transkrip Nilai (Apabila Mahasiswa sudah lulus)

Tanggal Lahir: Akte Kelahiran atau Surat Kenal Lahir atau Kartu Keluarga atau Ijazah, Kartu Tanda Mahasiswa, Ijazah dan Transkrip Nilai (Apabila Mahasiswa sudah lulus)

Periode Pendaftaran: Surat Penerimaan Mahasiswa

Jenis Kelamin: Mengikuti Persyaratan Umum`
}
