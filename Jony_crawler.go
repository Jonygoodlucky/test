package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"sync"
)

type Comment struct {
	PostID int    `json:"postId"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

func main() {
	var wg sync.WaitGroup
	emails := make(chan string)

	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d/comments", id)
			resp, err := http.Get(url)
			if err != nil {
				fmt.Printf("Error fetching URL %s: %s", url, err)
				return
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("Error reading response body for URL %s: %s", url, err)
				return
			}

			comments := parseComments(body)
			for _, comment := range comments {
				emails <- comment.Email
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(emails)
	}()

	for email := range emails {
		// do something with the email address, e.g. write to file
		fmt.Println(email)
	}
}

func parseComments(body []byte) []Comment {
	re := regexp.MustCompile(`"email"\s*:\s*"([^"]+)"`)
	matches := re.FindAllSubmatch(body, -1)

	comments := make([]Comment, len(matches))
	for i, match := range matches {
		comment := Comment{Email: string(match[1])}
		comments[i] = comment
	}

	return comments
}
