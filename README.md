# 活动api

## 快速开始
```
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
git clone https://github.com/hildxd/activity-api
cd activity-api
make init
go generate ./...
go build -o ./bin/ ./...
./bin/server -conf ./configs
```

## Docker
```bash
# build
docker build -t <your-docker-image-name> .

# run
docker run --rm -p 8000:8000 -p 9000:9000 -v </path/to/your/configs>:/data/conf <your-docker-image-name>
```
