package main

import (
	"fmt"
	"os"
	"strconv"
)

type Data struct {
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func main() {
	varkey, _ := strconv.Atoi(os.Args[1])
	getdata(varkey)
}

func getdata(key int) {
	var friend = []*Data{
		{nama: "Frisky", alamat: "jl durian", pekerjaan: "wiraswasta", alasan: "keren"},
		{nama: "Yoseph Bram", alamat: "jl pepaya", pekerjaan: "PNS", alasan: "iseng"},
		{nama: "Rizky", alamat: "jl Apel", pekerjaan: "mahasiswa", alasan: "biar pinter"},
		{nama: "Tyas", alamat: "jl Nanas", pekerjaan: "Pedagang", alasan: "gabut"},
		{nama: "Diko", alamat: "jl Mangga", pekerjaan: "Petani", alasan: "biar jago"},
	}

	var nama = func(name []*Data, key int) {
		for k, x := range name {
			if k == key {
				fmt.Println("Nama: "+x.nama, "Alamat: "+x.alamat, "Pekerjaan: "+x.pekerjaan, "Alasan: "+x.alasan)
			}
		}
	}
	nama(friend, key)
}
