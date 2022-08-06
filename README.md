# treco

Twitter トレンドの収集蓄積ツールです

## 準備

sample_envをコピーして.envもしくは.env_docker_local .env_prod を作成してください

- .env ローカルで動かす用
- .env_docker_local ローカルでdocker動かす用
- .env_prod prodで動かす用

## コンパイルと実行

```
CGO_ENABLED=0 go build -trimpath -ldflags '-s -w' -o treco
./treco
```

## Docker 

### build

```
docker build -t treco .
```

### run

```
docker run --rm treco ./treco
```

## Docker-Prod AWS ECR

### build

```
docker build -t treco_prod_x64 -f Dockerfile-Prod-x64 .
```
