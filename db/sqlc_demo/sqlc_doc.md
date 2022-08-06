#### 安装

```shell
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

# 命令行不支持windows, 用docker
docker pull kjconroy/sqlc

docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc generate
```

#### `QUERY ANNOTATIONS`

- :exec - 只执行
- :one - 返回一行记录
- :many - 返回多行记录

```shell

```

#### 参数

- emit_prepared_queries: 启用 `prepare`