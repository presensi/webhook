package nidn

import "github.com/whatsauth/itmodel"

// Pdm generates a reply string based on the provided IteungMessage.
func Nidn(Pesan itmodel.IteungMessage) (reply string) {
	return ` hai berikut adalah persyaratan umum regristasi dosen baru:
	
1.Surat Keterangan Sehat Jasmani Rohani, dan Bebas
Narkoba 
2.SK Dosen Tetap
3.Ijazah
4.Foto Formal
5.KTP
6.Surat Penyataan Pemimpin PT`
}
