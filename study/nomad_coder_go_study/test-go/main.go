package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

//func multiply(a int, b int)
func multiply(a, b int) int {
	return a * b
}

/*
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}
*/

func repaetMe(words ...string) {
	fmt.Println(words)
}

//naked return & defer(함수가 끝나고 임의의 동작이 필요할 때)
func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("잘했어요")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func superAdd(numbers ...int) int {
	//fmt.Println(numbers)
	total := 0
	//for i:= 0; i < len(number); i++ {
	for _, number := range numbers {
		//fmt.Println(index, number)
		//total = total + number
		total += number
	}
	return total
}

func canIDrink(age int) bool {
	/*
		if age < 18 {
			return false
		}
		return true
	*/

	//variable expression
	/*
		if koreanAge := age + 2; koreanAge < 18 {
			return false
		}
		return true
	*/

	switch age {
	case 10:
		return false
	case 20:
		return true
	}

	switch {
	case age < 18:
		return false
	case age == 40:
		return true
	case age > 50:
		return false
	}

	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 11:
		return true
	}

	return false

}

var errRequestFailed = errors.New("Request failed")

type requestResult struct {
	url    string
	status string
}

//func main() {
// Hello
//fmt.Println("Hello world")

//someting.SayHello()

/*
	const name string = "nico"
	name = "Lynn" 에러 발생
*/

/*
	var name string = "nico"
	name := "nico" 는 var name string = "nico"와 같은 의미이다. func안에서 변수 선언에서만 사용 가능
	name = "lynn"
	fmt.Println(name)
*/

//fmt.Println(multiply(2, 2))

/*
	totalLenght, upperName := lenAndUpper("jeongJiHyun")
	fmt.Println(totalLenght, upperName)
	totalLenght, _ := lenAndUpper("jeongJiHyun")
	fmt.Println(totalLenght)
*/

//repaetMe("jihyun", "jusung", "doraemo")
/*
	totalLenght, upperName := lenAndUpper("jeonjihyun")
	fmt.Println(totalLenght, upperName)

	total := superAdd(1, 2, 3, 4, 5)
	fmt.Println(total)

	fmt.Println(canIDrink(16))

	a := 2
	b := &a
	a = 10
	fmt.Println(&a, *b)

	names := []string{"nico", "jihuyn", "dodo"}
	names = append(names, "ffff")
	fmt.Println(names)

	jihyun := map[string]string{"name": "jehyun", "age": "12"}
	fmt.Println(jihyun)

	for key, value := range jihyun {
		fmt.Println(key, value)
	}
*/

/*
	account := accounts.Newaccount("jihyun")
	account.Deposit(10)
	err := account.Withdraw(5)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account)
	fmt.Println(account)
*/

/*
	dictionary := mydict.Dictionary{"first": "First word"}
	dictionary["hello"] = "hello"
	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
*/

/*
	dictionary := mydict.Dictionary{}
	word := "hello"
	definition := "Greeting"
	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}
	hello, _ := dictionary.Search(word)
	fmt.Println(hello)
	err2 := dictionary.Add(word, definition)
	if err2 != nil {
		fmt.Print(err2)
	}
*/

/*
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	}
	//word, _ := dictionary.Search(baseWord)
	fmt.Print(word)
*/
/*
	var results = make(map[string]string)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
	}
	for _, url := range urls {
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}
	for url, result := range results {
		fmt.Println(url, result)
	}
*/

/*
	c := make(chan string)
	people := [2]string{"nico", "flynn"}
	for _, person := range people {
		go isSexy(person, c)
	}

	for i := range people {
		fmt.Print("wating for", i)
		fmt.Println(i, " Received this message:", <-c)
	}
*/
//go sexyCount("nico")
//go sexyCount("to")

/*
	var results = make(map[string]string)
	var c = make(chan requestResult)

	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
	}
	for _, url := range urls {
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
*/

//}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 5)
	c <- person + " is sexy"

}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}

func hitURL(url string, c chan<- requestResult) {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAIL"
	}
	c <- requestResult{url: url, status: status}
}

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

func main() {
	pages := getPages()
	fmt.Println(pages)
}

func getPages() int {
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	fmt.Println(doc)
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Html())
	})
	return 0
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Requests failed with Status:", res.StatusCode)
	}
}
