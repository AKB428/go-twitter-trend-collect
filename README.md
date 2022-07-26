# treco

Twitter トレンドの収集蓄積ツールです


## Docker 

### build

```
docker build -t treco .
```

### run

```
docker run --rm -i treco ./treco
```


## Docker-Prod

### build

```
docker build -t treco_prod -f Dockerfile-Prod .
```

### run

```
docker run --rm -i treco_prod ./treco
```

## Docker-Prod AWS ECR

### build

```
docker build -t treco_prod_x64 -f Dockerfile-Prod-x64 .
```
