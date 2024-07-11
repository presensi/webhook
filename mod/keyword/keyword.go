package keyword

import (
	"fmt"

	"github.com/whatsauth/itmodel"
)

func Keyword(Pesan itmodel.IteungMessage) (reply string) {
	return fmt.Sprintf("Hai.. hai.. %s\n\nBerikut adalah Keyword yang tersedia:\n\n"+
		"1. *minta persyaratan pdm (Adalah keywords untuk meminta persyaratan perubahan data mahasiswa)\n"+
		"2. *minta persyaratan pdmk (Adalah keywords untuk meminta persyaratan perubahan data mahasiswa jenis keluar)\n"+
		"3. *minta persyaratan validasi ijazah (Adalah keywords untuk meminta persyaratan untuk memvalidasi ijazah mahasiswa)\n"+
		"4. *minta persyaratan nidn (Adalah keywords untuk meminta persyaratan Regristasi Dosen Baru NIDN)\n"+
		"5. *minta dokumen panduan (Adalah keywords untuk meminta dokumen panduan)\n",
		Pesan.Alias_name)
}
