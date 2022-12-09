# 

## API
### 获得验证码
```shell
curl --location --request POST 'http://localhost:8777/api/getCaptcha' \
--header 'Content-Type: application/json' \
--data-raw '{
    "CaptchaType":"digit",
    "DriverDigit":{"Length":5,"Width":240,"Height":80}
}'
```

### 验证
```shell
curl --location --request POST 'http://localhost:8777/api/verifyCaptcha' \
--header 'Content-Type: application/json' \
--data-raw '{
    "Id":"jfSU4vJfYVQhRaDsZcMO",
    "VerifyValue":"49308"
}'
```