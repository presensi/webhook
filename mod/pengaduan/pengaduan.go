package pengaduan

import (
	"fmt"

	"github.com/whatsauth/itmodel"
)

// Pdmk generates a reply string based on the provided IteungMessage.
func Pengaduan(Pesan itmodel.IteungMessage) (reply string) {
	return fmt.Sprintf("Hai.. hai.. %s\n\nBerikut adalah persyaratan pengaduan layanan publik:\n\n" +
		"1. *Surat pengantar dari pemohon\n" +
		"2. *Identitas pemohon\n" +
		"3. *Uraian pelayanan yang tidak sesuai dengan standar pelayanan dan uraian kerugian material atau imaterial yang diderita\n" +
		Pesan.Alias_name)
}
