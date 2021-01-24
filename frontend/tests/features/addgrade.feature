# language: pt

Funcionalidade: adicionar uma nova grade
  
  Contexto: acessar página de adicionar grades
    Dado que acesso a página de adicionar grades

  Cenário: acessar página de adicionar grade e adicionar uma grade
    Dado que preencho os campos das informações de adicionar uma grade
    Quando clico no botão de adicionar
    Então devo visualizar o resultado da inserção