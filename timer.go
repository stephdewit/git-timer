package main

import (
	"fmt"
	"log"
	"sort"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

var delay time.Duration = time.Minute * 60

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

	sessions := make([]result, 0, 10)
	var current result
	current.from = times[0]

	for i := 1; i < len(times); i++ {
		from := times[i-1]
		to := times[i]

		if to.Sub(from) < delay {
			current.duration = to.Sub(current.from)
		} else {
			if current.duration > 0 {
				sessions = append(sessions, current)
				current = result{}
			} else {
				log.Print("Lonely commit")
			}
			current.from = to
		}
	}

	if current.duration > 0 {
		sessions = append(sessions, current)
	}

	for _, session := range sessions {
		log.Print(session)
	}
}

type result struct {
	from     time.Time
	duration time.Duration
}

func (r result) String() string {
	return fmt.Sprintf("%s: %v", r.from.Format(time.RFC1123Z), r.duration)
}
