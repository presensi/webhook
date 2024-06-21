package pendirian

import (
	"fmt"

	"github.com/whatsauth/itmodel"
)

func Pendirian(Pesan itmodel.IteungMessage) (reply string) {
	return fmt.Sprintf("Hai.. hai.. %s\n\nBerikut adalah persyaratan rekomendasi pendirian perguruan tinggi swasta:\n\n"+
		"1. *Surat pengantar dari Ketua Badan Penyelenggara\n"+
		"2. *Akta notaris pendirian Badan Penyelenggara beserta semua perubahannya (jika pernah dilakukan perubahan) berupa scan dari dokumen asli\n"+
		"3. *Surat keputusan pejabat yang berwenang tentang pengesahan Badan Penyelenggara sebagai badan hukum, (Surat Keputusan Menkumham) berupa scan dari dokumen\n"+
		"4. Kerja sama dengan dunia usaha atau industri untuk program pendidikan vokasi (Perjanjian Kerjasama)\n"+
		"5. *Ketersediaan lahan untuk kampus perguruan tinggi sesuai dengan persyaratan dibuktikan dengan sertifikat/akta sewa\n"+
		"6. *Rekap dosen pada setiap program studi yang diajukan (format pdf)\n",

		Pesan.Alias_name)
}
