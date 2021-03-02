from behave import *
from nose.tools import assert_equal
from pages.addgrade_page import AddGradePage
from pages.grades_page import GradesPage
from pages.update_delete_grade_page import UpdateAndDeleteGradePage

gradesPage = GradesPage()
deleteGradePage = UpdateAndDeleteGradePage()

@given(u'que capturo a quantidade de grades existentes')
def step_impl(context):
    gradesPage.gradesNumber = gradesPage.get_quantity_of_grades()

@when(u'clico no botão de deletar')
def step_impl(context):
    deleteGradePage.click_delete_button()


@then(u'devo visualizar a página de grades com menos uma grade')
def step_impl(context):
    gradesNumberAfterDelete = gradesPage.get_quantity_of_grades()
    gradesNumberBeforeDelete = gradesPage.gradesNumber
    assert_equal(gradesNumberAfterDelete, gradesNumberBeforeDelete-1)
