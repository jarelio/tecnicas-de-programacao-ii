from behave import *
from nose.tools import assert_equal
from pages.grades_page import GradesPage
from pages.update_delete_grade_page import UpdateAndDeleteGradePage

gradesPage = GradesPage()
editGradePage = UpdateAndDeleteGradePage()


@given(u'que acesso a página de grades')
def step_impl(context):
    gradesPage.acess_page('http://localhost:3000/grade')


@given(u'que escolho a primeira grade')
def step_impl(context):
    gradesPage.select_first_grade()


@given(u'que clico no botão de editar')
def step_impl(context):
    gradesPage.click_edit_button()


@given(u'que preencho os campos das informações da grade')
def step_impl(context):
    editGradePage.send_keys_grades_inputs()


@when(u'clico no botão de atualizar')
def step_impl(context):
    editGradePage.click_edit_button()


@then(u'devo visualizar o resultado da atualização')
def step_impl(context):
    assert_equal(editGradePage.get_result_text_title(),
                 'The grade was updated successfully!')


@given(u'que volto para a página de grades')
def step_impl(context):
    editGradePage.click_grades_button()


@when(u'escolho a primeira grade')
def step_impl(context):
    gradesPage.select_first_grade()


@then(u'devo visualizar os dados da grade atualizada')
def step_impl(context):
    grade = gradesPage.get_first_grade_data()
    assert_equal(grade["student"],
                 'Student 2')
    assert_equal(grade["subject"],
                 'Subject 2')
    assert_equal(grade["type"],
                 'Type 2')
    assert_equal(grade["value"],
                 '20')
