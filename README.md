# FinalProject2 Aplikasi MyGram Kelompok 4
Aplikasi ini akan dilengkapi dengan proses CRUD dengan table Final Project 2 Golang Hacktiv8

Pembagian tugas pada aplikasi MyGram :
* Aless : Mengerjakan tabel USER
* Alfin : Mengerjakan tabel PHOTO
* Fitri : Mengerjakan tabel COMMENT
* Faiz  : Mengerjakan tabel SOCIALMEDIA

Aplikasi ini telah di deploy ke cloud application platform Heroku.com

### Base Local URL   : `http://localhost:8080/user/register`
### Base Heroku URL  : `https://fp2-kelompok4-hactiv8.herokuapp.com`

### End Points on POSTMAN or INSOMNIA:
**USER**
* POST :
    1. Register
       * Untuk melakukan Registrasi User dapat menggunakan URL : `http://localhost:8080/user/register` atau `https://fp2-kelompok4-hactiv8.herokuapp.com/user/register` dengan method **POST**
       * Lalu menggunakan JSON berikut untuk registrasi :
            ```
            {
            "age" : 21,
            "email" : "dummy@gm.com",
            "username" : "dummy",
            "password" : "1234567"
            }   
            ```
       * Output response yang dihasilkan adalah :
           ```
           {
            "age": 21,
            "email": "dummy@gm.com",
            "id": 10,
            "username": "dummy"
            }
           ```
    2. Login
        * Untuk melakukan Login User untuk mendapatkan token dapat menggunakan URL : `http://localhost:8080/user/login` atau `https://fp2-kelompok4-hactiv8.herokuapp.com/user/login` dengan method **POST**
        *  Lalu menggunakan JSON berikut untuk Login :
            ```
            {
            "email" : "dummy@gm.com",
            "password" : "1234567"
            }
            ```
       * Output response yang dihasilkan adalah :
           ```
           {
           "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImR1bW15QGdtLmNvbSIsImlkIjo5fQ.UYVr69zLMoF_V1glErJudNic_-hni4n5uO7Ba9-L1YE"
           }
           ```
* PUT
    * Untuk melakukan Update data User dapat menggunakan URL : `http://localhost:8080/user/` atau `https://fp2-kelompok4-hactiv8.herokuapp.com/user/` dengan method **PUT**
    * Jangan lupa menggunakan token yang didapatkan dari endpoint `/user/login` untuk autentikasi dan authorisasi.
    * Lalu menggunakan JSON berikut untuk Update user :
         ```
        {
        "email" : "alesss@gm.com",
        "username" : "AlesssandroT"
        }
        ```      
    * Output response yang dihasilkan adalah :
         ```
        {
        "age": 21,
        "email": "alesss@gm.com",
         "id": 9,
        "updated_at": "2022-11-11T20:18:09.314+07:00",
         "username": "AlesssandroT"
        }
        ```
* DELETE
    * Untuk melakukan Delete data User dapat menggunakan URL : `http://localhost:8080/user/` atau `https://fp2-kelompok4-hactiv8.herokuapp.com/user/` dengan method **DELETE**
    * Jangan lupa menggunakan token yang didapatkan dari endpoint `/user/login` untuk autentikasi dan authorisasi.
    * Output response yang dihasilkan adalah :
        ```
        {
        "message": "your account has been successfully deleted"
        }
        ```


**PHOTO**
* GET :
   * Untuk melakukan Get atau menampilkan Photo dapat menggunakan URL : `http://localhost:8080/photos` atau `https://fp2-kelompok4-hactiv8.herokuapp.com/photos` dengan method **GET**
    * Jangan lupa menggunakan token yang didapatkan dari endpoint `/user/login` untuk autentikasi dan authorisasi.   
    * Output response yang dihasilkan adalah :
    ```
    {
    "photo": [
        {
            "id": 14,
            "title": "hallooaaaaaasad edit",
            "caption": "no",
            "photo_url": "www.rahasia.comaaaa2",
            "created_at": "2022-11-10T15:48:15Z",
            "user_id": 4,
            "updated_at": "2022-11-11T02:58:17Z",
            "user": {
                "email": "test1@gm.com",
                "username": "alfin"
            }
        },
        {
            "id": 16,
            "title": "hallooaaaaa",
            "caption": "no",
            "photo_url": "www.rahasia.comaaaa2",
            "created_at": "2022-11-10T16:10:54Z",
            "user_id": 4,
            "updated_at": "2022-11-10T16:10:54Z",
            "user": {
                "email": "test1@gm.com",
                "username": "alfin"
            }
        }
    ]
}
    ```
 * POST :
    * Untuk melakukan Post Comment dapat menggunakan URL : `http://localhost:8080/photos` atau `https://fp2-kelompok4-hactiv8.herokuapp.com/photos` dengan method **POST**
    * Jangan lupa menggunakan token yang didapatkan dari endpoint `/user/login` untuk autentikasi dan authorisasi.
    * Lalu menggunakan JSON berikut untuk Buat Photos:   
	```
	{
    	    "title": "hallooaaaaaasad edit",
            "caption": "no",
            "photo_url": "www.rahasia.comaaaa2"
	}
	```
    * Output response yang dihasilkan adalah :
    ```
    {
    "id": 20,
    "title": "hallooaaaaaasad edit",
    "caption": "no",
    "photo_url": "www.rahasia.comaaaa2",
    "created_at": "2022-11-12T06:32:51.091Z",
    "user_id": 4
    }
    ```
