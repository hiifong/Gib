package main

import (
	"fmt"
	"io"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/memory"
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

	ref, err := r.Head()
	if err != nil {
		return
	}

	logs, err := r.Log(&git.LogOptions{
		From: ref.Hash(),
	})
	if err != nil {
		return
	}
	logs.ForEach(func(commit *object.Commit) error {
		fmt.Println(commit.Author.Name)
		return nil
	})

	fs := memfs.New()
	_, err = git.Clone(memory.NewStorage(), fs, &git.CloneOptions{
		URL: repoPath,
	})
	if err != nil {
		return
	}
	file, err := fs.Open("README.md")
	if err != nil {
		return
	}
	bytes, err := io.ReadAll(file)
	if err != nil {
		return
	}
	fmt.Println(string(bytes))
}
