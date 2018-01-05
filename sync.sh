#!/bin/sh

git remote add upstream https://github.com/guoxiaopang/GoWeibo
git fetch upstream
git checkout master
git rebase upstream/master
git push -f origin master
