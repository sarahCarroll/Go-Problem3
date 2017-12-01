package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

func Reflect(input string) string {
	// Split the input on word boundaries.
	boundaries := regexp.MustCompile(`\b`)
	tokens := boundaries.Split(input, -1)

	// List the reflections.
	reflections := [][]string{
		{`I`, `you`},
		{`me`, `you`},
		{`you`, `me`},
		{`my`, `your`},
		{`your`, `my`},
	}

	// Loop through each token, reflecting it if there's a match.
	for i, token := range tokens {
		for _, reflection := range reflections {
			if matched, _ := regexp.MatchString(reflection[0], token); matched {
				tokens[i] = reflection[1]
				break
			}
		}
	}

	// Put the tokens back together.
	return strings.Join(tokens, ``)
}

var responses = []string{
	"I’m not sure what you’re trying to say. Could you explain it to me?",
	"How does that make you feel?",
	"Why do you say that?",
}

func ElizaResponse(input string) string {

	if matched, _ := regexp.MatchString(`(?i).*\bfather\b.*`, input); matched {
		//match string

		return "why dont you tell me more about your father?"
	}

	if matched, _ := regexp.MatchString(`(?i).*\bmother\b.*`, input); matched {
		//match string

		return "Whats your mothers maiden name?"

	}

	if matched, _ := regexp.MatchString(`(?i).*\bbrother\b.*`, input); matched {
		//match string

		return "Does your brother annoy you too?"

	}

	re := regexp.MustCompile(`(?i).*\bi am|i'm|im\b.*([^.?!]*)[.?!]?`)
	if matched := re.MatchString(input); matched {
		return re.ReplaceAllString(input, "How do you know you are $1")
	}

	pe := regexp.MustCompile(`(?i).*\bi dislike\b.*([^.?!]*)[.?!]?`)
	if matched := pe.MatchString(input); matched {
		return pe.ReplaceAllString(input, "Why do you dislike it")
	}

	return responses[rand.Intn(len(responses))]

}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	fmt.Println("People say I look like both my mother and father.")
	fmt.Println(ElizaResponse("People say I look like both my mother and father."))
	fmt.Println()

	fmt.Println("\nFather was a teacher.")
	fmt.Println(ElizaResponse("Father was a teacher."))
	fmt.Println()

	fmt.Println("\nI was my father’s favourite.")
	fmt.Println(ElizaResponse("I was my father’s favourite."))
	fmt.Println()

	fmt.Println("\nI'm looking forward to the weekend.")
	fmt.Println(ElizaResponse("I'm looking forward to the weekend."))
	fmt.Println()

	fmt.Println("\nMy grandfather was French!")
	fmt.Println(ElizaResponse("My grandfather was French!"))
	fmt.Println()

	fmt.Println(" I am happy.")
	fmt.Println(ElizaResponse(" I am happy."))
	fmt.Println()

	fmt.Println("My mother is Nuala")
	fmt.Println(ElizaResponse("My mother is Nuala"))
	fmt.Println()

	fmt.Println("My brother is lazy")
	fmt.Println(ElizaResponse("My brother is lazy"))
	fmt.Println()

	fmt.Println("I dislike all my college work")
	fmt.Println(ElizaResponse("I dislike all my college work"))
	fmt.Println()
}
