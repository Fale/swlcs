package pr

import (
	"fmt"

	"github.com/google/go-github/github"
)

func (pr PR) Execute() error {
	ref, err := pr.createBranch()
	if err != nil {
		return fmt.Errorf("unable to create the commit branch: %s\n", err)
	}
	tree, err := pr.getTree(ref)
	if err != nil {
		return fmt.Errorf("unable to create the tree based on the provided files: %s\n", err)
	}

	if err := pr.pushCommit(ref, tree); err != nil {
		return fmt.Errorf("unable to create the commit: %s\n", err)
	}

	if err := pr.createPR(); err != nil {
		return fmt.Errorf("error while creating the pull request: %s", err)
	}
	return nil
}

func (pr PR) createBranch() (ref *github.Reference, err error) {
	var baseRef *github.Reference
	if baseRef, _, err = pr.repository.GitHubClient.Git.GetRef(pr.ctx, pr.repository.Owner, pr.repository.Name, fmt.Sprintf("refs/heads/%s", pr.repository.Branch)); err != nil {
		return nil, err
	}
	newRef := &github.Reference{Ref: github.String(fmt.Sprintf("refs/heads/%s", pr.commitBranch)), Object: &github.GitObject{SHA: baseRef.Object.SHA}}
	ref, _, err = pr.repository.GitHubClient.Git.CreateRef(pr.ctx, pr.repository.Owner, pr.repository.Name, newRef)
	return ref, err
}

func (pr PR) getTree(ref *github.Reference) (tree *github.Tree, err error) {
	entries := []github.TreeEntry{
		{
			Path:    &pr.comment.FileName,
			Type:    github.String("blob"),
			Content: &pr.comment.Content,
			Mode:    github.String("100644"),
		},
	}
	tree, _, err = pr.repository.GitHubClient.Git.CreateTree(pr.ctx, pr.repository.Owner, pr.repository.Name, *ref.Object.SHA, entries)
	return tree, err
}

func (pr PR) pushCommit(ref *github.Reference, tree *github.Tree) (err error) {
	// Get the parent commit to attach the commit to.
	parent, _, err := pr.repository.GitHubClient.Repositories.GetCommit(pr.ctx, pr.repository.Owner, pr.repository.Name, *ref.Object.SHA)
	if err != nil {
		return err
	}
	// This is not always populated, but is needed.
	parent.Commit.SHA = parent.SHA

	// Create the commit using the tree.
	commit := &github.Commit{
		Author:  &github.CommitAuthor{Date: &pr.comment.Time, Name: &pr.comment.AuthorName, Email: &pr.comment.AuthorEmail},
		Message: github.String(fmt.Sprintf("Add %s comment to post %s", pr.comment.AuthorName, pr.comment.Resource)),
		Tree:    tree,
		Parents: []github.Commit{*parent.Commit},
	}
	newCommit, _, err := pr.repository.GitHubClient.Git.CreateCommit(pr.ctx, pr.repository.Owner, pr.repository.Name, commit)
	if err != nil {
		return err
	}

	// Attach the commit to the master branch.
	ref.Object.SHA = newCommit.SHA
	_, _, err = pr.repository.GitHubClient.Git.UpdateRef(pr.ctx, pr.repository.Owner, pr.repository.Name, ref, false)
	return err
}

func (pr PR) createPR() (err error) {
	newPR := &github.NewPullRequest{
		Title:               &pr.subject,
		Base:                &pr.repository.Branch,
		Head:                &pr.commitBranch,
		Body:                github.String(fmt.Sprintf("Add %s comment to post %s", pr.comment.AuthorName, pr.comment.Resource)),
		MaintainerCanModify: github.Bool(true),
	}

	p, _, err := pr.repository.GitHubClient.PullRequests.Create(pr.ctx, pr.repository.Owner, pr.repository.Name, newPR)
	if err != nil {
		return err
	}

	fmt.Printf("PR created: %s\n", p.GetHTMLURL())
	return nil
}
