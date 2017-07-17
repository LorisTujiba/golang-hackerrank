package main

import (
	"bufio"
	"fmt"
	"github.com/LorisTujiba/gotraining/3_scope/values"
	"os"
	"strings"
)

/*
https://www.hackerrank.com/challenges/ctci-ransom-note

Input Format

The first line contains two space-separated integers describing the respective values of  (the number of words in the magazine) and  (the number of words in the ransom note).
The second line contains  space-separated strings denoting the words present in the magazine.
The third line contains  space-separated strings denoting the words present in the ransom note.

Constraints

Each word consists of English alphabetic letters (i.e.,  to  and  to ).
The words in the note and magazine are case-sensitive.
Output Format

Print Yes if he can use the magazine to create an untraceable replica of his ransom note; otherwise, print No.*/

var magazineLen, ransomLen, used int
var magazine, ransom string
var splitMagazine, splitRansom []string
var magazineBuckets = make([][]string, 12)
var ransomBuckets = make([][]string, 12)
var notAvailable = ""

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &magazineLen)
	fmt.Fscan(reader, &ransomLen)
	reader.ReadString('\n')
	magazine, _ = reader.ReadString('\n')
	ransom, _ = reader.ReadString('\n')

	splitMagazine = strings.Fields(magazine)
	splitRansom = strings.Fields(ransom)

	for i := 0; i < magazineLen; i++ {
		magazineBuckets[hash(splitMagazine[i])] = append(magazineBuckets[hash(splitMagazine[i])], splitMagazine[i])
	}

	for i := 0; i < ransomLen; i++ {
		ransomBuckets[hash(splitRansom[i])] = append(ransomBuckets[hash(splitRansom[i])], splitRansom[i])
	}

	/*
		fmt.Println(magazineBuckets)
		fmt.Println(ransomBuckets)
		values.PrintSeparation("=", 60)
	*/
	fmt.Println(check())

}

func hash(word string) int {
	letter := int(word[0])
	buckets := letter % 12
	return buckets
}

func check() string {

	for i := 0; i < len(ransomBuckets); i++ {
		if len(ransomBuckets[i]) != 0 {

			/*
				fmt.Println("Ransom len : ",len(ransomBuckets[i]))
				fmt.Println("Not Available : ",notAvailable)
			*/

			for j := 0; j < len(ransomBuckets[i]); j++ {
				targetedValue := hash(ransomBuckets[i][j])

				if len(magazineBuckets[targetedValue]) != 0 {
					for n := 0; n < len(magazineBuckets[targetedValue]); n++ {
						if ransomBuckets[i][j] == magazineBuckets[targetedValue][n] {
							used++
							magazineBuckets[targetedValue] = append(magazineBuckets[targetedValue][:n], magazineBuckets[targetedValue][n+1:]...)
							/*
								fmt.Println("Ransom word : ",ransomBuckets[i][j])
								fmt.Println("iteration M: ",n," : ",magazineBuckets)
								fmt.Println("iteration R: ",n," : ",ransomBuckets)
							*/
							notAvailable = ""
							break
						}
						if n == len(magazineBuckets[targetedValue])-1 {
							return "No"
						}
					}
				} else {
					return "No"
				}
			}

		}
	}
	if used == ransomLen {
		return "Yes"
	}

	return "No"

}
