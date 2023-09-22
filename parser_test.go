package gitlog

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	assert := assert.New(t)

	commitLog := `@@__GIT_LOG_SEPARATOR__@@HASH:51064a83516c60fdffd99a7d605d168298d91464 51064a8@@__GIT_LOG_DELIMITER__@@TREE:4b825dc642cb6eb9a060e54bf8d69288fbee4904 4b825dc@@__GIT_LOG_DELIMITER__@@AUTHOR:tsuyoshiwada<mail@example.com>[1517138361]@@__GIT_LOG_DELIMITER__@@COMMITTER:tsuyoshiwada<mail@example.com>[1517138361]@@__GIT_LOG_DELIMITER__@@TAG:tag: 1.2.0-beta.1, origin/refactor/branch@@__GIT_LOG_DELIMITER__@@SUBJECT:chore(*): Has commit body@@__GIT_LOG_DELIMITER__@@BODY:This is body comment
multiline
foo
bar

@@__GIT_LOG_SEPARATOR__@@HASH:6dccb5c65f984ec8857243017f506008683342c2 6dccb5c@@__GIT_LOG_DELIMITER__@@TREE:4b825dc642cb6eb9a060e54bf8d69288fbee4904 4b825dc@@__GIT_LOG_DELIMITER__@@AUTHOR:tsuyoshiwada<mail@example.com>[1517134427]@@__GIT_LOG_DELIMITER__@@COMMITTER:tsuyoshiwada<mail@example.com>[1517134427]@@__GIT_LOG_DELIMITER__@@TAG:@@__GIT_LOG_DELIMITER__@@SUBJECT:docs(readme): Test commit@@__GIT_LOG_DELIMITER__@@BODY:
@@__GIT_LOG_SEPARATOR__@@HASH:a6cca30eb0b0685c45c1da53c3629913895c1fd0 a6cca30e@@__GIT_LOG_DELIMITER__@@PARENT:e611cec238f51147c5dc903996f0ae9f1d7bdd7a c74c5f19709948665ea030b0a84fcaa5e0c4686f@@__GIT_LOG_DELIMITER__@@TREE:94f4b58ccbdbe231f8ea23b59278aae0e224dd61 94f4b58c@@__GIT_LOG_DELIMITER__@@AUTHOR:David Wertenteil<dwertent@armosec.io>[1694423357]@@__GIT_LOG_DELIMITER__@@COMMITTER:GitHub<noreply@github.com>[1694423357]@@__GIT_LOG_DELIMITER__@@TAG:@@__GIT_LOG_DELIMITER__@@SUBJECT:Merge pull request #1381 from XDRAGON2002/issue_1380@@__GIT_LOG_DELIMITER__@@BODY:feat: darken table borders
@@__GIT_LOG_SEPARATOR__@@HASH:c74c5f19709948665ea030b0a84fcaa5e0c4686f c74c5f19@@__GIT_LOG_DELIMITER__@@PARENT:e611cec238f51147c5dc903996f0ae9f1d7bdd7a@@__GIT_LOG_DELIMITER__@@TREE:94f4b58ccbdbe231f8ea23b59278aae0e224dd61 94f4b58c@@__GIT_LOG_DELIMITER__@@AUTHOR:DRAGON<anantvijay3@gmail.com>[1694358722]@@__GIT_LOG_DELIMITER__@@COMMITTER:DRAGON<anantvijay3@gmail.com>[1694362695]@@__GIT_LOG_DELIMITER__@@TAG:@@__GIT_LOG_DELIMITER__@@SUBJECT:feat: darken table borders@@__GIT_LOG_DELIMITER__@@BODY:Signed-off-by: DRAGON <anantvijay3@gmail.com>


core/core/list.go
core/pkg/resultshandling/printer/v2/prettyprinter/tableprinter/configurationprinter/categorytable.go
core/pkg/resultshandling/printer/v2/prettyprinter/tableprinter/configurationprinter/frameworkscan.go
@@__GIT_LOG_SEPARATOR__@@HASH:806512fe97c9c3397b7ed30c0b4076032112f697 806512f@@__GIT_LOG_DELIMITER__@@TREE:4b825dc642cb6eb9a060e54bf8d69288fbee4904 4b825dc@@__GIT_LOG_DELIMITER__@@AUTHOR:tsuyoshiwada<mail@example.com>[1517122160]@@__GIT_LOG_DELIMITER__@@COMMITTER:tsuyoshiwada<mail@example.com>[1517122160]@@__GIT_LOG_DELIMITER__@@TAG:tag: v0.2.1@@__GIT_LOG_DELIMITER__@@SUBJECT:chore(*): Initial commit@@__GIT_LOG_DELIMITER__@@BODY:`

	table := []*Commit{
		{
			Hash: &Hash{
				Long:  "51064a83516c60fdffd99a7d605d168298d91464",
				Short: "51064a8",
			},
			Tree: &Tree{
				Long:  "4b825dc642cb6eb9a060e54bf8d69288fbee4904",
				Short: "4b825dc",
			},
			Author: &Author{
				Name:  "tsuyoshiwada",
				Email: "mail@example.com",
				Date:  time.Unix(1517138361, 0),
			},
			Committer: &Committer{
				Name:  "tsuyoshiwada",
				Email: "mail@example.com",
				Date:  time.Unix(1517138361, 0),
			},
			Tag: &Tag{
				Name: "1.2.0-beta.1",
				Date: time.Unix(1517138361, 0),
			},
			Subject: "chore(*): Has commit body",
			Body: `This is body comment
multiline
foo
bar`,
		},
		{
			Hash: &Hash{
				Long:  "6dccb5c65f984ec8857243017f506008683342c2",
				Short: "6dccb5c",
			},
			Tree: &Tree{
				Long:  "4b825dc642cb6eb9a060e54bf8d69288fbee4904",
				Short: "4b825dc",
			},
			Author: &Author{
				Name:  "tsuyoshiwada",
				Email: "mail@example.com",
				Date:  time.Unix(1517134427, 0),
			},
			Committer: &Committer{
				Name:  "tsuyoshiwada",
				Email: "mail@example.com",
				Date:  time.Unix(1517134427, 0),
			},
			Tag: &Tag{
				Name: "",
				Date: time.Unix(1517134427, 0),
			},
			Subject: "docs(readme): Test commit",
			Body:    "",
		},
		{
			Hash: &Hash{
				Long:  "a6cca30eb0b0685c45c1da53c3629913895c1fd0",
				Short: "a6cca30e",
			},
			Tree: &Tree{
				Long:  "94f4b58ccbdbe231f8ea23b59278aae0e224dd61",
				Short: "94f4b58c",
			},
			Author: &Author{
				Name:  "David Wertenteil",
				Email: "dwertent@armosec.io",
				Date:  time.Unix(1694423357, 0),
			},
			Committer: &Committer{
				Name:  "GitHub",
				Email: "noreply@github.com",
				Date:  time.Unix(1694423357, 0),
			},
			Tag: &Tag{
				Name: "",
				Date: time.Unix(1694423357, 0),
			},
			Subject: "Merge pull request #1381 from XDRAGON2002/issue_1380",
			Body:    "feat: darken table borders",
			Parents: []string{"e611cec238f51147c5dc903996f0ae9f1d7bdd7a", "c74c5f19709948665ea030b0a84fcaa5e0c4686f"},
		},
		{
			Hash: &Hash{
				Long:  "c74c5f19709948665ea030b0a84fcaa5e0c4686f",
				Short: "c74c5f19",
			},
			Tree: &Tree{
				Long:  "94f4b58ccbdbe231f8ea23b59278aae0e224dd61",
				Short: "94f4b58c",
			},
			Author: &Author{
				Name:  "DRAGON",
				Email: "anantvijay3@gmail.com",
				Date:  time.Unix(1694358722, 0),
			},
			Committer: &Committer{
				Name:  "DRAGON",
				Email: "anantvijay3@gmail.com",
				Date:  time.Unix(1694362695, 0),
			},
			Tag: &Tag{
				Name: "",
				Date: time.Unix(1694358722, 0),
			},
			Subject: "feat: darken table borders",
			Body:    "Signed-off-by: DRAGON <anantvijay3@gmail.com>",
			Files: []string{
				"core/core/list.go",
				"core/pkg/resultshandling/printer/v2/prettyprinter/tableprinter/configurationprinter/categorytable.go",
				"core/pkg/resultshandling/printer/v2/prettyprinter/tableprinter/configurationprinter/frameworkscan.go",
			},
			Parents: []string{"e611cec238f51147c5dc903996f0ae9f1d7bdd7a"},
		},
		{
			Hash: &Hash{
				Long:  "806512fe97c9c3397b7ed30c0b4076032112f697",
				Short: "806512f",
			},
			Tree: &Tree{
				Long:  "4b825dc642cb6eb9a060e54bf8d69288fbee4904",
				Short: "4b825dc",
			},
			Author: &Author{
				Name:  "tsuyoshiwada",
				Email: "mail@example.com",
				Date:  time.Unix(1517122160, 0),
			},
			Committer: &Committer{
				Name:  "tsuyoshiwada",
				Email: "mail@example.com",
				Date:  time.Unix(1517122160, 0),
			},
			Tag: &Tag{
				Name: "v0.2.1",
				Date: time.Unix(1517122160, 0),
			},
			Subject: "chore(*): Initial commit",
			Body:    "",
		},
	}

	parser := &parser{}
	commits, err := parser.parse(commitLog)

	assert.Nil(err)
	assert.Equal(table, commits)
}
