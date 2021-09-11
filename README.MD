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



Writed ,11 Sept' 2021
Developer ,Andre ,Jakarta-Indonesia