* PUT
	* Untuk melakukan Update atau mengedit Comment dapat menggunakan URL : `http://localhost:8080/photos/id` atau `https://fp2-kelompok4-hactiv8.herokuapp.com/photos/id` dengan method **PUT** misalnya ingin mengedit comment dengan id 8, maka id di urlnya diganti menjadi 20 : `http://localhost:8080/photos/20` atau `https://fp2-kelompok4-hactiv8.herokuapp.com/photos/20`
	    * Jangan lupa menggunakan token yang didapatkan dari endpoint `/user/login` untuk autentikasi dan authorisasi.
	    * Lalu menggunakan JSON berikut untuk Update comment :
	    ```
	    {
		    "title": "hallooaaaaaasad edit",
		    "caption": "no",
		    "photo_url": "edit"
	    }
	    ```
	   * Output response yang dihasilkan adalah :
	   ```
	   {
	    "id": 20,
	    "title": "hallooaaaaaasad edit",
	    "caption": "no",
	    "photo_url": "www.rahasia.comaaaa2",
	    "created_at": "2022-11-12T06:32:51.091Z",
	    "user_id": 4
	   }
	   ```
	   * Jika mengupdate id yang dimiliki oleh userid lain
	   ```
	   {
	    "message": "you are not allowed to update this id"
	   }
	   ```
* Delete
* Untuk melakukan Delete Comment dapat menggunakan URL : `http://localhost:8080/photos/id` atau `https://fp2-kelompok4-hactiv8.herokuapp.com/photos/id` dengan method **DELETE**. Misalnya ingin menghapus comment dengan id 8, maka id di urlnya diganti menjadi 20 : `http://localhost:8080/photos/20` atau `https://fp2-kelompok4-hactiv8.herokuapp.com/photos/20`
    * Jangan lupa menggunakan token yang didapatkan dari endpoint `/user/login` untuk autentikasi dan authorisasi.
    * Output response yang dihasilkan adalah :
    ```
    {
    "message": "Your photo has been successfully deleted"
    }
    ```
    * Jika menghapus id yang dimiliki oleh userid lain
    ```
    {
    "message": "you are not allowed to delete this id"
    }
    ```
**COMMENT**
* GET :
    * Untuk melakukan Get atau menampilkan Comment dapat menggunakan URL : `http://localhost:8080/comment/` atau `https://fp2-kelompok4-hactiv8.herokuapp.com/comment/` dengan method **GET**
    * Jangan lupa menggunakan token yang didapatkan dari endpoint `/user/login` untuk autentikasi dan authorisasi.   
    * Output response yang dihasilkan adalah :
        ```
        {
	        "comment": [
		        {
			        "id": 2,
			        "message": "keren",
			    "caption": 10,
			    "user_id": 1,
			    "updated_at": "2022-11-11T07:35:08Z",
			    "created_at": "2022-11-11T07:35:08Z",
			    "user": {
				    "id": 1,
				    "email": "aless@gm.com",
				    "username": "AlessandroT"
			            },
			    "photo": {
				    "id": 10,
				    "title": "test1@gm.com",
				    "caption": "aa",
				    "photo_url": "test",
				    "user_id": 1
			            }
		}
        ```
* POST :
    * Untuk melakukan Post Comment dapat menggunakan URL : `http://localhost:8080/comment/` atau `https://fp2-kelompok4-hactiv8.herokuapp.com/comment/` dengan method **POST**
    * Jangan lupa menggunakan token yang didapatkan dari endpoint `/user/login` untuk autentikasi dan authorisasi.
    * Lalu menggunakan JSON berikut untuk Update user :
         ```
        {
		"message" : "hi",
		"photo_id" : 10
	    }
        ```      
    * Output response yang dihasilkan adalah :
        ```
        {
	    "created_at": "2022-11-12T06:52:15.251+07:00",
	    "id": 8,
	    "message": "hi",
	    "photo_id": 10,
	    "user_id": 1
        }
        ```
