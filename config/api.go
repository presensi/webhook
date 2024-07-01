package config

import (
	"github.com/gocroot/helper"
)

var WAAPIQRLogin string = "https://api.wa.my.id/api/whatsauth/request"

var WAAPIMessage string = "https://api.wa.my.id/api/send/message/text"

var WAAPIGetToken string = "https://api.wa.my.id/api/signup"

var PublicKeyWhatsAuth string

var WAAPIToken string

var DOKUMENPANDUAN, ErrDokumen = helper.DownloadFileBase64("https://github.com/4c3d2765-882f-4557-bc45-eed332449888")
