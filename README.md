single2double.exe
=================

In Windows, we have to write one-liner like this

```sh
cat issues.json | jq -r ". | map(\"#\"+(.number|tostring)+\" [\"+ .user.login + \"](\"+ .user.html_url + \")\") | join(\"\r\n\")"
```

Single2double.exe simplifies command-lines

```sh
cat issues.json | single2double.exe jq -r ". | map('#'+(.number|tostring)+' ['+ .user.login + ']('+ .user.html_url + ')') | join('\r\n')"
```
