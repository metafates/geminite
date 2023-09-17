# Geminite

Article reader for you terminal!
No ads, no JavaScript, no CSS, nothing bascially.

[![asciicast](https://asciinema.org/a/JKLi3sao0ZDuKFiU49bnn6jGy.svg)](https://asciinema.org/a/JKLi3sao0ZDuKFiU49bnn6jGy)

> [!NOTE]  
> A proof-of-concept app. Lacks many features. I'm just having fun 😜

## What?

1. You give it an URL
2. It downloads HTML (or fetches it from cache)
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

## Bookmarks

You can press <kbd>b</kbd> when reading an article to bookmark it.

To open your bookmarks list, run `geminite` without any arguments

## Config

To show config file location run

```bash
geminite where
```

Config is in TOML format. Default config example

```toml
# Words per minute reading speed
# You can get your own here
# https://swiftread.com/reading-speed-test
wpm = 250

# Enable caching
cache = true
```
