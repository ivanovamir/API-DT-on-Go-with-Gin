# API Deshevle TUT
___

![Logo](https://github.com/ivanovamir/API-DT-on-Go-with-Gin/blob/main/python_backend/media/admin-interface/logo/logo.png)


**Deshevle TUT Api** - this is a API that we developed with our team for a commercial site selling auto parts

## Technologies used in this project:
- Python (3.10.3)
  - Pillow (9.1.1)
  - gunicorn (20.1.0)
  - django-cors-headers (3.13.0)
  - psycopg2-binary (2.9.3)
  - django-admin-interface (0.20.0)
  - attr
  - django-object-actions (4.0.0)
  - requests (2.28.1)
  - crispy-bootstrap5 (0.3.1)
  - Django (4.0.5)
- Go (1.19.1)
  - Gin (1.18.1)
  - joho / godotenv
  - golang-jwt / jwt
  - mitchellh / mapstructure
  - mileusna / useragent
  - go-gorm / gorm
  - and other ...
- JavaScript
- HTML
- CSS
- Nginx
- PostgreSQL (14)

___

## Project objectives:
+ Create a fast and highly loaded **`API`** for an online store selling car parts
+ Develop a pleasant design and competent **`UX / UI`**
+ Convenient admin panel system
+ Correct ranking by search engines using **`SSR`**
+ Correct database architecture
+ Extensibility for future project updates

___

## Installation

**For security reasons, all files that can be compromised have been moved to **`.gitignore`** and used through the virtual environment. Therefore, you need to enter your own data and change the code a little.**

**At first you need to cteate database and register user in `PostgreSQL`:**
```sql
CREATE DATABASE DATABAS_NAME;
CREATE USER user_name WITH PASSWORD 'PASSWORD';
```
**Then you need connect `go` and `python` to this db:**

**In golang_backend/config/db.go:**
```go
db, err := gorm.Open(postgres.Open("postgres://user_name:password@host/port")
```
**In python_backend/main_settings/settings.py:**
```python
DATABASES = {
   'default': {
       'ENGINE': 'django.db.backends.postgresql',
       'NAME': 'db_name',
       'USER': 'user_name',
       'PASSWORD': 'password',
       'HOST': 'host',
       'PORT': 'port',
   }
}
```

**To run the Go api go to `cd golang_backend` and run this:**
```go
go mod build
go run cmd/main.go
```
**or**
```go
go mod build
go guild cmd/main.go
```

**To run the Admin panel go to `cd python_backend` and run this commands:**
```python
pip install -r req.txt

python manage.py makemigrations

python manage.py migrate

python manage.py runserver
```
___

## To contact with me:

### [Site link](https://deshevle-tut.ru/)
### [My Telegram](https://t.me/amirich18)


