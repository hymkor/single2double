[![GoDev](https://pkg.go.dev/badge/github.com/hymkor/single2double)](https://pkg.go.dev/github.com/hymkor/single2double)

single2double.exe
=================

In Windows, we have to write one-liner with `\"` for string literals.

```sh
cat issues.json | jq -r ". | map(\"#\"+(.number|tostring)+\" [\"+ .user.login + \"](\"+ .user.html_url + \")\") | join(\"\r\n\")"
```

The single2double.exe replaces ' with \" on the command line and executes it.

```sh
cat issues.json | single2double.exe jq -r ". | map('#'+(.number|tostring)+' ['+ .user.login + ']('+ .user.html_url + ')') | join('\r\n')"
```

Install
-------

Download the binary package from [Releases](https://github.com/hymkor/single2double/releases) and extract the executable.

### via go install

```
go install github.com/hymkor/single2double@latest
```

### via scoop-installer

```
scoop install https://raw.githubusercontent.com/hymkor/single2double/master/single2double.json
```

or

```
scoop bucket add hymkor https://github.com/hymkor/scoop-bucket
scoop install single2double
```
