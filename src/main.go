package main

import (
	"fmt"
    "os"
	"path/filepath"
	"regexp"
    // external
    "github.com/google/uuid"
	"github.com/go-git/go-git/v5"
	//gitplumbing "github.com/go-git/go-git/v5/plumbing"
	//gitobject "github.com/go-git/go-git/v5/plumbing/object"
	githttp "github.com/go-git/go-git/v5/plumbing/transport/http"
)

type LocalRepo struct {
    Path string
    Repo *git.Repository
}

func NewLocalRepo(remote, user, token, directory string) (*LocalRepo, error) {
    
    repo_name := regexp.MustCompile(`[a-zA-Z0-9_-]+\.git$`).FindString(remote)
    if len(repo_name) < 1 {
        return nil, fmt.Errorf("Invalid remote %s", remote)
    }

    abs_path_to_directory, err := filepath.Abs(directory)
    if err != nil {
        return nil, fmt.Errorf("Invalid directory %s, getting absolute path failed with error %s", directory, err)
    }

    path := filepath.Join(abs_path_to_directory, uuid.New().String())

    err = os.MkdirAll(path, 0755)
	if err != nil {
		return nil, fmt.Errorf("Creating directory %s failed with error %s", directory, err)
	}

    repo, err := git.PlainClone(path, false,
        &git.CloneOptions{
            URL: remote,
            NoCheckout: true,
            // Token can be supplied instead of a password
            Auth: &githttp.BasicAuth{
                Username: user,
                Password: token,
            },
    })
    if err != nil {
        return nil, fmt.Errorf("Checking out %s from remote %s failed with error: %s", repo_name, remote, err)
    }

    return &LocalRepo{
        Repo: repo,
        Path: path,
    }, nil
}

func (local *LocalRepo) CleanFiles() error {
    // STUB
    return nil
}

func (local *LocalRepo) Checkout() error {
    // STUB
    return nil
}

func main() {
    
}
