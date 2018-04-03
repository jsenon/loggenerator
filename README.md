# loggenerator
 
This is a Log Generator Server with Logrus and Zapper samples

## Build

Build version
```
go build main.go
```

Check name of your binairies against name used in dockerfile


### Create Container

```
docker build -t jsenon/loggenerate:vX.X .
```

### Launch Container

```
docker run -d -p 80:9030 jsenon/loggenerator:vX.X
```

### Func

- Log with Zap Suggar 
- Log with Zap Logger
- Log with Logrus

### Usage

Logrus
```
curl -X POST 127.0.0.1:9030/log/logrus
```

ZapSuggar
```
curl -X POST 127.0.0.1:9030/log/zapsugar
```

ZapLogger
```
curl -X POST 127.0.0.1:9030/log/zaplogger
```

### TIPS

If you receive `cannot execute binary file: Exec format error`

Compile with  `CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go`

And build `docker build -t jsenon/loggenerator:vx.x .`

