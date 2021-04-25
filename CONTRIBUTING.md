[TOC levels=1-6]: #

# Table of Contents
- [Branching Model](#branching-model)
    - [GitFlow](#gitflow)
    - [Branching Introduction](#branching-introduction)
    - [The Main Branches](#the-main-branches)
    - [Supporting Branches](#supporting-branches)
        - [Feature Branches](#feature-branches)
            - [Creating a feature branch](#creating-a-feature-branch)
            - [Incorporating a finished feature on development](#incorporating-a-finished-feature-on-development)
        - [Release Branches](#release-branches)
            - [Creating a Release Branch](#creating-a-release-branch)
            - [Finishing a Release Branch](#finishing-a-release-branch)
        - [Hotfix Branches](#hotfix-branches)
            - [Creating the Hotfix Branch](#creating-the-hotfix-branch)
            - [Finishing a Hotfix Branch](#finishing-a-hotfix-branch)
    - [Branching References](#branching-references)
- [Git Commit Message Style Guide](#git-commit-message-style-guide)
    - [Commit Message Introduction](#commit-message-introduction)
    - [Commit Messages](#commit-messages)
        - [Message Structure](#message-structure)
        - [The Type](#the-type)
        - [The Scope](#the-scope)
        - [The subject](#the-subject)
        - [The Body](#the-body)
        - [The Footer](#the-footer)
        - [Example Commit Message](#example-commit-message)
- [Versioning Style Guide](#versioning-style-guide)
    - [Semantic Versioning](#semantic-versioning)


# Branching Model

## GitFlow

The branching model this project uses is based on Gitflow. For further
reading please refer to [Branching References](#branching-references)

## Branching Introduction

This Guide documents the development model that this team uses. This
includes branching strategy and release management.

## The Main Branches

The central repo (located at origin) holds two main branches with an
infinite lifetime:

* master
* development

The origin/master branch is the main branch where the source code of
HEAD always reflects a production-ready state

The origin/development branch is the main branch where the source code
of HEAD always reflects a state of the latest delivered development
changes for the next release. Some would refer to this as the
"Integration Branch". This is where any automatic nightly builds are
built from.

When the source code in the development branch reaches a stable point
and is ready to be released, all changes will be merged back into master
and tagged with a release number.

Therefore, each time when changes are merged back into master, this is a
new production release by definition.

## Supporting Branches

The development model uses a variety of supporting branches to aid in
parallel development between team members, ease tracking of features,
prepare for production releases and to assist in quickly fixing live
production problems. Unlike the main branches, these branches always
have a limited life time, since they will be removed when no-longer
needed.

The different types of branches we may use are:

* Feature Branches
* Release Branches
* Hotfix Branches

Each of these branches have a specific purpose and are bound to strict
rules as to which branches may be their originating branch and which
branches must be their merge targets.

### Feature Branches

May Branch off from:
* development

Must merge back into:
* development

Branch naming convention:
* anything except master, develop*, release-\*, or hotfix-\*

Feature branches (Sometimes called topic branches) are used to develop
new features for the upcoming or a distant future release. When starting
development of a feature, the target release in which this feature will
be incorporated may well be unknown at that point. A feature branch
exists as long as the feature is in development, but will eventually be
merged back into development (so that it can be added to an upcoming
release) or discarded (in the case that it was a failed experiment)

#### Creating a feature branch

When starting work on a new feature, branch off from the development
branch.

```git
git checkout -b myfeature development
```

#### Incorporating a finished feature on development

Finished features may be merged into the development branch by creating
a merge request from the repo.

### Release Branches

May Branch off from:
* development

Must merge back into:
* development and master

Branch naming convention:
* release-\*

Release branches support preparation of a new production release. They
allow for last-minute reviews and testing before release. Furthermore,
they allow for minor bug fixes and preparing meta-data for a release
(version number, build dates, etc.). By performing this work in a
release branch, the development branch is cleared to receive features
for the next release.

The time to create a new release branch from development is when the
development branch (almost) reflects the desired state of the new
release. All features that are targeted for the release-to-be-built must
be merged into development at this point. All features targeted at
future releases may not. They must wait until after the release branch
has been created.

At the start of the release branch, the release is assigned a version
number. The development branch reflects the changes for the "next
release", but it is unclear whether that "next release" will eventually
become 0.3 or 1.0, etc... until the release branch is started. The
version is determined by our version system discussed later in this
document.

#### Creating a Release Branch

Release branches are created from the development branch. For example,
version 1.1.5 is the current production release and we a big release
coming up. The state of development is ready for the "next release" and
we have decided that this will become version 1.2 (instead of 1.1.6 or
2.0). We create a branch and give the release branch a name reflecting
the new version number:

```git
git checkout -b release-1.2 development

$ .\bump-version.sh 1.2 # Run our versioning tool

git commit -a -m "chore: Bump version number 1.2"
```

After creating a new branch and switching to it, we bump the version
number. bump-version.sh is currently a fictional shell script that
changes the version across all the required files for the release. Then,
the bumped version number is committed.

The release branch may exist there for a while, until the release is has
been fully rolled out. During that time, bug fixes may be applied in
this branch (rather than on the develop branch). Adding large new
features here is strictly prohibited. They must be merged into develop,
and therefore, wait for the next big release.

#### Finishing a Release Branch

When the state of the release branch is ready to become a real release,
the branch is merged into master (Since every commit on master is a new
release). Next the commit on master must be tagged for easy future
reference. Finally, the changes made on the release branch need to be
merged back into development to sync up the changes (bug fixes, ext...).
This is completed by submitting a merge request through the repo.

```git
$ git pull # sync changes

$ git checkout master # checkout current master

$ git tag -a 1.2 # Tag the committed release
```

The release is now done, and tagged for future reference.

We now need to merge the changes back into the development branch. This
is done by submitting a merge request through the repo. This step may
lead to a merge conflict, if so, resolve it and commit.

The release branch can now be removed at this stage.

### Hotfix Branches

May branch off from:
* master

Must merge back into:
* development
* master

Branch naming convention:
* hotfix-\*

Hotfix branches are similar to release branches in that they are
intended to prepare for a new production release, though unplanned. They
arise from the necessity to act immediately upon an undesired state of a
live production version. When a critical bug in a production version
must be resolved immediately, a hotfix branch may be branched off from
the corresponding tag on the master branch that marks the production
version.

The idea is that work on the development branch can continue, while
preparing a quick production fix.

#### Creating the Hotfix Branch

Hotfix branches are created from the master branch. For example, say
version 1.2 is the current production release running live and causing
troubles due to a severe bug. But changes on develop are yet unstable.
We may then branch off a hotfix branch and start fixing the problem.

```git
git checkout -b hotfix-1.2.1 master

.\bump-version.sh 1.2.1 # Run our versioning tool

git commit -a -m "chore: Bumped version number to 1.2.1"
```

You must bump the version number after branching off.

Then, fix the bug and commit the fix in one or more separate commits

```git
git commit -m "fix: severe production problem"
```

#### Finishing a Hotfix Branch

When the bug has been fixed it now needs to be merged back into master,
but also needs to be merged back into development in order to insure
that it is included in the next release.

Submit a merge request to master through the repo. Then perform the
following:

```git
git pull # this syncs your local copy with origin

git tag -a 1.2.1 # tag the release with the ner version
```

Next, submit a merge request to development:

The one exception to the rule here is when a release branch currently
exists, the hotfix changes need to be merged into that release branch,
instead of development. Back-merging the hotfix into he release branch
will eventually result in the hotfix being merged into development when
the release branch is finished.

If work in the development branch immediately requires this hotfix and
cannot wait for the release branch to be finished, you may safely merge
the hotfix into the development branch now as well.

The final step is to remove the temporary branch created for the hotfix.

## Branching References

The following references where consulted to create the branching model:

* [GitFlow](http://nvie.com/posts/a-successful-git-branching-model/)

# Git Commit Message Style Guide

## Commit Message Introduction

This style guide acts as the official guide to follow when crafting your
Git Commit Messages. There are many opinions on the "ideal" style in the
world of development. Therefore, in order to reduce confusion on what
style contributors should following during the development life cycle,
contributors should refer to this style guide.

## Commit Messages

### Message Structure

A commit message consists of three distinct parts separated by a blank
line: the title, an optional body and an optional footer. The layout
looks like this:

```
type(scope): subject

body

footer
```

The title consists of the type of the message and subject

### The Type

The type is contained within the title and can be one of these types:

* **feat**: a new feature
* **fix**: a bug fix
* **wip**: work in progress. fomat for use is `wip: <type>: <description>`
  This type is used when you need to make a commit but are not at a
  logical point but want/need to create a save point. This type will be
  ignored during change log creation. your final commit should be the
  `<type>` of change you where making
* **docs**: changes to documentation
* **style**: formatting, missing semi colons, etc; no code change
* **refactor**: refactoring production code
* **test**: adding tests, refactoring test; no production code change
* **chore**: updating build tasks, package manager configs, etc; no
  production code change

### The Scope

The scope is used to specify the location of the commit change. for
example if you where working on the document build.md, the scope would
be build.md. doc(build.md): U

### The subject

Subjects should be no greater than 50 characters, should begin with a
capital letter and do not end with a period

Use an imperative tone to describe what a commit does, rather than what
it did. For example, use change; not changed or changes.

### The Body

Not all commits are complex enough to warrant a body, therefore it is
optional and only used when a commit requires a bit of explanation and
context. Use the body to explain the **what** and **why** of a commit,
not the how.

When writing a body, the blank line between the title and the body is
required and you should limit the length of each line to no more than 72
characters.

### The Footer

All breaking changes have to be mentioned in the footer with the
description of the change, justification and migration notes.

The footer is also used to reference issue tracker IDs.

### Example Commit Message

```
feat(scope): Summarize changes in around 50 characters or less

More detailed explanatory text, if necessary. Wrap it to about 72
characters or so. In some contexts, the first line is treated as the
subject of the commit and the rest of the text as the body. The
blank line separating the summary from the body is critical (unless
you omit the body entirely); various tools like `log`, `shortlog`
and `rebase` can get confused if you run the two together.

Explain the problem that this commit is solving. Focus on why you
are making this change as opposed to how (the code explains that).
Are there side effects or other unintuitive consequenses of this
change? Here's the place to explain them.

Further paragraphs come after blank lines.

 - Bullet points are okay, too

 - Typically a hyphen or asterisk is used for the bullet, preceded
   by a single space, with blank lines in between, but conventions
   vary here

BREAKING CHANGE: Describe the change that was made that breaks existing
changes.

To migrate the code follow the example below:

Before:

(Before usage)

After:

(After usage)

(Justification)

If you are working on a issue from the issue tracker, put references to them at the bottom,
like this:

Resolves: #123
See also: #456, #789
```

# Versioning Style Guide

### Semantic Versioning

This project uses [Semantic Versioning](http://semver.org/) to perform
versioning related tasks.

Given a version number MAJOR.MINOR.PATCH, increment the

1. MAJOR version when you make incompatible API changes
2. MINOR version when you add functionality in a backwards-compatible
   manner
3. PATCH version when you make backwards-compatible bug fixes

Additional labels for pre-release and build metadata are available as
extensions to the MAJOR.MINOR.PATCH format.

Acceptable pre-release lables are the following:

* -rc.* - The * represents the current release candidate. For example
  1.0.0-rc.1
* -alpha - this represents an early alpha release that has not be fully
  vetted or had its features solidified. For example 1.0.0-alpha
* -beta - this represents a pre-release that has had its features set,
  but may still contain bugs. for example 1.0.0-beta
