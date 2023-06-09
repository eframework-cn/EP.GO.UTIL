#!/bin/sh

git fetch --tags
latest=$(git describe --tags `git rev-list --tags --max-count=1`)
echo -e "Latest tag is "$latest
read -p "Please input version(ex: v1.0.0) for publising: " version
git tag ${version}
git push origin ${version}
read -p "publish done."