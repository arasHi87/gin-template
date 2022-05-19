# Gin Template

A simple gin RESTful API template with following function.

- complete backend architecture
- use gorm build data model
- use logrus and other plugin to build logger service
- user swagger to generate docs and mock test
- simple user auth service with JWT
- simple test framework
- ci/cd example

## Arch

```
--- pkg
 |------- common
 |------- controller
 |------- middleware
 |------- model
 |------- router
 |------- setting
```

- common
  - put tool which use in other service, ex. jwt, logger, etc.
- controller
  - accept input come from router
- middleware
  - some middleware, ex. logger, auth, etc.
- router
  - put all router here
- setting
  - convert enviroment value to CONFIG struct here

## Enviroment

All env this project use, you can change default value in [here](https://github.com/arasHi87/gin-template/blob/4135c0937fb69b9d301eae120ae833893443a775/pkg/setting/setting.go#L51).

| Name        | Description                         | Value        | Type   |
| ----------- | ----------------------------------- | ------------ | ------ |
| APP_ADDRESS | web app address                     | 8080         | string |
| APP_PORT    | web app port                        | 0.0.0.0      | string |
| DB_HOST     | DB host                             | localhost    | string |
| DB_PORT     | DB port                             | 5432         | string |
| DB_NAME     | DB name                             | gintemplate  | string |
| DB_USERNAME | DB user name                        | arashi87     | string |
| DB_PASSWORD | DB user password                    | m3ow87       | string |
| DB_TIMEZONE | DB timezone                         | Asia/Taipei  | string |
| JWT_EXPIRE  | JWT expire time (hours)             | 1            | int    |
| JWT_SECRET  | JWT secret                          | arashi87     | string |
| JWT_ISSUER  | JET issuer                          | arashi87     | string |
| LOG_PATH    | Log file save path                  | /tmp/gin.log | string |
| LOG_EXPIRE  | Log file expire time (hours)        | 720          | int    |
| LOG_ROTATE  | When should log file rotate (hours) | 12           | int    |

## Logger

This project use logrus + rotatelogs + lfshook to build log service.

- logrus use to instead of origin go logger
- rotatelogs use to split log file periodically
- lfshook is the logrus hook plugin, it will match log level to decide which log file should be write.

## Test

Not implement yet.

## CI/CD

Not implement yet.

## Run

| Name | Version       |
| ---- | ------------- |
| Go   | go1.17.6      |
| Make | GNU Make 3.81 |

0. Initialize environment variable

```
cp env-sample .env
```

1. Start database

```
docker-compose up -d
```

2. Build service

It will build swagger document in ./docs at the same time, and u will get a executable file named app.

```
make build
```

3. Start service

```
./app
```

4. Result

You should be able to accept service by `http://localhost:8080/health` and accept swagger docs by `http://localhost:8080/swagger/index.html`.
