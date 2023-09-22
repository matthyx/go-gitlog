package gitlog

import "time"

// Hash of commit
type Hash struct {
	Long  string
	Short string
}

// Tree hash of commit
type Tree struct {
	Long  string
	Short string
}

// Author of commit
type Author struct {
	Name  string
	Email string
	Date  time.Time
}

// Tag of commit
type Tag struct {
	Name string
	Date time.Time
}

// Committer of commit
type Committer struct {
	Name  string
	Email string
	Date  time.Time
}

// Commit data
type Commit struct {
	Hash      *Hash
	Tree      *Tree
	Author    *Author
	Committer *Committer
	Tag       *Tag
	Subject   string
	Body      string
	Files     []string
	Parents   []string
}
