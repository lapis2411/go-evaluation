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
}

func (p Person) String() string {
	return fmt.Sprintf("%s(%d) %s", p.Name, p.Age, p.Job)
}
