# language: pt

Funcionalidade: deletar uma grade
  
  Contexto: Adicionar uma grade, acessar página de grades
    Dado que acesso a página de adicionar grades
    Dado que preencho os campos das informações de adicionar uma grade
    Quando clico no botão de adicionar
    Dado que acesso a página de grades

  Cenário: página de grades acessada, selecionar uma grade e deletá-la
    Dado que escolho a primeira grade
    Dado que capturo a quantidade de grades existentes
    Dado que clico no botão de editar
    Quando clico no botão de deletar
    Então devo visualizar a página de grades com menos uma grade