package mod

import (
	"github.com/gocroot/mod/dokumen"
	"github.com/gocroot/mod/idgrup"
	"github.com/gocroot/mod/idname"
	"github.com/gocroot/mod/ijazah"
	"github.com/gocroot/mod/lldikti"
	"github.com/gocroot/mod/mutasi"
	"github.com/gocroot/mod/nidn"
	"github.com/gocroot/mod/panduan"
	"github.com/gocroot/mod/pdm"
	"github.com/gocroot/mod/pdmk"
	"github.com/gocroot/mod/pendirian"
	"github.com/gocroot/mod/pengaduan"
	"github.com/whatsauth/itmodel"
	"go.mongodb.org/mongo-driver/mongo"
)

func Caller(Modulename string, Pesan itmodel.IteungMessage, db *mongo.Database, Profile itmodel.Profile) (reply string) {
	switch Modulename {
	case "idgrup":
		reply = idgrup.IDGroup(Pesan)
	case "idname-masuk":
		reply = idname.IDNameMasuk(Pesan, db)
	case "idname-pulang":
		reply = idname.IDNamePulang(Pesan, db)
	// case "selfie-masuk":
	// 	reply = idname.CekSelfieMasuk(Profile, Pesan, db)
	// case "selfie-pulang":
	// 	reply = idname.CekSelfiePulang(Pesan, db)
	case "lldikti":
		reply = lldikti.Lldikti(Pesan)
	case "panduan":
		reply = panduan.Panduan(Pesan)
	case "pdm":
		reply = pdm.Pdm(Pesan)
	case "pdmk":
		reply = pdmk.Pdmk(Pesan)
	case "nidn":
		reply = nidn.Nidn(Pesan)
	case "ijazah":
		reply = ijazah.Ijazah(Pesan)
	case "pengaduan":
		reply = pengaduan.Pengaduan(Pesan)
	case "mutasi":
		reply = mutasi.Mutasi(Pesan)
	case "pendirian":
		reply = pendirian.Pendirian(Pesan)
	case "dokumen":
		reply = dokumen.PanduanPDDIKTI(Pesan)
	}
	return
}
