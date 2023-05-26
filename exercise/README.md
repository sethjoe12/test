



A brief description of what this project does and who it's for

# Golang Rest api with Mysql and authentication with Prometheus and Grafana monitoring.



# Local Mysql Database Setup
$ sudo docker-compose up -d

$ docker exec -it my-mysql /bin/bash

# Enter Root Password sethjoe
$ maysql -uroot -p -A

$ update mysql.user set host='%' where user='root';

$ flush privileges;







## Register User



```sh
curl --location --request POST 'http://localhost:8000/api/v1/user/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "robie3",
    "username": "distroyer3",
    "email": "ambot3",
    "password": "ambot3"
}'
```

## Generate Token



```sh
curl --location --request POST 'http://localhost:8000/api/v1/token' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "ambot3",
    "password": "ambot3"
}'
```


## Get all Users



```sh
curl --location --request GET 'http://localhost:8000/api/v2/allusers' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImRpc3Ryb3llcjMiLCJlbWFpbCI6ImFtYm90MyIsImV4cCI6MTY4NDMxMDA1Mn0.BT0LNHN_aZX8ki2xonTETcuewuygCBGQolFEWejuLvg'
```

## Get User By Id



```sh
curl --location --request GET 'http://localhost:8000/api/v2/user/2' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImRpc3Ryb3llcjMiLCJlbWFpbCI6ImFtYm90MyIsImV4cCI6MTY4NDMxMDA1Mn0.BT0LNHN_aZX8ki2xonTETcuewuygCBGQolFEWejuLvg'
```

## Update User By Id



```sh
curl --location --request PUT 'http://localhost:8000/api/v2/user/6' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImRpc3Ryb3llcjMiLCJlbWFpbCI6ImFtYm90MyIsImV4cCI6MTY4NDMxMDA1Mn0.BT0LNHN_aZX8ki2xonTETcuewuygCBGQolFEWejuLvg'
```

## Dlete User By Id



```sh
curl --location --request DELETE 'http://localhost:8000/api/v2/user/2' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImRpc3Ryb3llcjMiLCJlbWFpbCI6ImFtYm90MyIsImV4cCI6MTY4NDMxMDA1Mn0.BT0LNHN_aZX8ki2xonTETcuewuygCBGQolFEWejuLvg'

```

Mysql Connection string

```sh

"root:Sethjoe12345@tcp(localhost:3306)/test_demo?parseTime=true"
