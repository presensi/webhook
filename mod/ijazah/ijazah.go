package ijazah

import (
	"fmt"

	"github.com/whatsauth/itmodel"
)

// Pdmk generates a reply string based on the provided IteungMessage.
func Ijazah(Pesan itmodel.IteungMessage) (reply string) {
	return fmt.Sprintf("Hai.. hai.. %s\n\nBerikut adalah persyaratan validasi ijazah:\n\n"+
		"1. *Surat Permohonan dari Perguruan Tinggi/Yayasan terkait\n"+
		"2. *Salinan ijazah\n"+
		"3. Hasil pindai dokumen ujian negara (untuk yang lulus sebelum th 2001 dengan status PT yang terdaftar dan diakui)\n"+
		"4. Hasil cetak profil mahasiswa pada Pangkalan Data Pendidikan Tinggi (PDDikti)\n"+
		"5. *Surat Pernyataan (form terlampir)",
		Pesan.Alias_name)
}
