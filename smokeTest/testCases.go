package main

import (
	"encoding/json"
	"fmt"
)

type testCase struct {
	Name     string
	Method   string
	Url      string
	Headers  map[string]string
	Body     string
	Validate func(int, []byte) error // status code, body
}

// Create user 1
// Create user 2
// Create user 3
// User 1 creates content 1.1
// User 3 follows user 1
// User 2 creates content 2.1
// User 2 creates content 2.2
// User 2 creates content 2.3
// User 3 follows user 2
// User 1 creates content 1.2
// User 2 creates content 2.4
// User 1 creates content 1.3
// User 2 creates content 2.5
// User 1 creates content 1.4
// User 2 creates content 2.6
// User 1 creates content 1.5
// User 2 creates content 2.7
// User 3 ask feed --> len = 12, order: [2.7, 1.5, 2.6, 1.4, 2.5, 1.3, 2.4, 1.2, 2.3, 2.2, 2.1, 1.1]

var testCases = []testCase{
	{
		Name:   "create user 1",
		Method: "POST",
		Url:    "http://localhost:8090/users/create",
		Body:   `{"name":"user1","password":"password1"}`,
		Validate: func(sc int, body []byte) error {
			if sc != 201 {
				return fmt.Errorf("status code should be 201, not %d", sc)
			}
			toDelete(users, body)
			return nil
		},
	},
	{
		Name:   "create user 2",
		Method: "POST",
		Url:    "http://localhost:8090/users/create",
		Body:   `{"name":"user2","password":"password2"}`,
		Validate: func(sc int, body []byte) error {
			if sc != 201 {
				return fmt.Errorf("status code should be 201, not %d", sc)
			}
			toDelete(users, body)
			return nil
		},
	},
	{
		Name:   "create user 3",
		Method: "POST",
		Url:    "http://localhost:8090/users/create",
		Body:   `{"name":"user3","password":"password3"}`,
		Validate: func(sc int, body []byte) error {
			if sc != 201 {
				return fmt.Errorf("status code should be 201, not %d", sc)
			}
			toDelete(users, body)
			return nil
		},
	},
	{
		Name:   "user 1 creates content 1.1",
		Method: "POST",
		Url:    "http://localhost:8091/contents/create",
		Headers: map[string]string{
			"X-Tuala-Name":     "user1",
			"X-Tuala-Password": "password1",
		},
		Body: `{"text":"1.1"}`,
		Validate: func(sc int, body []byte) error {
			if sc != 201 {
				return fmt.Errorf("status code should be 201, not %d", sc)
			}
			toDelete(content, body)
			return nil
		},
	},
	{
		Name:   "user 3 follows user 1",
		Method: "POST",
		Url:    "http://localhost:8090/users/follow/user1",
		Headers: map[string]string{
			"X-Tuala-Name":     "user3",
			"X-Tuala-Password": "password3",
		},
		Validate: func(sc int, body []byte) error {
			if sc != 200 {
				return fmt.Errorf("status code should be 200, not %d", sc)
			}
			return nil
		},
	},
	{
		Name:   "user 2 creates content 2.1",
		Method: "POST",
		Url:    "http://localhost:8091/contents/create",
		Headers: map[string]string{
			"X-Tuala-Name":     "user2",
			"X-Tuala-Password": "password2",
		},
		Body: `{"text":"2.1"}`,
		Validate: func(sc int, body []byte) error {
			if sc != 201 {
				return fmt.Errorf("status code should be 201, not %d", sc)
			}
			toDelete(content, body)
			return nil
		},
	},
	{
		Name:   "user 2 creates content 2.2",
		Method: "POST",
		Url:    "http://localhost:8091/contents/create",
		Headers: map[string]string{
			"X-Tuala-Name":     "user2",
			"X-Tuala-Password": "password2",
		},
		Body: `{"text":"2.2"}`,
		Validate: func(sc int, body []byte) error {
			if sc != 201 {
				return fmt.Errorf("status code should be 201, not %d", sc)
			}
			toDelete(content, body)
			return nil
		},
	},
	{
		Name:   "user 2 creates content 2.3",
		Method: "POST",
		Url:    "http://localhost:8091/contents/create",
		Headers: map[string]string{
			"X-Tuala-Name":     "user2",
			"X-Tuala-Password": "password2",
		},
		Body: `{"text":"2.3"}`,
		Validate: func(sc int, body []byte) error {
			if sc != 201 {
				return fmt.Errorf("status code should be 201, not %d", sc)
			}
			toDelete(content, body)
			return nil
		},
	},
	{
		Name:   "user 3 follows user 2",
		Method: "POST",
		Url:    "http://localhost:8090/users/follow/user2",
		Headers: map[string]string{
			"X-Tuala-Name":     "user3",
			"X-Tuala-Password": "password3",
		},
		Validate: func(sc int, body []byte) error {
			if sc != 200 {
				return fmt.Errorf("status code should be 200, not %d", sc)
			}
			return nil
		},
	},
	{
		Name:   "user 1 creates content 1.2",
		Method: "POST",
		Url:    "http://localhost:8091/contents/create",
		Headers: map[string]string{
			"X-Tuala-Name":     "user1",
			"X-Tuala-Password": "password1",
		},
		Body: `{"text":"1.2"}`,
		Validate: func(sc int, body []byte) error {
			if sc != 201 {
				return fmt.Errorf("status code should be 201, not %d", sc)
			}
			toDelete(content, body)
			return nil
		},
	},
	{
		Name:   "user 2 creates content 2.4",
		Method: "POST",
		Url:    "http://localhost:8091/contents/create",
		Headers: map[string]string{
			"X-Tuala-Name":     "user2",
			"X-Tuala-Password": "password2",
		},
		Body: `{"text":"2.4"}`,
		Validate: func(sc int, body []byte) error {
			if sc != 201 {
				return fmt.Errorf("status code should be 201, not %d", sc)
			}
			toDelete(content, body)
			return nil
		},
	},
	{
		Name:   "user 1 creates content 1.3",
		Method: "POST",
		Url:    "http://localhost:8091/contents/create",
		Headers: map[string]string{
			"X-Tuala-Name":     "user1",
			"X-Tuala-Password": "password1",
		},
		Body: `{"text":"1.3"}`,
		Validate: func(sc int, body []byte) error {
			if sc != 201 {
				return fmt.Errorf("status code should be 201, not %d", sc)
			}
			toDelete(content, body)
			return nil
		},
	},
	{
		Name:   "user 2 creates content 2.5",
		Method: "POST",
		Url:    "http://localhost:8091/contents/create",
		Headers: map[string]string{
			"X-Tuala-Name":     "user2",
			"X-Tuala-Password": "password2",
		},
		Body: `{"text":"2.5"}`,
		Validate: func(sc int, body []byte) error {
			if sc != 201 {
				return fmt.Errorf("status code should be 201, not %d", sc)
			}
			toDelete(content, body)
			return nil
		},
	},
	{
		Name:   "user 1 creates content 1.4",
		Method: "POST",
		Url:    "http://localhost:8091/contents/create",
		Headers: map[string]string{
			"X-Tuala-Name":     "user1",
			"X-Tuala-Password": "password1",
		},
		Body: `{"text":"1.4"}`,
		Validate: func(sc int, body []byte) error {
			if sc != 201 {
				return fmt.Errorf("status code should be 201, not %d", sc)
			}
			toDelete(content, body)
			return nil
		},
	},
	{
		Name:   "user 2 creates content 2.6",
		Method: "POST",
		Url:    "http://localhost:8091/contents/create",
		Headers: map[string]string{
			"X-Tuala-Name":     "user2",
			"X-Tuala-Password": "password2",
		},
		Body: `{"text":"2.6"}`,
		Validate: func(sc int, body []byte) error {
			if sc != 201 {
				return fmt.Errorf("status code should be 201, not %d", sc)
			}
			toDelete(content, body)
			return nil
		},
	},
	{
		Name:   "user 1 creates content 1.5",
		Method: "POST",
		Url:    "http://localhost:8091/contents/create",
		Headers: map[string]string{
			"X-Tuala-Name":     "user1",
			"X-Tuala-Password": "password1",
		},
		Body: `{"text":"1.5"}`,
		Validate: func(sc int, body []byte) error {
			if sc != 201 {
				return fmt.Errorf("status code should be 201, not %d", sc)
			}
			toDelete(content, body)
			return nil
		},
	},
	{
		Name:   "user 2 creates content 2.7",
		Method: "POST",
		Url:    "http://localhost:8091/contents/create",
		Headers: map[string]string{
			"X-Tuala-Name":     "user2",
			"X-Tuala-Password": "password2",
		},
		Body: `{"text":"2.7"}`,
		Validate: func(sc int, body []byte) error {
			if sc != 201 {
				return fmt.Errorf("status code should be 201, not %d", sc)
			}
			toDelete(content, body)
			return nil
		},
	},
	{
		Name:   "user 3 ask feed",
		Method: "GET",
		Url:    "http://localhost:8091/feed/recent",
		Headers: map[string]string{
			"X-Tuala-Name":     "user3",
			"X-Tuala-Password": "password3",
		},
		Validate: func(sc int, body []byte) error {
			if sc != 200 {
				return fmt.Errorf("status code should be 200, not %d", sc)
			}
			type tweet struct {
				Text string `json:"text"`
			}
			var content []tweet
			err := json.Unmarshal(body, &content)
			if err != nil {
				return err
			}
			var expectedText = []string{
				"2.7", "1.5", "2.6",
				"1.4", "2.5", "1.3",
				"2.4", "1.2", "2.3",
				"2.2", "2.1", "1.1",
			}
			if len(content) != len(expectedText) {
				return fmt.Errorf(
					"read content length of %d when expected %d",
					len(content), len(expectedText),
				)
			}
			for i := 0; i < len(content); i++ {
				if content[i].Text != expectedText[i] {
					return fmt.Errorf(
						"invalid text %s for %d position; should be invalid sorting",
						content[i].Text, i,
					)
				}
			}
			return nil
		},
	},
}

// user 3 ask feed --> len = 12, order: [2.7, 1.5, 2.6, 1.4, 2.5, 1.3, 2.4, 1.2, 2.3, 2.2, 2.1, 1.1]
