package main

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"math/bits"
	"time"
)

func main() {

	fmt.Println("Welcome to Coinflip!")
	var input int
	fmt.Println("New Coinflip every 10 seconds >>11<<\nCheck randomness >>22<<")
	_, err := fmt.Scanf("%d", &input)
	if err != nil {
		fmt.Println(err)
	}
	if input == 11 {
		fmt.Println("Every 10 seconds new coinflip")
		for { //endless loop
			if headortails() == 0 { //if return value is 0
				fmt.Println("HEADS!") //its heads
			} else { //if return value is 1
				fmt.Println("TAILS!") //its tails
			}
			fmt.Println("Timestamp: ", time.Now().UTC()) //make a timestamp for better performance
			time.Sleep(10 * time.Second)                 //let it sleep for 5 sec and do again
		}
	} else if input == 22 {
		checkRandomness() //to check randomness
	}
}

func checkRandomness() { //to check if its really random
	var counter1, counter2 int     //initialise new variables to count tails or heads
	for i := 0; i < 1000000; i++ { //let the coin flipping 1 million times
		//fmt.Println("")
		if headortails() == 1 { //if return value is 1, counter1 +=1
			counter1++
		} else { //else counter2 += 1
			counter2++
		}
	}
	fmt.Println("Counter1: ", counter1, " Counter2: ", counter2) //print result
}

func headortails() int {
	var randomUser1, randomUser2 string //make variables for two random strings

	randomUser1 = GenerateRandomString(12) //generate first random string with 12 character
	//fmt.Println("Input RandomUser1: ", randomUser1)
	randomUser2 = GenerateRandomString(12) //generate second random string with 12 character
	//fmt.Println("Input RandomUser2: ", randomUser2)

	hashUser1 := sha512.Sum512([]byte(randomUser1)) //make sha512 hash of first random string
	//fmt.Printf("User 1 sha512: %x \n", hashUser1)
	hashUser2 := sha512.Sum512([]byte(randomUser2)) //make sha512 hash of second random string
	//fmt.Printf("User 2 sha512: %x \n", hashUser2)

	var countUser1, countUser2 int        //initialise new variables
	for i := 0; i < len(hashUser1); i++ { //for loop from 0 to len of hash
		countUser1 += bits.OnesCount(uint(hashUser1[i])) //count setted bits in hash1
		countUser2 += bits.OnesCount(uint(hashUser2[i])) //count setted bits in hash2
	}
	//fmt.Println("Bits count User1: ", countUser1)
	//fmt.Println("Bits count User2: ", countUser2)
	//fmt.Println("Coinflip is ...")
	if countUser1%2 == 0 && countUser2%2 == 0 { //same procedure as in p2p-beta.py
		//fmt.Println("Head")
		return 0
	} else if countUser1%2 == 1 && countUser2%2 == 1 {
		return 0
	} else {
		return 1
	}
}

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
func GenerateRandomString(s int) string {
	b := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b)		//encode random byte array to base64 encoding
}

func GenerateRandomBytes(n int) []byte {
	b := make([]byte, n)		//new byte array of length n
	_, err := rand.Read(b)		//fill array with random
	if err != nil {				//if error print
		println(err)
	}
	return b					//return array
}
