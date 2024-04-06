package dto

import (
	"fmt"
	"regexp"
	"strings"
)

// type Phone struct {
// 	Phone string
// }

type Phones struct {
	PhonesStrings []string
	MainNumber    string
	SnabNumber    string
	AllNumbers    string
}

func (p *Phones) MergePhones() {
	p.AllNumbers = strings.Join(p.PhonesStrings, "|")
}

func (p *Phones) ParsePhones() *Phones {

	p.SnabNumber = ""

	if len(p.PhonesStrings) == 0 {
		return p
	}

	if len(p.PhonesStrings) > 0 {
		p.MainNumber = ExtractPhoneNumber(p.PhonesStrings[0])
		p.MergePhones()
	}

	for _, phone := range p.PhonesStrings {
		//phone = CloseBrackets(phone)
		if checkForKeywords(phone) {
			p.SnabNumber = phone
		}
	}

	return p
}

// Функция для добавления недостающих закрывающих скобок
// func CloseBrackets(s string) string {
// 	openCount, closeCount := 0, 0
// 	for _, char := range s {
// 		if char == '(' {
// 			openCount++
// 		} else if char == ')' {
// 			closeCount++
// 		}
// 	}

// 	// Добавляем недостающие закрывающие скобки
// 	for openCount > closeCount {
// 		s += ")"
// 		closeCount++
// 	}

// 	return s
// }

// Функция для извлечения номера телефона из строки
func ExtractPhoneNumber(s string) string {

	// Регулярное выражение для поиска номеров телефонов
	// Поддерживает различные форматы записи номеров
	re := regexp.MustCompile(`\+?\d[\d\s()-]*\d`)

	// Поиск и возврат номера телефона
	return re.FindString(s)
}

// func TestPhoneNumber() {
// 	examples := []string{
// 		"(812)777-89-84 доб. 420 (Отдел снабжения)",
// 		"+7 (49624) 2-31-51 (доб 223) (Заместитель Генерального директора по развитию производства)",
// 		"8(812)520-62-87 (Отдел снабжения)",
// 		"+7 901 303 05 58 (Офис)",
// 	}

// 	// Применение функции к каждой строке из списка примеров
// 	for _, example := range examples {
// 		phoneNumber := extractPhoneNumber(example)
// 		fmt.Println(phoneNumber)
// 	}
// }

func TestSnab() {
	// Пример описания для проверки функции
	descriptions := []string{
		"Отдел снабжения",
		"Заместитель Генерального директора по закупкам",
		"Секретарь",
	}

	// Проверяем каждое описание
	for _, description := range descriptions {
		if checkForKeywords(description) {
			fmt.Printf("Описание \"%s\" содержит ключевые слова.\n", description)
		} else {
			fmt.Printf("Описание \"%s\" не содержит ключевых слов.\n", description)
		}
	}
}

func checkForKeywords(description string) bool {
	// Определяем список ключевых слов для поиска
	keywords := []string{"снаб", "закуп"}

	// Преобразуем описание в нижний регистр для регистронезависимого поиска
	descriptionLower := strings.ToLower(description)

	for _, keyword := range keywords {
		// Проверяем, содержит ли описание ключевое слово
		if strings.Contains(descriptionLower, keyword) {
			return true
		}
	}

	// Если ни одно из ключевых слов не найдено, возвращаем false
	return false
}
