Dockerized-Golang-Postgres-API-Transaction

Db = automigrate ,mounted on database_postgres volumes.

Stack docker-compose ,build di macbook. 
Namun bisa juga ke swarm mode untuk horizontal scalling jika dibutuhkan ,krna yang saya pakai sudah versi 3 higher.

Untuk menjalankan program ,silahkan menggunakan perintah

docker-compose up -d


1.Web app : localhost:8002/api/user/new 
Harus insert data dulu ke db untuk user ,agar mendapatkan token 

Method post 
Contoh insert data user dengan format ,json di body 
{
"email" : "developerpdak@gmail.com",
"name" : "Andre",
"password" : "S2komputer"
}

2.login for tokenizer [POST]
localhost:8002/api/user/login
{
"email" : "developerpdak@gmail.com",
"password" : "S2komputer"
}

bearer example : eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjE2fQ.Td6mWqCCy4dS9YApsEh43BSyvlUWmmWayU-ESEHrxYU

3.get data transaction [GET]
localhost:8002/api/transaction
attach bearer / bawa token ,klw tidak akan ada pesan perlu token ,dan data tidak akan di tampilkan


4.localhost:8002/api/transaction/new
insert data transaction [POST]
attach bearer / bawa token ,klw tidak akan ada pesan perlu token ,dan data tidak akan di simpan

{
"uuid" : "2",
"user_id" : 16,
"device_timestamp" : "2021-09-06T01:59:17",
"total_amount" : "890000",
"paid_amount" : "4550000",
"change_amount" : "66000",
"payment_method" : "cash"
}
5.update data transaction [PATCH]
localhost:8002/api/transaction/update?id=1
attach bearer / bawa token ,klw tidak akan ada pesan perlu token ,dan data tidak akan di update
{
"uuid" : "2",
"user_id" : 16,
"device_timestamp" : "2021-09-06T01:59:17",
"total_amount" : "10000",
"paid_amount" : "20000",
"change_amount" : "77000",
"payment_method" : "cash"
}
6.delete data transaction [DELETE]
localhost:8002/api/transaction/delete?id=1
attach bearer / bawa token ,klw tidak akan ada pesan perlu token ,dan data tidak akan di delete

7.detail data transaction [GET]
localhost:8002/api/transaction/detail?id=1
attach bearer / bawa token ,klw tidak akan ada pesan perlu token ,dan data tidak akan di tampilkan

Local collection = local_collection.postman_collection.json
Public collection = Public Collection.txt

Writed ,5 Sept' 2021
Developer ,Andre ,Jakarta-Indonesia
