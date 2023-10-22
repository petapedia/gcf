# Google Cloud Function Example
Contoh Penerapan Google Cloud Function dengan penggunaan Package [peda](https://pkg.go.dev/github.com/petapedia/peda).
1. [GET](./get)
2. [POST](./post)

## Langkah 
Buat database di mongodb.com. Contoh tampilan struktur collection mongo  
![image](https://github.com/petapedia/gcf/assets/11188109/1d100401-afc8-4451-81aa-be35f2d13ea1)

Jangan lupa setting env database mongo di google cloud function  
![image](https://github.com/petapedia/gcf/assets/11188109/a927c980-e81f-471a-a100-f437e330b185)

Isi file function.go dan go.mod dengan yang ada di repo gcf petapedia, klik deploy.  
![image](https://github.com/petapedia/gcf/assets/11188109/84f1be81-08e1-4d4c-9004-e3c905159b78)

Tunggu Proses Build Selesai  
![image](https://github.com/petapedia/gcf/assets/11188109/0f3ccfe9-7ec8-4cff-a7f1-e0e8a1375951)

Liat Log Jika Ada Error  
![image](https://github.com/petapedia/gcf/assets/11188109/80ead846-b81a-4e45-a0bc-d527f822948b)

Agar API bisa diakses Publik Berikan Akses Publik dari Menu Cloud Run dan Tambahkan permission sebagai berikut:  
New Principals : AllUsers  
Role : Cloud Run Invoker  
![image](https://github.com/petapedia/gcf/assets/11188109/45b84091-3e42-4124-9264-ae86e08d49b3)
