package main

import "fmt"

func main() {

	// Önce değişkenleri hazırlıyoruz.

	var islem int
	var sayi1 int
	var sayi2 int

	//Giriş ekranı.

	fmt.Println("Yapmak istediğiniz işlemi tablodakine göre seçiniz:\n 1- Toplama\n 2- Çıkartma\n 3- Çarpma\n 4- Bölme")
	fmt.Scanln(&islem)
	fmt.Println("Sırasıyla işlem yapmak istedikleriniz sayıları giriniz: (Maksimum 2 tane)")
	fmt.Scanln(&sayi1)
	fmt.Scanln(&sayi2)

	// Ardından şartları switch kullanarak şartları belirtiyoruz.

	switch {
	case islem == 1:
		sonuc := sayi1 + sayi2
		fmt.Println("Sonuç: ", sonuc)

	case islem == 2:
		sonuc := sayi1 - sayi2
		fmt.Println("Sonuç: ", sonuc)

	case islem == 3:
		sonuc := sayi1 * sayi2
		fmt.Println("Sonuç: ", sonuc)

	case islem == 4:
		sonuc := sayi1 / sayi2
		fmt.Println("Sonuç: ", sonuc)

	default:
		fmt.Println("Geçersiz girdi.")
	}
}
