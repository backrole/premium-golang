# Projek Golang
![Golang](https://user-images.githubusercontent.com/53844625/108027912-845eb200-705d-11eb-833d-fb82961c291e.png?raw=true)


## Login

/api/v1/sessions

POST
```
{
    "email" : "email@gmail.com",
    "password" : "bahrul99"
}
```
Response
```
{
    "meta": {
        "message": "Login SUKSES",
        "code": 422,
        "status": "Sukses"
    },
    "data": {
        "id": 14,
        "nama": "bahrul",
        "kampus": "ayam",
        "email": "email@gmail.com",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxNH0.gQbqJn2o4v9bhV7TXuCZR474U1qsi2x_P7cHtLSYCWc"
    }
}
```

## Register

/api/v1/users

POST
```
{
    "nama" : "jwt auth",
    "email" : "jwt@gmail.com",
    "kampus" : "disitu",
    "password" : "bahrul99"
}
```

Response
```
{
    "meta": {
        "message": "Akun terregistrasi",
        "code": 200,
        "status": "Succes"
    },
    "data": {
        "id": 15,
        "nama": "jwt auth",
        "kampus": "disitu",
        "email": "jwt@gmail.com",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxNX0.-iwQTOtWMb4CxPUhfNevyLhEW8Nbh0PvlBrcFZbX7Fw"
    }
}
```

## Cek Register

/api/v1/email_ceks

POST
```
{
    "email" : "email@gmail.com"
}
```

Response
```
{
    "meta": {
        "message": "Email has been Registered",
        "code": 200,
        "status": "sukses"
    },
    "data": {
        "is_available": false
    }
}
```

## Upload Avatar

/api/v1/avatars

POST
```
Key : avatar
Type : File

Value : gambar anda
```

Response
```
{
    "meta": {
        "message": "Sukses upload gambar",
        "code": 200,
        "status": "sukses"
    },
    "data": {
        "is_uploaded": true
    }
}
```
