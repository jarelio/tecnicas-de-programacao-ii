# language: pt

Funcionalidade: atualizar uma grade
  
  Contexto: acessar página de grades
    Dado que acesso a página de grades

  Cenário: acessar página de grades, selecionar uma grade e atualizar os dados dela
    Dado que escolho a primeira grade
    Dado que clico no botão de editar
    Dado que preencho os campos das informações da grade
    Quando clico no botão de atualizar
    Então devo visualizar o resultado da atualização

  Cenário: acessar página de grades, visualizar a primeira grade atualizada e validar os dados
    Dado que volto para a página de grades
    Quando escolho a primeira grade
    Então devo visualizar os dados da grade atualizada