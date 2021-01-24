from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as ec
from browser import Browser


class GradesPageLocator(object):
    # Seletores dos elementos utilizados na página
    FIRST_GRADE = '//*[@id="root"]/div/div/div/div[2]/ul/li[1]'
    EDIT_BUTTON = '//*[@id="root"]/div/div/div/div[3]/div/a'
    STUDENT_DATA = '//*[@id="root"]/div/div/div/div[3]/div/div[1]'
    SUBJECT_DATA = '//*[@id="root"]/div/div/div/div[3]/div/div[2]'
    TYPE_DATA = '//*[@id="root"]/div/div/div/div[3]/div/div[3]'
    VALUE_DATA = '//*[@id="root"]/div/div/div/div[3]/div/div[4]'


class GradesPage(Browser):
    def get_element(self, locator):
        # aguarda elemento estar visível na tela
        WebDriverWait(self.driver, 10).until(
            ec.visibility_of_element_located((By.XPATH, locator)))
        # retorna elemento
        return self.driver.find_element(By.XPATH, locator)

    def acess_page(self, url):
        # acessa url passada
        self.driver.get(url)

    def select_first_grade(self):
        first_grade = self.get_element(GradesPageLocator.FIRST_GRADE)
        first_grade.click()

    def click_edit_button(self):
        # clica no botão edit
        button = self.get_element(GradesPageLocator.EDIT_BUTTON)
        button.click()

    def get_grade_data(self):
        # retorna os textos dos elementos
        student = self.get_element(
            GradesPageLocator.STUDENT_DATA).text.split(' ', 1)[1]
        subject = self.get_element(
            GradesPageLocator.SUBJECT_DATA).text.split(' ', 1)[1]
        type_data = self.get_element(
            GradesPageLocator.TYPE_DATA).text.split(' ', 1)[1]
        value = self.get_element(
            GradesPageLocator.VALUE_DATA).text.split(' ', 1)[1]

        grade = {"student": student, "subject": subject,
                 "type": type_data, "value": value}
        return grade
