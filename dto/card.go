package dto

import (
	"encoding/csv"
	"encoding/hex"
	"fmt"
	"hash/fnv"
	"log"
	"os"
	"regexp"

	"strings"
)

// название, категория, адрес, телефон, Эл. почта, Сайт, source, file

const CsvPath = "files/factories.csv"

// PCard is a data transfer object for pcard data.
type Card struct {
	Name      string
	Category  string
	Adress    string
	Email     string
	Site      string
	Source    string
	File      string
	Phone     string
	SnabPhone string
	Phones    string
}

// NewPCard creates a new PCard.
func NewCard() *Card {
	return &Card{}
}

// // String returns a string representation of the PCard.
func (p *Card) String() []string {

	p.Clear()

	return []string{p.Name, p.Category, p.Adress, p.Email, p.Site, p.Source, p.File, p.Phone, p.SnabPhone, p.Phones}

}

func (p *Card) Clear() *Card {

	occ := []string{
		"\n",
		"\t",
		"\r",
		"  ",
	}

	for _, r := range occ {
		p.Name = strings.Replace(p.Name, r, "", -1)
		p.Category = strings.Replace(p.Category, r, "", -1)
		p.Adress = strings.Replace(p.Adress, r, "", -1)
		p.Email = strings.Replace(p.Email, r, "", -1)
		p.Site = strings.Replace(p.Site, r, "", -1)
		p.Phone = strings.Replace(p.Phone, r, "", -1)
		p.SnabPhone = strings.Replace(p.SnabPhone, r, "", -1)
		p.Phones = strings.Replace(p.Phones, r, "", -1)
	}

	return p
}

func (p *Card) Hash() string {
	h := fnv.New32a()
	h.Write([]byte(p.Name))
	file := hex.EncodeToString(h.Sum(nil)) + ".html"
	p.File = file
	return file
}

