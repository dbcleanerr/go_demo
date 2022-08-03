#### 安装

```shell
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

# 命令行不支持windows, 用docker
docker pull kjconroy/sqlc

docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc generate
```

