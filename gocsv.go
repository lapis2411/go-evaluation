package main

import (
	"fmt"
	"log"

	"github.com/gocarina/gocsv"
)

type (
	Person struct {
		Name string
		Age  int
		Job  string
	}
	// 大文字にしないとgocsvから見えない
	PersonCSV struct {
		FirstName  string `csv:"名"`
		SecondName string `csv:"姓"`
		Age        int    `csv:"年齢"`
		Job        string `csv:"職業"`
	}
)

func main() {
	test := `姓,名,年齢,職業
	山田,太郎,30,エンジニア
	佐藤,花子,25,デザイナー
	田中,一郎,35,医師
	渡辺,和子,28,教師
	高橋,健二,32,自由業`
	var p []Person
	if err := gocsv.UnmarshalBytesToCallback([]byte(test), func(pc PersonCSV) {
		p = append(p, Person{
			Name: pc.FirstName + pc.SecondName,
			Age:  pc.Age,
			Job:  pc.Job,
		})
	}); err != nil {
		log.Fatalln(err)
	}
	for _, v := range p {
		fmt.Println(v)
	}
	fmt.Println("pattern 1 done")

	// 処理が複雑な場合1
	var p2 []Person
	cp := ComplexProcess(&p2)
	if err := gocsv.UnmarshalBytesToCallback([]byte(test), cp); err != nil {
		log.Fatalln(err)
	}
	for _, v := range p2 {
		fmt.Println(v)
	}
	fmt.Println("pattern 2 done")

	// 処理が複雑な場合2
	var p3 []Person
	if err := gocsv.UnmarshalBytesToCallback([]byte(test), func(pc PersonCSV) {
		ComplexProcess2(&p3, pc)
	}); err != nil {
		log.Fatalln(err)
	}
	for _, v := range p3 {
		fmt.Println(v)
	}
	fmt.Println("pattern 3 done")
}

func (p Person) String() string {
	return fmt.Sprintf("%s(%d) %s", p.Name, p.Age, p.Job)
}

// 処理が複雑な場合は関数を返すようにする
func ComplexProcess(person *[]Person) func(PersonCSV) {
	return func(pc PersonCSV) {
		*person = append(*person, Person{
			Name: pc.FirstName + pc.SecondName,
			Age:  pc.Age,
			Job:  pc.Job,
		})
		// さらに何らかの処理。。。
	}
}

func ComplexProcess2(person *[]Person, personCSV PersonCSV) {
	// pに値を加工して代入するような何らかの複雑な処理
	*person = append(*person, Person{
		Name: personCSV.FirstName + personCSV.SecondName,
		Age:  personCSV.Age,
		Job:  personCSV.Job,
	})
}
