## Session

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
