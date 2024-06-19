package pdm

import "github.com/whatsauth/itmodel"

// Pdm generates a reply string based on the provided IteungMessage.
func Pdm(Pesan itmodel.IteungMessage) (reply string) {
	return `hai berikut adalah persayaratan pdm pokok jika merubah:
	
Nomor Induk Mahasiswa: 
1.Kartu Tanda Mahasiswa
2.Ijazah dan Transkrip Nilai (Apabila Mahasiswa sudah lulus),
3.Kartu Hasil Studi

Nama Mahasiswa: 
1.Akte Kelahiran atau Surat Kenal Lahir atau Kartu Keluarga atau Ijazah
2.Kartu Tanda Mahasiswa 
3.Ijazah dan Transkrip Nilai (Apabila Mahasiswa sudah lulus)

Nama Ibu Kandung: 
1.Akte Kelahiran atau Surat Kenal Lahir atau Kartu Keluarga

Tempat Lahir:
1.Akte Kelahiran atau Surat Kenal Lahir atau Kartu Keluarga atau Ijazah
2.Kartu Tanda Mahasiswa
3.Ijazah dan Transkrip Nilai (Apabila Mahasiswa sudah lulus)

Tanggal Lahir:
1.Akte Kelahiran atau Surat Kenal Lahir atau Kartu Keluarga atau Ijazah
2.Kartu Tanda Mahasiswa, 
3.Ijazah dan Transkrip Nilai (Apabila Mahasiswa sudah lulus)

Periode Pendaftaran: Surat Penerimaan Mahasiswa

Jenis Kelamin: Mengikuti Persyaratan Umum`
}
