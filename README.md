# ns git helper

To install, from this directory:

```bash
go install .
```

Set your github fork repo repo:

```bash
git config --global --add ns.fork.org n3wscott
```

Don't forget to have git setup:

```bash
git config --global user.name "Your Name"
git config --global user.email your@email.com
```

`git ns lint` requires:

- `prettier`, needs npm to install: `npm install --global prettier`
- `goimports`, install `go get golang.org/x/tools/cmd/goimports`
