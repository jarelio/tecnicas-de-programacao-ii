from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as ec
from browser import Browser


class UpdateGradePageLocator(object):
    # Seletores dos elementos utilizados na página
    GRADES_BUTTON = '//*[@id="root"]/div/nav/div/li[1]/a'
    INPUT_STUDENT = '//*[@id="student"]'
    INPUT_SUBJECT = '//*[@id="subject"]'
    INPUT_TYPE = '//*[@id="type"]'
    INPUT_VALUE = '//*[@id="value"]'
    EDIT_BUTTON = '//*[@id="root"]/div/div/div/div/button[2]'
    RESULT_TITLE = '//*[@id="root"]/div/div/div/div/p'


class UpdateGradePage(Browser):
    def get_element(self, locator):
        # aguarda elemento estar visível na tela
        WebDriverWait(self.driver, 10).until(
            ec.visibility_of_element_located((By.XPATH, locator)))
        # retorna elemento
        return self.driver.find_element(By.XPATH, locator)

    def acess_page(self, url):
        # acessa url passada
        self.driver.get(url)

    def send_keys_grades_inputs(self):
        # envia para o elemento o texto 'Student 1'
        input_student = self.get_element(UpdateGradePageLocator.INPUT_STUDENT)
        input_student.clear()
        input_student.send_keys('Student 2')

        # envia para o elemento o texto 'Subject 1'
        input_subject = self.get_element(UpdateGradePageLocator.INPUT_SUBJECT)
        input_subject.clear()
        input_subject.send_keys('Subject 2')

        # envia para o elemento o texto 'Type 1'
        input_type = self.get_element(UpdateGradePageLocator.INPUT_TYPE)
        input_type.clear()
        input_type.send_keys('Type 2')

        # envia para o elemento o numero 10
        input_value = self.get_element(UpdateGradePageLocator.INPUT_VALUE)
        input_value.clear()
        input_value.send_keys(20)

    def click_edit_button(self):
        # clica no botão submit
        button = self.get_element(UpdateGradePageLocator.EDIT_BUTTON)
        button.click()

    def click_grades_button(self):
        button = self.get_element(UpdateGradePageLocator.GRADES_BUTTON)
        button.click()

    def get_result_text_title(self):
        # retorna o texto do elemento
        element = self.get_element(UpdateGradePageLocator.RESULT_TITLE)
        return element.text
