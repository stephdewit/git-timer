package main

import (
	"log"
	"sort"
	"time"

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

	times := make([]time.Time, 0, 10)

	err = commits.ForEach(func(c *object.Commit) error {
		times = append(times, c.Author.When)

		if c.Author.When != c.Committer.When {
			times = append(times, c.Committer.When)
		}

		return nil
	})

	if err != nil {
		log.Fatal("can't iterate over commits: %w", err)
	}

	if len(times) < 2 {
		log.Fatal("not enough commits")
	}

	sort.Slice(times, func(i, j int) bool {
		return times[i].Before(times[j])
	})

	log.Print(times)
}
