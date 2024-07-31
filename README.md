<p align="center">
  <img src="https://cdn-icons-png.flaticon.com/512/11150/11150628.png" width="100" />
</p>
<p align="center">
    <h1 align="center">GO-EXPERT STRESS TEST</h1>
</p>
<p align="center">
    <em>Desafio da pós em GO Expert</em>
</p>
<p align="center">
	<img src="https://img.shields.io/github/last-commit/mhayk/GO-Expert-rate-limiter?style=flat&logo=git&logoColor=white&color=0080ff" alt="last-commit">
	<img src="https://img.shields.io/github/languages/top/mhayk/GO-Expert-rate-limiter?style=flat&color=0080ff" alt="repo-top-language">
	<img src="https://img.shields.io/github/languages/count/mhayk/GO-Expert-rate-limiter?style=flat&color=0080ff" alt="repo-language-count">
<p>
<p align="center">
    <em>Developed with ❤️ by Mhayk Whandson</em>
</p>
<p align="center">
		<em>Developed with the language, software and tools below.</em>
</p>
<p align="center">
	<img src="https://img.shields.io/badge/YAML-CB171E.svg?style=flat&logo=YAML&logoColor=white" alt="YAML">
	<img src="https://img.shields.io/badge/V8-4B8BF5.svg?style=flat&logo=V8&logoColor=white" alt="V8">
	<img src="https://img.shields.io/badge/Docker-2496ED.svg?style=flat&logo=Docker&logoColor=white" alt="Docker">
	<img src="https://img.shields.io/badge/Go-00ADD8.svg?style=flat&logo=Go&logoColor=white" alt="Go">
</p>
<hr>

# Desafio Go: Stress Test

Este projeto é um sistema desenvolvido em Go para realizar testes de carga em serviços web. Ele permite configurar a URL do serviço a ser testado, definir o número total de requisições e a quantidade de requisições simultâneas. Ao final dos testes, o sistema gera um relatório que inclui informações como o tempo total de execução, o número total de requisições efetuadas e a distribuição dos status HTTP recebidos.

## Pré-requisitos

Certifique-se de ter os seguintes componentes instalados em sua máquina:

- Go (versão 1.22.5 ou superior)
- Docker (opcional, para execução em contêineres Docker)

## Instalação e Configuração

### Clonando o Repositório

```bash
git clone https://github.com/mhayk/GO-Expert-stress-test.git
cd GO-Expert-stress-test
```

## Compilação e Execução

### Executando Diretamente do Código-Fonte

1. **Compilar o código-fonte:**

   ```bash
   go build -o stress-test ./cmd/stress-test
   ```

   or

   ```bash
   make build
   ```

2. **Executar o teste de carga:**

   Configure os parâmetros `--url`, `--requests` e `--concurrency` conforme necessário.

   ```bash
   ./stress-test --url=https://google.com --requests=100 --concurrency=10
   ```

   or

   ```bash
   make run
   ```

### Executando via Docker

1. **Construir a imagem Docker:**

   ```bash
   docker build -t go-expert-stress-test .
   ```

2. **Executar o teste de carga utilizando Docker:**

   Ajuste os valores de `--url`, `--requests` e `--concurrency` conforme necessário.

   ```bash
   docker run go-expert-stress-test --url=http://google.com --requests=100 --concurrency=10
   ```

## Resultados do Teste

```
 ___  ____  ____  ____  ___  ___    ____  ____  ___  ____
/ __)(_  _)(  _ \( ___)/ __)/ __)  (_  _)( ___)/ __)(_  _)
\__ \  )(   )   / )__) \__ \\__ \    )(   )__) \__ \  )(
(___/ (__) (_)\_)(____)(___/(___/   (__) (____)(___/ (__)
------------------------------------------------------------
Relatório de teste de stress
Tempo total gasto: 8.543427291s
Total de requests: 100
Requests com status 200: 76
Status 0: 24
------------------------------------------------------------
```

Ao término da execução, o sistema gera um relatório detalhado que inclui:

- Tempo total de execução
- Número total de requisições realizadas
- Quantidade de requisições com status HTTP 200
- Distribuição dos outros códigos de status HTTP (como 404, 500, etc.)

Este README oferece uma visão geral do projeto, além de fornecer instruções de configuração e execução, e informações sobre como interpretar os resultados dos testes de carga.



### Dependências de desenvolvimento

- [Mockery](https://vektra.github.io/mockery/latest/)

```bash
$ go install github.com/vektra/mockery/v2@v2.20.0
$ mockery
```

## Executando os testes unitários

Para rodar os testes unitários, execute o comando a seguir:

```bash
$ make test
```