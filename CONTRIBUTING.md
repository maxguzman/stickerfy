# How to contribute

We are really glad you're reading this, because you are motivated to collaborate with me.

## Collaboration Submitting changes

Please send a [GitHub Pull Request](/.github/pull_request_template.md) with a clear list of what you've done.

Please keep it mind, I can require that pull requests are approved before being merged.
Always write a clear log message for your commits. One-line messages are fine for small changes, but bigger changes should look like this:

    $ git commit -m "feat: A brief summary of the commit
    >
    > A paragraph describing what changed and its impact."

## Coding conventions

I love OpenSource, so I follow standardized commit message while doing so.

The commit contains the following structural elements, to communicate intent to the consumers this repo:

* **fix**: a commit of the type fix patches a bug in your codebase (this correlates with PATCH in semantic versioning).
* **feat**: a commit of the type feat introduces a new feature to the codebase (this correlates with MINOR in semantic versioning).
* **Others**: for example recommends chore:, docs:, style:, refactor:, perf:, test:, and others. We also recommend improvement for commits that improve a current implementation without adding a new feature or fixing a bug.
* **BREAKING CHANGE**: a commit that has the text BREAKING CHANGE: at the beginning of its optional body or footer section introduces a breaking API change (correlating with MAJOR in semantic versioning). A breaking change can be part of commits of any type. e.g., a fix:, feat: & chore: types would all be valid, in addition to any other type.
