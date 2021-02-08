# Eagle Rock Command Line Utilities

## Setup

### Prerequisites

### Configuration

### Running Tests

## Contribution Guidelines
If you have ideas for improving this project, I welcome both opening an issue on
GitHub or even a PR with the suggested improvement. If you do open a PR, I ask
that you conform to the following guidelines:

### Coding Style---Go Projects
I ask that contributors consider the following:
- conform to community accepted practice for variable names,
- run `gofmt -s` on your code,
- provide documentation comments on each exported interface, struct, struct
  member (unless it implements a function from an interface), and function
  (unexported interfaces, etc. are optional to document, but strongly
  encouraged), and
- if any SQL DDL is written as a part of the change, please include table and
  column comments in the statements.

### Workflow
For the most part, I use the Git Flow workflow. For your forks, obviously, it
doesn't really matter what workflow you use, but I do ask that you:
- sign your commits,
- squash commits before opening a PR, and
- keep changes as incremental as reasonably possible. I find it acceptable to
  implement features over the course of multiple PRs and I promise I won't get
  annoyed if you do that.

### Commit Messages
The rules of [conventional commit
messages](https://www.conventionalcommits.org/en/v1.0.0-beta.2/) should be
observed. Please keep the first line of the commit message down to 50
characters and insert hard line-breaks at 72 characters for the rest of the
message body.

Since I ask that you squash commits before opening PRs, these rules can be
relaxed on topic branches as long as the commit message is conformant before
opening the PR.

#### Tags
  - fix&mdash;for bugfixes
  - feat&mdash;for any new functionality
  - BREAKING CHANGE&mdash;annotation in the commit message body for any changes
    that will affect backwards-compatability.
  - refactor&mdash;for reworked code that ends up being functionally the same
  - docs&mdash;for changes to docstrings, CHANGELOG.md, this README, etc.
  - chore&mdash;for changes to the repo that don't affect functional code or
    docs (i.e. Makefiles, Dockerfiles, etc.)

## License
This project is made available under the MIT License. The text of which you can
find in the LICENSE file of this repository or
[here](https://opensource.org/licenses/MIT).
