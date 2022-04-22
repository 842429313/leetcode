package main

// func isaeiou(word byte) bool {
// 	if (word == 'a') ||
// 		(word == 'e') ||
// 		(word == 'i') ||
// 		(word == 'o') ||
// 		(word == 'u') ||
// 		(word == 'A') ||
// 		(word == 'E') ||
// 		(word == 'I') ||
// 		(word == 'O') ||
// 		(word == 'U') {
// 		return true
// 	} else {
// 		return false
// 	}
// }

// func toGoatLatin(sentence string) string {
// 	var result string = ""
// 	var byte_word string = " "
// 	var word_num int = 0
// 	for i := 0; i < len(sentence); i++ {
// 		if sentence[i] == ' ' {
// 			if byte_word != " " {
// 				result = result + byte_word
// 				byte_word = " "
// 			}
// 			result = result + "ma"
// 			for j := 0; j <= word_num; j++ {
// 				result = result + "a"
// 			}
// 			word_num = word_num + 1
// 		}
// 		if i != 0 {
// 			if sentence[i-1] == ' ' && !isaeiou((sentence[i])) {
// 				byte_word = string(sentence[i])
// 			} else {
// 				result = result + string(sentence[i])
// 			}
// 		} else {
// 			if !isaeiou((sentence[i])) {
// 				byte_word = string(sentence[i])
// 			} else {
// 				result = result + string(sentence[i])
// 			}
// 		}
// 	}
// 	if byte_word != " " {
// 		result = result + byte_word
// 		byte_word = " "
// 	}
// 	result = result + "ma"
// 	for j := 0; j <= word_num; j++ {
// 		result = result + "a"
// 	}
// 	return result
// }
// func main() {
// 	var temp string = "I speak Goat Latin"
// 	fmt.Println(toGoatLatin(temp))
// }
