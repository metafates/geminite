# Geminite

Article reader for you terminal!
No ads, no JavaScript, no CSS, nothing bascially.

[![asciicast](https://asciinema.org/a/JKLi3sao0ZDuKFiU49bnn6jGy.svg)](https://asciinema.org/a/JKLi3sao0ZDuKFiU49bnn6jGy)

> [!NOTE]  
> A proof-of-concept app. Lacks many features. I'm just having fun 😜

## What?

1. You give it an URL
2. It downloads HTML
3. Extracts readable part (like reader mode does in Firefox; [thx Mozilla](https://github.com/mozilla/readability))
4. Converts into markdown
5. Nicely displays it inside TUI
6. ...
7. PROFIT!!!

## Build

To build `geminite` binary clone this repo and run

```bash
go build .

# You can also use go install
go install .
```

## TODO

- Bookmarks (like your cool browser does)
- Caching (do not download the same page twice)
- Configuration (e.g. set you own reading speed in words per minute for your personal reading time estimation)
