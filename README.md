# git-diffstat

`git-diffstat` is a tool to format the output of `git diff --numstat`.

## Example

```console
> git-diffstat HEAD^
path       add  delete
----------------------
README.md   14       0
```

Get the diff of the first commit(694cb8272e1f1b7f917f151bd2fab05755552d8a). 4b825dc642cb6eb9a060e54bf8d69288fbee4904 is the empty tree object.
```console
> git-diffstat 4b825dc642cb6eb9a060e54bf8d69288fbee4904..694cb8272e1f1b7f917f151bd2fab05755552d8a
path         add  delete
------------------------
LICENSE.txt   21       0
README.md     18       0
git-stat.go   94       0
git.go        25       0
```

## License

This software is released under the MIT License, see LICENSE.txt.

## Author

TAKAHASHI Satoshi <hikobae@gmail.com>