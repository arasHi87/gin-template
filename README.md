# Gin template

A RESTful API service built with gin.

## Developement

### Prerequisite

| Name     | Version |
| -------- | ------- |
| GNU make | 4.2.1   |
| go       | 1.17    |

### Environment setup

0. Initialize environment variable

```
cp env-sample .env
```

1. Download required packages

```
make init
```

2. Start development API service

```
make run
```

3. (Optional) Create binary executable output

```
make build
```

### Formatting

This project use `gofmt` for formatting

```
make format
```
