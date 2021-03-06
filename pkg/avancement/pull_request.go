package avancement

import (
	"fmt"

	"github.com/jenkins-x/go-scm/scm"
	"github.com/rhd-gitops-example/services/pkg/util"
)

// TODO: OptionFuncs for Base and Title?
// TODO: For the Head, should this try and determine whether or not this is a
// fork ("user" of both repoURLs) and if so, simplify the Head?
func makePullRequestInput(fromURL, toURL, branchName string) (*scm.PullRequestInput, error) {
	_, fromRepo, err := util.ExtractUserAndRepo(fromURL)
	if err != nil {
		return nil, err
	}
	fromUser, toRepo, err := util.ExtractUserAndRepo(toURL)
	if err != nil {
		return nil, err
	}
	title := fmt.Sprintf("promotion from %s to %s", fromRepo, toRepo)
	return &scm.PullRequestInput{
		Title: title,
		Head:  fmt.Sprintf("%s:%s", fromUser, branchName),
		Base:  "master",
		Body:  "this is a test body",
	}, nil
}
