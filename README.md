# Testes de carga

Este repositório apresenta o desenvolvimento de um `Teste de Carga` como parte do desafio proposto para cumprimento da Pós-graduação utilizando linguagem de programação [Go](https://go.dev).

---

## Como executar

- **Primeira opção:**

```shell
go run main.go --url http://www.google.com --concurrency=10 --requests=100
```

- **Segunda opção:**

Para criar a imagem:

```shell
docker build -t loadtester .
```

Para eecutar a imagem criada:

```shell
docker run loadtester ./loadtester --url=http://www.google.com --requests=100 --concurrency=10
```
