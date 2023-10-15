# Google Cloud Function HTTP POST

Dicontohkan dengan menggunakan databse mongo  
![image](https://github.com/petapedia/gcf/assets/11188109/46863a76-b87e-436d-9598-93253e7df8e2)

Test API  
![image](https://github.com/petapedia/gcf/assets/11188109/b36ffe44-f71e-4322-83e6-ee2803a19381)

Setting Environment Variabel  
![image](https://github.com/petapedia/gcf/assets/11188109/ac2c505e-34d4-4a2a-862a-a6926bd14f4a)

if u want to get token in the header of post just use this function 
```go
func GetToken(r *http.Request) string {
    return r.Header.Get("Authorization")
}
```
