# Mini Task Go 

## Membuat Program Interaktif Hitung Luas dan Keliling Lingkaran

## Implementasi nya
1. Buatkan 2 fungsi untuk menghitung luas dan keliling lingkaran dengan isi parameter jari jari yaitu r dengan tipe data integer dan point dari kedua ini yaitu menghasilkan nilai float32

```go
func hitungLuas(r int) float32{

}
func kelLingkaran(r int) float32{
    
}
```

2. Menentukan nilai phi berdasarkan kondisi r, jika r bisa di modulus 7 sama dengan 0 (hasil bagi dari 7 sama dengan 0) maka nilai phi 22/7 jika tidak phi 3.14
CATATAN : karena variabel phi di deklarasikan dengan tipe float32 maka phi 22/7 menjadi 22.0/7.0

```go
var phi float32
if r%7 == 0 {
    if r%7 == 0 {
		phi = 22.0 / 7.0
	} else{
		phi = 3.14
	}
}

```
3. Masukan kode kondisi tersebut ke dalam 2 fungsi tadi, hitungLuas dan kelLingkarang
4. Untuk fungsi hitungLuas, point nya menghasilkan nilai dari rumus luas lingkaran dengan kode berikut
```go
return float32(r) * float32(r) * phi
```
5. Untuk fungsi kelLingkaran, point nya menghasilkan nilai dari rumus keliling lingkaran dengan kode berikut
```go
return  2 * phi * float32(r)
```
6. Masukan/panggil fungsi hitung dan luas lingkaran dalam fungsi main dengan kode yang dibuatkan secara interaktif seperti berikut
```go
func main() {
	var r int
	fmt.Print("Masukan jari-jari lingkaran : ")
	fmt.Scan(&r)

	keliling := kelLingkaran(r)
	luas := hitungLuas(r)

	fmt.Println("======= Hasil =======")
	fmt.Println("Jari - Jari : ", r)
	fmt.Println("Luas Lingkarang : ", luas)
	fmt.Println("Keliling Lingkaran : ", keliling)
}
```
