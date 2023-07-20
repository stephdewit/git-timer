package main

import (
	"log"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func main() {
	repo, err := git.PlainOpen(".")
	if err != nil {
		log.Fatal("can't open repository: %w", err)
	}

	commits, err := repo.Log(&git.LogOptions{})
	if err != nil {
		log.Fatal("can't get repository log: %w", err)
	}

	err = commits.ForEach(func(c *object.Commit) error {
		log.Print(c)
		return nil
	})

	if err != nil {
		log.Fatal("can't iterate over commits: %w", err)
	}
}
