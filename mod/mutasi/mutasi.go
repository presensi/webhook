package mutasi

import (
	"fmt"

	"github.com/whatsauth/itmodel"
)

func Mutasi(Pesan itmodel.IteungMessage) (reply string) {
	return fmt.Sprintf("Hai.. hai.. %s\n\nBerikut adalah persyaratan mutasi mahasiswa:\n\n"+
		"1. *Surat Pengantar dari Pimpinan Penerima\n"+
		"2. *Surat Keterangan Mutasi Mahasiswa dari PT asal\n"+
		"3. *Terdaftar di PDDIKTI (tangkapan layar)\n"+
		"4. *Salinan Transkrip Nilai\n"+
		"5. *Salinan KTM\n",
		Pesan.Alias_name)
}
