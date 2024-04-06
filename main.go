package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"parseFactories/dto"

	//"parseFactories/parsesample"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

//var CsvPath = "files/factories.csv"

func init() {

	path := dto.CsvPath

	// Removing file from the directory
	// if exists
	// if _, err := os.Stat(path); err == nil {
	// 	return
	// }
	//e := os.Remove(path)
	//if e != nil {
	//	log.Fatal(e)
	//}

	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	// writer.Comma = ';'
	// writer.UseCRLF = true
	defer writer.Flush()

	writer.Write(GetHeaders())
}

func GetHeaders() []string {
	return []string{"name", "category", "adress", "email", "site", "source", "file", "phone", "snabphone", "phones"}
}

func main() {

	//parsesample.ParseS()
	crawl()
}

func crawl() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("заводы.рф"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {

		link := e.Attr("href")

		// convert relative url to absolute
		url := e.Request.AbsoluteURL(link)

		visit, err := OnLink(url)
		if err != nil {
			fmt.Println(err)
			return
		}

		if !visit {
			return
		}

		// Visit link found on page on a new thread
		c.Visit(url)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// After making a request print "Visited ..."
	c.OnResponse(func(r *colly.Response) {

		page, err := OnPage(r)
		if err != nil {
			fmt.Println(err)
			return
		}

		if !page {
			return
		}

		fmt.Println("THIS IS THE PAGE!!!!", r.Request.URL)
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://заводы.рф")
}

// func parseInfo() {

// }

// func savePage() {

// }

func OnLink(url string) (bool, error) {

	if strings.Contains(url, "/exporters") {
		return false, nil
	}
	if strings.Contains(url, "/products") {
		return false, nil
	}
	if strings.Contains(url, "/news") {
		return false, nil
	}
	return true, nil
}

func OnPage(e *colly.Response) (bool, error) {

	url := e.Request.URL.String()

	fmt.Println("CHECK IF PAGE", url)
	if !bytes.Contains(e.Body, []byte(`<a class="active" href="#company-descr">О компании</a>`)) {
		// just skip this url, no errors triggered
		return false, nil
	}

	product, err := ParsePage(bytes.NewBuffer(e.Body), url)
	if err != nil {
		return true, err
	}

	product.Source = url

	srcfile := product.Hash()
	srcfile = strings.Replace(srcfile, "oy", "", -1)

	// err = SavePage(srcfile, string(e.Body))
	// if err != nil {
	// 	return true, err
	// }

	product.WriteCsv()

	// err = WriteCsv(dto.CsvPath, product)
	// if err != nil {
	// 	return true, err
	// }

	return true, nil
}

func ParsePage(html *bytes.Buffer, source string) (*dto.Card, error) {
	// название, категория, адрес, телефон, Эл. почта, Сайт, source, file

	doc, err := goquery.NewDocumentFromReader(html)
	if err != nil {
		return nil, err
	}

	card := dto.NewCard()

	// parse name from div.company-head-title h1 text trim space
	card.Name = strings.TrimSpace(doc.Find("div.company-head-title h1").Text())

	// parse category from ul.content-list first li span.content-list__descr text trim space
	categoryli := doc.Find("ul.content-list li").First()
	card.Category = strings.TrimSpace(categoryli.Find("span.content-list__descr").Text())

	//address := ""

	phones := dto.Phones{}

	doc.Find("div#contact-company ul.content-list li").Each(func(i int, s *goquery.Selection) {
		// if s isset span.icon-map then address = s.content-list__descr text trim space
		if s.Find("span.icon-map").Length() > 0 {
			card.Adress = strings.TrimSpace(s.Find("span.content-list__descr").Text())
		}

		// if s.Find("span.icon-mail").Length() > 0 {
		// 	email = strings.TrimSpace(s.Find("span.content-list__descr").Text())
		// }

		if s.Find("span.icon-mail").Length() > 0 {
			//email = strings.TrimSpace(s.Find("span.content-list__descr").Text())
			s.Find("span.content-list__descr span.__cf_email__").Each(func(j int, t *goquery.Selection) {
				card.Email += Cf(t.AttrOr("data-cfemail", "")) + "|"
				//phone += strings.TrimSpace(strings.Replace(t.Text(), ",", " ", -1)) + "|"
			})
		}

		if s.Find("span.icon-site").Length() > 0 {
			card.Site = strings.TrimSpace(s.Find("span.content-list__descr").Text())
		}

		if s.Find("span.icon-tel").Length() > 0 {
			//phone = strings.TrimSpace(s.Find("span.content-list__descr").Text())
			s.Find("span.content-list__descr div").Each(func(j int, t *goquery.Selection) {
				phones.PhonesStrings = append(phones.PhonesStrings, strings.TrimSpace(t.Text()))
			})
		}
	})

	pp := phones.ParsePhones()

	card.Phone = pp.MainNumber
	card.SnabPhone = pp.SnabNumber
	card.Phones = pp.AllNumbers

	//card.Adress = strings.Replace(address, ";", ".", -1)

	//file := Hash(name) + ".html"

	return card, nil
}

func WriteCsv(path string, product []string) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Comma = ';'
	defer writer.Flush()

	writer.Write(product)

	return nil
}

// redeclare writer
// func NewWriter(w io.Writer) (writer *csv.Writer) {
//     writer = csv.NewWriter(w)
//     writer.Comma = '\t'

//     return
// }

func SavePage(filename string, html string) error {

	path := "files/pages/" + filename

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(html)
	if err != nil {
		return err
	}

	return nil

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

func ReplaceString(s string, occ []string) string {

	for _, oldString := range occ {
		s = strings.Replace(s, oldString, "", -1)
	}

	return strings.TrimSpace(strings.Replace(s, "\n", "", -1))
}
