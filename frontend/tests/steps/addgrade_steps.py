from behave import *
from nose.tools import assert_equal
from pages.addgrade_page import AddGradePage

addGradePage = AddGradePage()


@given(u'que acesso a página de adicionar grades')
def step_impl(context):
    addGradePage.acess_page('http://localhost:3000/add')


@given(u'que preencho os campos das informações de adicionar uma grade')
def step_impl(context):
    addGradePage.send_keys_grades_inputs()


@when(u'clico no botão de adicionar')
def step_impl(context):
    addGradePage.click_submit_button()


@then(u'devo visualizar o resultado da inserção')
def step_impl(context):
    assert_equal(addGradePage.get_result_text_title(),
                 'You submitted successfully!')
