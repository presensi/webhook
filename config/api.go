package config

import (
	"github.com/gocroot/helper/file"
)

var WAAPIQRLogin string = "https://api.wa.my.id/api/whatsauth/request"

var WAAPIMessage string = "https://api.wa.my.id/api/send/message/text"

var WAAPIGetToken string = "https://api.wa.my.id/api/signup"

var PublicKeyWhatsAuth string

var WAAPIToken string

var DOKUMENPANDUAN, ErrDokumenPanduan = file.DownloadFileBase64("https://raw.githubusercontent.com/haryadi14/jurnal_peningkatan-/master/panduan_compressed.pdf")

var PDMK, ErrDokumenPDMK = file.DownloadFileBase64("https://raw.githubusercontent.com/haryadi14/jurnal_peningkatan-/master/Prosedur_pdmk.pdf")
