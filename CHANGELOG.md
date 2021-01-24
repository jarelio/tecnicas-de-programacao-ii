# Changelog

Todas as modificações neste projeto serão documentadas neste arquivo.

O formato é baseado em [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
e esse projeto adere à [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [2.0.0] - 2021-01-24

### Added

- Política de CORS no backend permitindo qualquer requisição e qualquer método
- Código do frontend feito em React com as operações de add, delete, update e list de grades
- Testes automatizados nas operações de add e update de uma grade usando behave, selenium e nose

### Updated

- Arquivo .gitignore para ignorar arquivos python
- README.md com os requerimentos e os passos para executar o frontend e os testes

### Fixed

- Campo de value das grades passa a ser string ao invés de float

## [1.1.0] - 2020-12-21

### Added

- Novo endpoint para consumir todas as grades de um estudante
- Router ao código de testes para validar as URLs usadas por eles
- Teste para o novo endpoint
- Função executada a cada teste para limpar a store utilizada pelo controlador

## [1.0.0] - 2020-12-17

### Added

- Pasta /backend contendo os controladores, serviços e testes para uma rest API de grades com as operações do CRUD.
- Passos para execução do serviço e dos testes no README.md

### Fixed

- Formatação dos arquivos baseada nos linters do go, code spell checker e prettier

## [0.0.1] - 2020-11-27

### Added

- Arquivo CHANGELOG.md contendo as informações iniciais adicionadas baseado em [Keep a change log](https://keepachangelog.com/pt-BR/0.3.0/).
- Arquivo de Licença MIT do projeto baseado em [Choose a License](https://choosealicense.com/licenses/mit/).
- Arquivo .gitignore gerado por [Toptal](https://www.toptal.com/developers/gitignore/api/visualstudiocode,go,react).
- Arquivo README.md contendo descrição inicial do projeto
