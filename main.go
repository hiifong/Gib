package main

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

const repoPath = "./testdata"

func main() {
	fmt.Println("Hello Gib!")
	r, err := git.PlainOpen(repoPath)
	if err != nil {
		fmt.Println(err)
	}

	branches, err := r.Branches()
	if err != nil {
		fmt.Println(err)
	}

	err = branches.ForEach(func(reference *plumbing.Reference) error {
		fmt.Println("branch name: ", reference.Name().Short())
		return nil
	})
	if err != nil {
		return
	}
}
