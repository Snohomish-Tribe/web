# Snohomish Tribe

Website for the Snohomish Tribe

## Starting the app

```sh
go run cmd/web/*.go
```

## Pull request preview deploys

This repository includes a GitHub Actions workflow at `.github/workflows/pr-deploy.yml` that deploys PR branches to Fly when a pull request is opened, reopened, updated, or marked ready for review.

The workflow creates a branch-specific Fly app name using the PR branch name. It normalizes the branch name by converting it to lowercase, replacing non-alphanumeric characters with `-`, trimming leading/trailing hyphens, and collapsing repeated hyphens.

Example branch names:

- `feature/New_UI` → `web-pr-feature-new-ui`
- `bugfix/123-fix` → `web-pr-bugfix-123-fix`

After a successful deploy, the workflow comments on the pull request with the preview URL. When the pull request is closed or merged, `.github/workflows/pr-teardown.yml` destroys only the corresponding `web-pr-*` preview app.

The workflow expects the `FLY_API_TOKEN` secret to be configured in the repository settings.

## Go package used

-[Gomap](https://pkg.go.dev/github.com/cwinters8/gomap#section-readme)
