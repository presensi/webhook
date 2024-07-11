package keyword

import (
	"fmt"

	"github.com/whatsauth/itmodel"
)

func Keyword(Pesan itmodel.IteungMessage) (reply string) {
	return fmt.Sprintf("Hai.. hai.. %s\n\nBerikut adalah Keyword yang tersedia:\n\n"+
		"1. *minta persyaratan pdm\n"+
		"2. *minta persyaratan pdmk\n"+
		"3. *minta persyaratan validasi ijazah\n"+
		"4. *minta persyaratan nidn\n"+
		"5. *minta dokumen panduan\n",
		Pesan.Alias_name)
}
