# Dcard Demo
### Simple Live Demo
![Demo](https://upload.cc/i1/2022/03/31/c5VgQY.png)
#### Front-end

I've deployed a simple shorten-URL Demo that have developed by React. If You wanna demo,visit the link below(a bit slow~):

```url
https://shortener-url.herokuapp.com/
```



#### Backend

**API Server have been deployed on AWS EC2 and the host name:**

```
https://ec2-18-141-10-193.ap-southeast-1.compute.amazonaws.com
```

**Using this API to check server active state. If it's being active,you'll receive a `pong` message**

```
https://ec2-18-141-10-193.ap-southeast-1.compute.amazonaws.com/ping
```

### Tool in use

* Gin-Gonic(backend framework)
* Mysql database 
* GROM
* Validator v10 
* Viper
* Swagger-API

#### APIs Document

Allowing to access Swagger-API if server is in **debug mode**

```yaml
Mode: debug
```

```
{HOST}/swagger/index.html
```

### APIs
|Uri|Method|Desc|
|---|---|---|
|/api/v1/urls|POST|Upload Longrest URL with expired date in UTC format|
|/{url_id}|GET|Redirect to original url with url_id|

### How it work
#### Upload url
* Getting uploaded data from the HTTP Form and validating  the data via Validator including URL and time
* Create a new tuple/record to the DB with Form data if data is validated successfully
* After tuple is created, server'll use **base62 Algorithm** to generate the short_URL depending on Unix time
  * base62 Algorithm Implementation
  * Input int64(unix time) and output a string which string consists of a-z, A-Z,and 0-9(total 62 characters) 
  ```go
  //Base62URL in order to generate shorten url string,it's consider to use base 62 approach(6byte)
  func Base62URL(id int64) string {
    if id < 0 {
      return ""
    }
    //[a-z,A-Z,0-9]
    urlStr := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" //total 62 character
    url := ""
  
    for {
      if id <= 0 {
        break
      }
      str := urlStr[int(id%62)]
      url += string(str)
      id = id / 62
    }
    return url
  }
  ```
* After the short_URL is generated by the algorithm, the short_URL will be updated in DB

#### Redirect to original URL
* Getting short URL id from URL path, for example: the short URL_id is `abc123`
* We'll query the database and check the URL_id (Checking steps below)
  * Step 1 :whether the URL_id exist
    * 1.the URL_id **not exist**,return *404 Not Found* Immediately
    * 2.the URL_id **exist** in the database,it goes to the next step
  * Step 2 :whether the url_id expire
    * 1.the URL_id expired, return *404 Not Found* Immediately
    * 2.the URL_id does not expire, redirect to the original_URL

