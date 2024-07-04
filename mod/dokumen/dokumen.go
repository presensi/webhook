package dokumen

import (
	"fmt"

	h "github.com/gocroot/config/helper"
	"github.com/gocroot/helper/atapi"
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
		Base64Doc: h.GetDokumen(),
		Filename:  "Panduan PDDIKTI Admin.pdf",
		Isgroup:   false,
		Caption:   "Ini Dokumennya yaa...",
	}
	profile, _ := h.GetAppProfile("62895800006000", h.GetMongo())
	_, _, err := atapi.PostStructWithToken[itmodel.Response]("Token", profile.Token, data, "https://api.wa.my.id/api/send/message/document")
	if err != nil {
		return fmt.Sprintf("Gagal mengirim dokumen: %v", err)
	}

	return
}
