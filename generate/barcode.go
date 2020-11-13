package generate

import (
	"github.com/massarakhsh/lik"
	"fmt"
	"github.com/boombuler/barcode/ean"
	"image/png"
	"math/rand"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func GenerateQR(code string) string {
	// Create the barcode
	qrCode, _ := qr.Encode(code, qr.M, qr.Auto)

	// Scale the barcode to 200x200 pixels
	qrCode, _ = barcode.Scale(qrCode, 200, 200)

	// create the output file
	rnd := rand.Intn(1000000000)
	path := fmt.Sprintf("var/code/%09d.png", rnd)
	if match := lik.RegExParse(path, "(.+)/[^/]*$"); match != nil {
		os.MkdirAll(match[1], os.ModePerm)
	}
	file, _ := os.Create(path)
	defer file.Close()

	// encode the barcode as png
	png.Encode(file, qrCode)
	return path
}

func GenerateEan13(code string) string {
	path := "/ean13/" + code + ".png"
	return DirectEan13(code, path)
}

func DirectEan13(code string, path string) string {
	if code = complement_ean13(code); code != "" {
		eanCode, _ := ean.Encode(code)
		barCode, _ := barcode.Scale(eanCode, 1280, 320)
		eanCode = barCode.(barcode.BarcodeIntCS)
		loc := path
		if match := lik.RegExParse(path, "^/(.+)"); match != nil {
			loc = match[1]
		}
		if match := lik.RegExParse(loc, "(.+)/[^/]*$"); match != nil {
			os.MkdirAll(match[1], os.ModePerm)
		}
		file, _ := os.Create(loc)
		defer file.Close()
		png.Encode(file, eanCode)
	} else {
		path = "/images/noean13.png"
	}
	return path
}

func complement_ean13(code string) string {
	if lc := len(code); lc == 12  || lc == 13 {
		sum := 0
		for nd := 0; nd < 12; nd += 2 {
			if dg := code[nd]; dg >= '0' && dg <= '9' {
				sum += (int(dg) - '0') * 1
			} else {
				sum = -1
				break
			}
			if dg := code[nd + 1]; dg >= '0' && dg <= '9' {
				sum += (int(dg) - '0') * 3
			} else {
				sum = -1
				break
			}
		}
		if sum < 0 {
			code = ""
		} else {
			crc := '9' - uint8((sum + 9) % 10)
			if lc == 13 && crc != code[12] {
				code = ""
			} else {
				code = code[:12] + string(crc)
			}
		}
	}
	if len(code) != 13 {
		code = ""
	}
	return code
}