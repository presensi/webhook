package pdmk

import (
	"fmt"

	"github.com/whatsauth/itmodel"
)

// Pdmk generates a reply string based on the provided IteungMessage.
func Pdmk(Pesan itmodel.IteungMessage) (reply string) {
	return fmt.Sprintf("Hai.. hai..: %s\n\nBerikut adalah persyaratan pdm jenis keluar:\n\n"+
		"1. Surat Pengantar sesuai dengan format surat permohonan\n"+
		"2. Ijazah dan Transkrip Nilai\n"+
		"3. SK kelulusan / SK Dikeluarkan / SK Mutasi / SK Pengunduran diri (Perubahan status mahasiswa aktif menjadi lulus/dikeluarkan/mutasi/pengundyran diri)\n"+
		"4. Pakta Integritas (Perubahan status dari keluar/mutasi/pengunduran diri ke lulus)",
		Pesan.Alias_name)
}