* PUT
    * Untuk melakukan Update atau mengedit Comment dapat menggunakan URL : `http://localhost:8080/comment/id` atau `https://fp2-kelompok4-hactiv8.herokuapp.com/comment/id` dengan method **PUT** misalnya ingin mengedit comment dengan id 8, maka id di urlnya diganti menjadi 8 : `http://localhost:8080/comment/8` atau `https://fp2-kelompok4-hactiv8.herokuapp.com/comment/8`
    * Jangan lupa menggunakan token yang didapatkan dari endpoint `/user/login` untuk autentikasi dan authorisasi.
    * Lalu menggunakan JSON berikut untuk Update comment :
         ```
        {
		"message" : "hi, apa kabar?"
	    }
        ```      
    * Output response yang dihasilkan adalah :
         ```
        {
	    "id": 8,
	    "message": "hi, apa kabar?",
	    "photo_id": 10,
	    "updated_at": "2022-11-12T06:59:01.934+07:00",
	    "user_id": 1
        }
        ```
* DELETE
    * Untuk melakukan Delete Comment dapat menggunakan URL : `http://localhost:8080/comment/id` atau `https://fp2-kelompok4-hactiv8.herokuapp.com/comment/id` dengan method **DELETE**. Misalnya ingin menghapus comment dengan id 8, maka id di urlnya diganti menjadi 8 : `http://localhost:8080/comment/8` atau `https://fp2-kelompok4-hactiv8.herokuapp.com/comment/8`
    * Jangan lupa menggunakan token yang didapatkan dari endpoint `/user/login` untuk autentikasi dan authorisasi.
    * Output response yang dihasilkan adalah :
        ```
        {
	    "message": "your comment has been successfully deleted"
        }
        ```

**SOCIAL MEDIA**
* POST :
    * Untuk membuat data social media baru dapat dengan menggunakan url :
    `http://localhost:8080/socialmedias` atau `https://fp2-kelompok4-hactiv8.herokuapp.com/socialmedias`
    * Kemudian gunakan json berikut untuk membuat datanya:
    ```json
    {
	"name":"Instagram",
	"social_media_url":"https://instagram.com/faiz"
    }
    ```
    * Untuk akses endpointnya dibutuhkan request autorisasi token yang didapatkan dari response endpoint user/login
    * Output response yang dihasilkan adalah :
    ```json
    {
	"id": 5,
	"name": "Instagram",
	"social_media_url": "https://instagram.com/faiz",
	"user_id": 5,
	"created_at": "2022-11-10T21:16:25.68+07:00"
    }
    ```
* GET :
    * Untuk menampilkan semua data social media dapat dengan menggunakan url :
    `http://localhost:8080/socialmedias` atau `https://fp2-kelompok4-hactiv8.herokuapp.com/socialmedias`
    * Untuk akses endpointnya dibutuhkan request autorisasi token yang didapatkan dari response endpoint user/login
    * Output response yang dihasilkan adalah :
    ```json
    {
	"social_medias": [
	    {
			"id": 5,
			"name": "Instagram",
			"social_media_url": "https://instagram.com/faiz",
			"UserId": 5,
			"createdAt": "2022-11-10T12:52:05Z",
			"updatedAt": "2022-11-10T12:52:05Z",
			"User": {
				"id": 5,
				"username": "faiz09"
			}
        },
		{
			"id": 6,
			"name": "Facebook",
			"social_media_url": "https://facebook.com/faizz",
			"UserId": 5,
			"createdAt": "2022-11-10T14:16:25Z",
			"updatedAt": "2022-11-10T15:00:10Z",
			"User": {
				"id": 5,
				"username": "faiz09"
			}
		}
	  ]   
    }
    ```
* PUT :
    * Untuk mengedit data social media dengan id 6 dapat dengan menggunakan url :
    `http://localhost:8080/socialmedias/6` atau `https://fp2-kelompok4-hactiv8.herokuapp.com/socialmedias/6`
    * Kemudian gunakan json berikut untuk mengedit datanya:
    ```json
    {
	"name": "Whatsapp",
	"social_media_url": "https://wa.me/12345"
    }
    ```
    * Untuk akses endpointnya dibutuhkan request autorisasi token yang didapatkan dari response endpoint user/login
    * Output response yang dihasilkan adalah :
    ```json
    {
	"id": 6,
	"name": "Whatsapp",
	"social_media_url": "https://wa.me/12345",
	"user_id": 5,
	"created_at": "2022-11-10T14:16:25Z"
    }
    ```
* DELETE :
    * Untuk menghapus data social media dengan id 6 dapat dengan menggunakan url :
    `http://localhost:8080/socialmedias/6` atau `https://fp2-kelompok4-hactiv8.herokuapp.com/socialmedias/6`
    * Untuk akses endpointnya dibutuhkan request autorisasi token yang didapatkan dari response endpoint user/login
    * Output response yang dihasilkan adalah :
    ```json
    {
	"message": "Your social media has been successfully deleted"
    }
    ```
