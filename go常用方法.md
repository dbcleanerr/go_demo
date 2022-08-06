### 交叉编码

#### `Windows` 下编译 `Mac` 和 `Linux 64`位可执行程序

```shell
# mac
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build main.go

# Linux 64
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build main.go
```

#### `Mac`下编译 `Windows` 和 `Linux 64`位可执行程序

```shell
# Linux 64
CGO_ENABLED=0 
GOOS=linux 
GOARCH=amd64 
go build main.go

# Windows
CGO_ENABLED=0 
GOOS=windows 
GOARCH=amd64 
go build main.go
```