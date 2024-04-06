package parsesample

import (
	//"bytes"

	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ParseS() {

	fileContent, err := ioutil.ReadFile("parsesample/s.html")
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string
	//text := string(fileContent)
	//fmt.Println(text)

	// Use strings.NewReader to create a reader from the string
	//reader := strings.NewReader(text)

	// Use goquery.NewDocumentFromReader to read the reader
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(fileContent)))
	if err != nil {
		log.Fatal(err)
	}

	//doc, err := goquery.NewDocumentFromReader(bytes.NewBufferString("parsesample/s.html")) /home/andrew/Рабочий стол/go_pr/parseFactories/parsesample/s.html
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // parse name from div.company-head-title h1 text trim space
	name := strings.TrimSpace(doc.Find("div.company-head-title h1").Text())

	// parse category from ul.content-list first li span.content-list__descr text trim space
	categoryli := doc.Find("ul.content-list li").First()
	category := strings.TrimSpace(categoryli.Find("span.content-list__descr").Text())

	address := ""
	phone := ""
	email := ""
	site := ""

	doc.Find("div#contact-company ul.content-list li").Each(func(i int, s *goquery.Selection) {
		// if s isset span.icon-map then address = s.content-list__descr text trim space
		if s.Find("span.icon-map").Length() > 0 {
			address = strings.TrimSpace(s.Find("span.content-list__descr").Text())
		}

		// if s.Find("span.icon-mail").Length() > 0 {
		// 	email = strings.TrimSpace(s.Find("span.content-list__descr").Text())
		// }

		if s.Find("span.icon-mail").Length() > 0 {
			//email = strings.TrimSpace(s.Find("span.content-list__descr").Text())
			s.Find("span.content-list__descr span.__cf_email__").Each(func(j int, t *goquery.Selection) {
				email += Cf(t.AttrOr("data-cfemail", "")) + "|\n"
				//phone += strings.TrimSpace(strings.Replace(t.Text(), ",", " ", -1)) + "|"
			})
		}

		if s.Find("span.icon-site").Length() > 0 {
			site = strings.TrimSpace(s.Find("span.content-list__descr").Text())
		}

		if s.Find("span.icon-tel").Length() > 0 {
			//phone = strings.TrimSpace(s.Find("span.content-list__descr").Text())
			s.Find("span.content-list__descr div").Each(func(j int, t *goquery.Selection) {
				phone += strings.TrimSpace(strings.Replace(t.Text(), ",", " ", -1)) + "|\n"
			})
		}
	})

	cfMail := Cf("0e6c617765614e6f657c7b6c6760207c7b")

	fmt.Println("name :", name)
	fmt.Println("category :", category)
	fmt.Println("address :", address)

	fmt.Println("phone :", phone)

	fmt.Println("SFmail :", cfMail)

	fmt.Println("site :", site)

	fmt.Println("email :", email)

}

func Cf(a string) (s string) {
	var e bytes.Buffer
	r, _ := strconv.ParseInt(a[0:2], 16, 0)
	for n := 4; n < len(a)+2; n += 2 {
		i, _ := strconv.ParseInt(a[n-2:n], 16, 0)
		e.WriteString(string(i ^ r))
	}
	return e.String()
}
