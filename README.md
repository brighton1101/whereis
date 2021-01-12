# whereis

#### A simple Go application telling you where those shortened links are redirecting you to.

### Example Usage:
```sh
🌴🌴🌴 whereis (master) $ go run cli.go -b buff.ly/1irhfHu
Warning: original url modified from buff.ly/1irhfHu to https://buff.ly/1irhfHu

http://www.huffingtonpost.com/entry/jiff-the-dog-wins-halloween_56327e41e4b00aa54a4d7a89?ncid=fcbklnkushpmg00000022&utm_content=buffer83279&utm_medium=social&utm_source=twitter.com&utm_campaign=buffer
Should open in browser (y/n): n
```

### Cli
- Optional args (must come before link as positional arg)
    - `-b` Prompts you to open link in browser, then will open link in browser
    - `-c` Will copy resulting link to clipboard (NOTE: only tested on MacOS as of now)
- Positional arg is shortened link
- Example `(executable) -b -c [shortened link here]`
