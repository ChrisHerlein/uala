package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := runTestCases()
	if err != nil {
		fmt.Println("TEST FAILED:", err)
	}
	cleanup()
}

func runTestCases() error {
	for i := 0; i < len(testCases); i++ {
		time.Sleep(1 * time.Second) // fake user interaction
		tc := testCases[i]
		fmt.Println("Running:", tc.Name)
		if i == len(testCases)-1 {
			fmt.Println("await for cache creating")
			time.Sleep(10 * time.Second)
		}
		req, err := http.NewRequest(tc.Method, tc.Url, bytes.NewBufferString(tc.Body))
		if err != nil {
			return err
		}

		req.Header.Add("Content-Type", "application/json")
		for k, v := range tc.Headers {
			req.Header.Add(k, v)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		err = tc.Validate(res.StatusCode, body)
		if err != nil {
			return err
		}
	}
	fmt.Println("All cases passed!")
	return nil
}

func cleanup() {
	connStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		"localhost", "uala", "challenge", "tuala", "5432",
	)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	defer func() {
		rawdb, err := db.DB()
		if err != nil {
			return
		}
		rawdb.Close()
	}()

	for i := 0; i < len(toDeleteResources); i++ {
		for table, id := range toDeleteResources[i] {
			q := fmt.Sprintf("DELETE FROM %s WHERE id = %+v;", table, id)
			db.Exec(q)
		}
	}
}
