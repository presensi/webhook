package pdmk

import "github.com/whatsauth/itmodel"

// Pdm generates a reply string based on the provided IteungMessage.
func Pdmk(Pesan itmodel.IteungMessage) (reply string) {
	return ` hai berikut adalah persyaratan pdm jenis keluar:
	
Surat Pengantar sesuai dengan format surat permohonan

Ijazah dan Transkrip Nilai

SK kelulusan / SK Dikeluarkan / SK Mutasi / SK Pengunduran diri (Perubahan status mahasiswa aktif menjadi lulus/dikeluarkan/mutasi/pengundyran diri)

Pakta Integritas (Perubahan status dari keluar/mutasi/pengunduran diri ke lulus`
}
