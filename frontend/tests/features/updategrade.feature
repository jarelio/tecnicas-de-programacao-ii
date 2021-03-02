# language: pt

Funcionalidade: atualizar uma grade
  
  Contexto: Adicionar uma grade, acessar página de grades
    Dado que acesso a página de adicionar grades
    Dado que preencho os campos das informações de adicionar uma grade
    Quando clico no botão de adicionar
    Dado que acesso a página de grades

  Cenário: acessar página de grades, selecionar uma grade e atualizar os dados dela
    Dado que escolho a primeira grade
    Dado que clico no botão de editar
    Dado que preencho os campos das informações da grade
    Quando clico no botão de atualizar
    Então devo visualizar o resultado da atualização
    Dado que volto para a página de grades
    Quando escolho a primeira grade
    Então devo visualizar os dados da grade atualizada

