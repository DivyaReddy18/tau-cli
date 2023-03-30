package repositoryLib

import (
	git "github.com/taubyte/go-simple-git"
	"github.com/taubyte/tau/singletons/config"
)

func (info *Info) Push(project config.Project, message, url string) (*git.Repository, error) {
	repo, err := info.Open(project, url)
	if err != nil {
		return nil, err
	}

	err = repo.Commit(message, ".")
	if err != nil {
		return nil, err
	}

	err = repo.Push()
	if err != nil {
		return nil, err
	}

	return repo, nil
}
