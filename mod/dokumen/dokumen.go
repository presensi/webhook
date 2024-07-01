package dokumen

import (
	"fmt"

	"github.com/gocroot/helper"

	"github.com/gocroot/config"
	"github.com/whatsauth/itmodel"
)

type DocumentMessage struct {
	To        string  `json:"to"`
	Base64Doc *string `json:"base64doc"`
	Filename  string  `json:"filename"`
	Isgroup   bool    `json:"isgroup"`
	Caption   string  `json:"caption"`
}

func PanduanPDDIKTI(Pesan itmodel.IteungMessage) (reply string) {
	data := DocumentMessage{
		To:        Pesan.Phone_number,
		Base64Doc: config.GetDokumen(),
		Filename:  "Panduan PDDIKTI Admin.pdf",
		Isgroup:   false,
		Caption:   "Ini Dokumennya yaa...",
	}
	profile, _ := helper.GetAppProfile(Pesan.Phone_number, config.Mongoconn)
	_, err := helper.PostStructWithToken[itmodel.Response]("Token", profile.Token, data, "https://wa.my.id/send/message/document")
	if err != nil {
		return fmt.Sprintf("Gagal mengirim dokumen: %v", err)
	}

	return
}
