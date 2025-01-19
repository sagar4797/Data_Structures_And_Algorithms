// Problem link: https://leetcode.com/discuss/interview-question/4944166/Microsoft-Software-Engineer-I-OA/

// Prepare a notification of the given message which will be displayed on a mobile device. The message is made of words separated by single spaces. The length of the notification is limited to K characters. If the message is too long to be displayed fully, some words from the end of the message should be cropped, keeping in mind that:
// • the notification should be as long as possible;
// • only whole wirds can be cropped;
// • if any words are cropped, the notification should end with " ..."; the dots should be separated from the last word by a single space;
// • the notification may not exceed the K-character limit, including the dots.
// If all the words need to be cropped, the notification is"..." (three dots without a space).
// For example, for message = "And now here is my secret" and K = 15,
// the notification is "and now". Note that:
// • the notification "And ..." would be incorrect, because there is a longer correct notification;
// • the notification "And now her ..." would be incorrect, because the original message is cropped through the middle of a word;
// • the notification "And now ... " would be incorrect, because it ends with a space;
// • the notification "And now here..." would be incorrect, because it ends with a space;
// • the notificatign "And now here..." would be incorrect, because there is no space before the three dots;
// • the notification "And now here ..." would be incorrect, because it exceeds the 15-character limit.
// Write a function:
// class Solution & public String solution(String message, int si
// that, given a string message and an integer K, returns the notification t display, which has no more than K characters, as described above.
// Examples:
// For message = "And now here is my secret" and K = 15, the
// function should return "And now...", as explained above.
// For message = "There is an animal with four legs" and K = 15,
// the function should return "There is an ...".
// For message = "super dog" and K - 4, the function should return
// For message = "how are you" and K = 20, the function should return
// "how are you".
// Assume that:
// • Kis an integer within the range |3.500;
// • the length of string message is within the range [1.500;
// • string message is made of English letters (a-z, 'A-2) and spaces ():
// • message does not contain spaces at the beginning or at the end;
// • words are separated by a single space (there are never two or more consecutive spaces).


package main

import (
	"fmt"
	"strings"
)

func prepareNotification(message string, k int) string {
	words := strings.Split(message, " ")
	if len(message) <= k {
		return message
	}

	ellipsis := "..."
	ellipsisLen := len(ellipsis) + 1 // including the space before ellipsis
	remainingLen := k - ellipsisLen

	var result []string
	currentLen := 0

	for _, word := range words {
		wordLen := len(word)
		if currentLen+wordLen > remainingLen {
			break
		}
		result = append(result, word)
		currentLen += wordLen + 1 // including the space
	}

	if len(result) == 0 {
		return ellipsis
	}

	return strings.Join(result, " ") + " " + ellipsis
}

func main() {
	message1 := "Prepare a notification of the given message which will be displayed on a mobile device"
	k1 := 50
	fmt.Println(prepareNotification(message1, k1)) // Output: "Prepare a notification of the given message which will..."

	message2 := "The quick brown fox jumps over the lazy dog"
	k2 := 20
	fmt.Println(prepareNotification(message2, k2)) // Output: "The quick brown..."

	message3 := "Hello world"
	k3 := 15
	fmt.Println(prepareNotification(message3, k3)) // Output: "Hello world"

	message4 := "This is a test message"
	k4 := 10
	fmt.Println(prepareNotification(message4, k4)) // Output: "This..."

	message5 := "Short"
	k5 := 10
	fmt.Println(prepareNotification(message5, k5)) // Output: "Short"
}
