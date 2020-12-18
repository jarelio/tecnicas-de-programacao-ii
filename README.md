# Técnicas de Programação II

## Descrição do Projeto

Este projeto será uma aplicação simples de controle de notas de trabalhos/atividades com todas as operações de CRUD (Create, Read, Update, Delete) utilizando **React** no frontend e uma REST API feita em **GO** no backend. O intuito inicialmente desse projeto é desenvolver o backend e melhorar o frontend, se necessário, sendo guiado por testes (TDD). Utilizar a linguagem **GO** ajudará neste processo pois ela é simples e possui testes nativamente.

- **Equipe:**
  - Jarélio Filho - 399683

As convenções utilizadas nos(as) commits/branchs serão:

- Criar uma branch diferente para cada atualização (bugfix ou feature) e fazer o merge dela após ser revisada.
- O primeiro commit de uma atividade deverá ser da forma **[#X] - Nome da Atividade**, os próximos são livres. Podendo **X** ser:
  - **FEAT** (Qualquer feature nova)
  - **FIX** (Qualquer bugfix ou refatoração)
  - **TEST** (Qualquer código relacionado a testes)
  - **UTIL** (Qualquer modificação útil. Ex: Modificação do README, CHANGELOG)
- Caso a atividade se encaixe em mais de um dos títulos é possível citar todos separando por "\_", por exemplo [#FEAT_TEST].

## Passos iniciais

```properties
curl -O "https://dl.google.com/go/go1.15.5.linux-amd64.tar.gz"

tar -C /usr/local -xzf go1.15.5.linux-amd64.tar.gz

export PATH=$PATH:/usr/local/go/bin
```

## Como executar a aplicação

```go
go run main.go
```

## Como executar os testes

```go
go test ./tests/
```
