---
title: Release
---

data-control-center tries to follow [Semantic Versioning](https://semver.org/).

## Process

**INFO**: You must be on the `main` branch to run this script!

Use the following script to update the data-control-center version:

```console
./internal/scripts/bump_version.sh X.Y.Z
```

After the release has been successfully created make sure to rebase and force push the `develop` branch on `main` branch.

```console
git checkout main
git pull
git checkout develop
git rebase main
git push -u develop --force
```
