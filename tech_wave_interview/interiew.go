package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("Hello World")

	file, err := readCSV("path/c.csv")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	b := make([]byte, 20)
	file.Read(b)
	//s1, c, 2010, 85, A = []byte

	// read each line in a CSV file - print if marks more than 80

	//s := string(b)
	//
	//arS := strings.Split(s, ", ")
	//
	//marks := arS[3]

	// strconv.ParseInt(marks, 10)

}

func readCSV(path string) (*os.File, error) {
	// read csv file and return link
	return os.Open(path)

}

func getstudents() {
	//for link.Next {
	//	// for each row get studentID, language,year,  marks, grade
	//	// if marks > 80 {
	//
	//	fmt.Print("studentID", "language")
	//	// }
	//}
}

// studentID course yearOfPassing

// get num of grad by year

// studentNum year

// select count(*) as studentNum, year  from Student group by year; -- group by

// select s1.studentID ,s2.course, s1.course from Student S1 join Student S2 on S1.StudentID = S2.StudentID where course = 'c' and course = 'JAVA'
// group by s1.StudentID;

// s1 , java, c -

// select studentID from (
// select distinct s1.studentID ,s2.course, s1.course from Student S1 join Student S2 on S1.StudentID = S2.StudentID where course = 'c' and course = 'JAVA';
// ) -- self join