func (p *Card) WriteCsv() error {
	file, err := os.OpenFile(CsvPath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	//writer.Comma = ';'
	//writer.UseCRLF = true
	defer writer.Flush()

	fmt.Println(p)

	writer.Write(p.String())

	return nil
}

func MapPhone() {
	// str := "(3843) 34–58–28 (отдел кадров)"

	// parts := strings.SplitN(str, " (", 2)

	// phone := parts[0]
	// description := strings.TrimPrefix(parts[1], "(")
	// description = strings.TrimSuffix(description, ")")

	// fmt.Println("Phone:", phone)
	// fmt.Println("Description:", description)

	// str := "+7 (495) 223-00-00 (доб. 306) (Сервисная служба фабрики)"

	// re := regexp.MustCompile(`(\+\d+ \(\d+\) \d+-\d+-\d+) \((доб\. \d+)\) \((.+)\)`)
	// matches := re.FindStringSubmatch(str)

	// if len(matches) != 4 {
	// 	fmt.Println("No match found")
	// 	return
	// }

	// phone := matches[1]
	// extension := matches[2]
	// description := matches[3]

	// fmt.Println("Phone:", phone)
	// fmt.Println("Extension:", extension)
	// fmt.Println("Description:", description)

	//str := "+7(34153)3-15-30 (Начальник отдела материально-технического снабжения)"
	// str := "+7 (495) 223-00-00 (доб. 306) (Сервисная служба фабрики)"

	// descriptionStart := strings.Index(str, "(")
	// descriptionEnd := strings.LastIndex(str, ")")

	// if descriptionStart == -1 || descriptionEnd == -1 {
	// 	fmt.Println("No description found")
	// 	return
	// }

	// description := str[descriptionStart+1 : descriptionEnd]

	// if strings.Contains(description, "снаб") || strings.Contains(description, "закуп") {
	// 	fmt.Println("Description contains 'снабжения' or 'закупки'")
	// } else {
	// 	fmt.Println("Description does not contain 'снабжения' or 'закупки'")
	// }

	// lines := []string{
	// 	"(812)777-89-84  доб. 420 (Отдел снабжения)",
	// 	"+7 (49624) 2-31-51 (доб 223) (Заместитель Генерального директора по развитию производства)",
	// 	"+7 (495) 223-00-00 (добавлена Сервисная служба фабрики)",
	// }

	// reWithExtension := regexp.MustCompile(`(\(?\+?\d+\)?\d+-\d+-\d+)  доб\. (\d+) \((.+)\)`)
	// reWithExtensionInBrackets := regexp.MustCompile(`(\(?\+?\d+\)?\d+-\d+-\d+) \((доб \d+)\) \((.+)\)`)
	// reWithoutExtension := regexp.MustCompile(`(\(?\+?\d+\)?\d+-\d+-\d+) \((.+)\)`)

	// for _, line := range lines {
	// 	matchesWithExtension := reWithExtension.FindStringSubmatch(line)
	// 	matchesWithExtensionInBrackets := reWithExtensionInBrackets.FindStringSubmatch(line)
	// 	matchesWithoutExtension := reWithoutExtension.FindStringSubmatch(line)

	// 	if len(matchesWithExtension) == 4 {
	// 		fmt.Println("Phone:", matchesWithExtension[1])
	// 		fmt.Println("Extension:", matchesWithExtension[2])
	// 		fmt.Println("Description:", matchesWithExtension[3])
	// 	} else if len(matchesWithExtensionInBrackets) == 4 {
	// 		fmt.Println("Phone:", matchesWithExtensionInBrackets[1])
	// 		fmt.Println("Extension:", matchesWithExtensionInBrackets[2])
	// 		fmt.Println("Description:", matchesWithExtensionInBrackets[3])
	// 	} else if len(matchesWithoutExtension) == 3 {
	// 		fmt.Println("Phone:", matchesWithoutExtension[1])
	// 		fmt.Println("No extension")
	// 		fmt.Println("Description:", matchesWithoutExtension[2])
	// 	} else {
	// 		fmt.Println("No match found for line:", line)
	// 	}

	// 	fmt.Println()
	// }

	// detect extension number in the string (доб 223)
	// line := "+7 (49624) 2-31-51 (доб 223) (Заместитель Генерального директора по развитию производства)"

	// re := regexp.MustCompile(`\(доб (\d+)\)`)
	// matches := re.FindStringSubmatch(line)

	// if len(matches) > 1 {
	// 	fmt.Println("The extension number is:", matches[1])
	// } else {
	// 	fmt.Println("The string does not contain an extension number.")
	// }

	// detect extension number in the string доб. 223
	// line2 := "+7 (49624) 2-31-51 доб. 223 (Заместитель Генерального директора по развитию производства)"

	// re2 := regexp.MustCompile(`доб\. (\d+)`)
	// matches2 := re2.FindStringSubmatch(line2)

	// if len(matches2) > 1 {
	// 	fmt.Println("The extension number is:", matches2[1])
	// } else {
	// 	fmt.Println("The string does not contain an extension number.")
	// }

}

// func DetectIf() {
// 		// Исходная строка
// 		text := "+7 (49624) 2-31-51 (доб 223) (Заместитель Генерального директора по развитию производства)"

// 		// Регулярное выражение для поиска номера телефона, добавочного номера и описания
// 		// Учтём возможные вариации форматирования номера и добавочного
// 		re := regexp.MustCompile(`(?P<phone>\(\d+\)\d+-\d+-\d+)( доб\. (?P<ext>\d+))?( \((?P<desc>.+)\))?`)

// 		// Использование регулярного выражения для поиска совпадений
// 		match := re.FindStringSubmatch(text)

// 		// Создание карты для хранения групп и их значений
// 		result := make(map[string]string)
// 		for i, name := range re.SubexpNames() {
// 			if i > 0 && i <= len(match) {
// 				result[name] = match[i]
// 			}
// 		}

// 		// Вывод результатов
// 		fmt.Println("Номер телефона:", result["phone"])
// 		if result["ext"] != "" {
// 			fmt.Println("Добавочный номер:", result["ext"])
// 		}
// 		if result["desc"] != "" {
// 			fmt.Println("Описание:", result["desc"])
// 		}
// }

// func DetectIfAddons() {
// 	lines := []string{
// 		"+7 (49624) 2-31-51 (доб 223) (Заместитель Генерального директора по развитию производства)",
// 		"+7 (49624) 2-31-51 доб. 223 (Заместитель Генерального директора по развитию производства)",
// 	}

// 	re1 := regexp.MustCompile(`\(доб (\d+)\)`)
// 	re2 := regexp.MustCompile(`доб\. (\d+)`)

// 	for _, line := range lines {
// 		matches1 := re1.FindStringSubmatch(line)
// 		matches2 := re2.FindStringSubmatch(line)

// 		if len(matches1) > 1 {
// 			fmt.Println("The extension number in the line is:", matches1[1])
// 		} else if len(matches2) > 1 {
// 			fmt.Println("The extension number in the line is:", matches2[1])
// 		} else {
// 			fmt.Println("The line does not contain an extension number.")
// 		}
// 	}
// }

// (812)777-89-84  доб. 420 (Отдел снабжения)
// +7 (49624) 2-31-51 (доб 223) (Заместитель Генерального директора по развитию производства)
// +7 (495) 223-00-00 (добавлена Сервисная служба фабрики)
// "+7 (49624) 2-31-51 (доб 223) (Заместитель Генерального директора по развитию производства)"
// "+7 (49624) 2-31-51 доб. 223 (Заместитель Генерального директора по развитию производства)"

func DetectIfAddons2() {
	examples := []string{
		"+7 (49624) 2-31-51 (доб 223) (Заместитель Генерального директора по развитию производства)",
		"+7 (495) 223-00-00 (добавлена Сервисная служба фабрики)",
		"8(812)520-62-87 (Отдел снабжения)",
		"+7 (812) 766-58-78",
		"8 (813 62) 27 918",
	}

	for _, example := range examples {
		phone, description := splitPhoneAndDescription(example)
		fmt.Printf("Номер телефона: %s\n", phone)
		if description != "" {
			fmt.Printf("Описание: %s\n", description)
		}
		fmt.Println("-----------")
	}

}

// Функция для разделения строки на номер телефона и описание
func splitPhoneAndDescription(text string) (string, string) {
	// Регулярное выражение для поиска описания в конце строки
	re := regexp.MustCompile(`^(.*?)(\s+\(([^()]+)\))$`)

	// Использование регулярного выражения для поиска совпадений
	match := re.FindStringSubmatch(text)

	if match != nil {
		return match[1], match[3] // Возврат номера телефона и описания
	}

	return text, "" // Возврат исходной строки, если описание не найдено
}
