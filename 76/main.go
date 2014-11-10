package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/shurcooL/go/vfs_util"
	"github.com/sourcegraph/go-vcs/vcs"
	_ "github.com/sourcegraph/go-vcs/vcs/gitcmd"
	_ "github.com/sourcegraph/go-vcs/vcs/hgcmd"
)

func main() {
	foo()
}

func foo() (interface{}, interface{}, error) {
	rev := ""

	repo, err := vcs.Open("git", "/Users/Dmitri/Dropbox/Work/2013/GoLand/src/github.com/gopherjs/websocket/")
	if err != nil {
		return nil, nil, err
	}

	var commitId vcs.CommitID
	if rev != "" {
		commitId, err = repo.ResolveRevision(rev)
	} else {
		commitId, err = repo.ResolveBranch("master")
	}
	if err != nil {
		return nil, nil, err
	}
	
	fmt.Println(commitId)

	fs, err := repo.FileSystem(commitId)
	if err != nil {
		return nil, nil, err
	}

	{
		walkFn := func(path string, fi os.FileInfo, err error) error {
			if err != nil {
				log.Printf("can't stat file %s: %v\n", path, err)
				return nil
			}
			if strings.HasPrefix(fi.Name(), ".") {
				if fi.IsDir() {
					return filepath.SkipDir
				} else {
					return nil
				}
			}
			fmt.Println(path)
			return nil
		}

		err = vfs_util.Walk(fs, "./", walkFn)
		if err != nil {
			panic(err)
		}
	}

	panic("all good")
}