package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	// Kullanıcıdan RGB değeri al
	var input string
	fmt.Print("RGB değeri girin (örn: 255,255,255): ")
	fmt.Scanln(&input)

	parts := strings.Split(input, ",")
	if len(parts) != 3 {
		fmt.Println("Geçersiz format. Örn: 255,255,255 şeklinde girin.")
		return
	}

	r, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
	g, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))
	b, err3 := strconv.Atoi(strings.TrimSpace(parts[2]))
	if err1 != nil || err2 != nil || err3 != nil || r < 0 || r > 255 || g < 0 || g > 255 || b < 0 || b > 255 {
		fmt.Println("Geçersiz RGB değerleri.")
		return
	}

	// Çalışma dizinini al
	baseDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}

	// iconlar klasörünün yolunu oluştur
	iconlarDir := filepath.Join(baseDir, "iconlar")

	// iconlar klasörü var mı kontrol et
	if _, err := os.Stat(iconlarDir); os.IsNotExist(err) {
		fmt.Println("iconlar klasörü bulunamadı.")
		return
	}

	// iconlar klasöründeki dosyaları tara
	files, err := os.ReadDir(iconlarDir)
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(strings.ToLower(file.Name()), ".png") {
			filePath := filepath.Join(iconlarDir, file.Name())
			fmt.Println(file.Name(), "işleniyor...")

			f, err := os.Open(filePath)
			if err != nil {
				fmt.Println("Açma hatası:", err)
				continue
			}
			img, err := png.Decode(f)
			f.Close()
			if err != nil {
				fmt.Println("Decode hatası:", err)
				continue
			}

			bounds := img.Bounds()
			newImg := image.NewNRGBA(bounds)

			for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
				for x := bounds.Min.X; x < bounds.Max.X; x++ {
					_, _, _, a := img.At(x, y).RGBA()
					if a == 0 {
						newImg.Set(x, y, color.NRGBA{0, 0, 0, 0}) // transparan koru
					} else {
						newImg.Set(x, y, color.NRGBA{uint8(r), uint8(g), uint8(b), uint8(a >> 8)})
					}
				}
			}

			outFile, err := os.Create(filePath)
			if err != nil {
				fmt.Println("Yazma hatası:", err)
				continue
			}
			err = png.Encode(outFile, newImg)
			outFile.Close()
			if err != nil {
				fmt.Println("Encode hatası:", err)
				continue
			}
			fmt.Println(file.Name(), "tamamlandı.")
		}
	}

	fmt.Println("Tüm iconlar klasöründeki PNG'ler işlendi.")
	fmt.Println("Çıkmak için Enter'a basın...")
	fmt.Scanln()
}
