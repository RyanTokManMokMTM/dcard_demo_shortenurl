# Dcard Demo
Allowing to access local open apis if server is in **debug mode**
```yaml
Mode: debug
```
```
http://127.0.0.1:8000/swagger/index.html
```
### Tool in use
* Gin-Gonic(backend framwork)
* Mysql database with GORM
* Validator v10 
* Viper
* Swagger-API

### APIs
|Uri|Method|Desc|
|---|---|---|
|/api/v1/urls|POST|Upload Longrest URL with expired date in UTC format|
|/:url_id|GET|Redirect to original url with url_id|



### How it work
