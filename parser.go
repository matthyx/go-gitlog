package gitlog

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type parser struct{}

const nl = "\n"

func (p *parser) parse(str string) ([]*Commit, error) {
	lines := strings.Split(str, separator)
	if len(lines) < 1 {
		return nil, errors.New("wip")
	}

	lines = lines[1:]
	commits := make([]*Commit, len(lines))

	for i, line := range lines {
		commits[i] = p.parseCommit(line)
	}

	return commits, nil
}

func (p *parser) parseCommit(str string) *Commit {
	segments := strings.Split(str, delimiter)
	commit := &Commit{}

	for _, segment := range segments {
		endFieldType := strings.Index(segment, ":")
		fieldType := segment[0:endFieldType]
		content := segment[endFieldType+1:]

		switch fieldType {
		case hashField:
			commit.Hash = p.parseHash(content)
		case treeField:
			commit.Tree = p.parseTree(content)
		case authorField:
			commit.Author = p.parseAuthor(content)
		case committerField:
			commit.Committer = p.parseCommitter(content)
		case tagField:
			commit.Tag = p.parseTag(content)
		case subjectField:
			commit.Subject = p.parseSubject(content)
		case bodyField:
			commit.Body, commit.Files = p.parseBody(content)
		case parentField:
			commit.Parents = p.parseParents(content)
		}
	}

	commit.Tag.Date = commit.Author.Date

	return commit
}

func (p *parser) parseHash(str string) *Hash {
	parts := strings.Split(str, " ")
	return &Hash{
		Long:  parts[0],
		Short: parts[1],
	}
}

func (p *parser) parseTree(str string) *Tree {
	hash := p.parseHash(str)

	return &Tree{
		hash.Long,
		hash.Short,
	}
}

func (p *parser) parseAuthor(s string) *Author {
	beginEmail := strings.Index(s, "<")
	endEmail := strings.Index(s, ">")

	beginDate := strings.LastIndex(s, "[")
	endDate := strings.LastIndex(s, "]")

	name := s[:beginEmail]
	email := s[beginEmail+1 : endEmail]
	timestamp, _ := strconv.Atoi(s[beginDate+1 : endDate])

	return &Author{
		name,
		email,
		time.Unix(int64(timestamp), 0),
	}
}

func (p *parser) parseCommitter(str string) *Committer {
	author := p.parseAuthor(str)

	return &Committer{
		author.Name,
		author.Email,
		author.Date,
	}
}

var tagRegex = regexp.MustCompile("tag:\\s([\\w\\.\\-_/]+)")

func (p *parser) parseTag(str string) *Tag {
	res := tagRegex.FindAllStringSubmatch(str, -1)
	tag := &Tag{
		Name: "",
	}

	if len(res) > 0 {
		tag.Name = res[0][1]
	}

	return tag
}

func (*parser) convNewline(str string) string {
	return strings.NewReplacer(
		"\r\n", nl,
		"\r", nl,
		"\n", nl,
	).Replace(str)
}

func (p *parser) parseSubject(str string) string {
	return strings.TrimSpace(p.convNewline(str))
}

func trimAll(s string) string {
	s = strings.TrimSpace(s)
	s = strings.Trim(s, "\"")
	s = strings.TrimSpace(s)
	s = strings.Trim(s, "\"")
	s = strings.TrimSpace(s)
	return s
}

func (p *parser) parseBody(str string) (string, []string) {
	split := strings.SplitN(p.convNewline(str), "\n\n", 2)
	body := trimAll(split[0])
	if len(split) < 2 || len(split[1]) == 0 {
		return body, nil
	}
	files := trimAll(split[1])
	return body, strings.Split(files, "\n")
}

func (p *parser) parseParents(s string) []string {
	return strings.Split(s, " ")
}
