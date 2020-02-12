#!/bin/bash

# seriously doubt about a upper line. Is it necessary?
cobraMain=$HOME/go/src/github.com/jjwow73/MeerChat/cmd/cobra/main.go
shitShell=$HOME/.zshrc


# 할 것
# 삭제하는 것도 있으면 좋을 듯
# 쉘 재시작 없게끔.
# 다른 쉘들에서도 동작할 수 있게 global한 것은 없을까?

if [[ $SHELL == *"zsh"* ]]; then
    echo "Zsh confirmed"
else
    echo "Only Zsh"
    exit 1
fi

if [ -e $cobraMain ]; then
    if [[ $EUID -ne 0 ]]; then
        echo "You are not running as ROOT(sudo)"
        echo "Please run as ROOT(sudo)"
        echo "Aborted"
        exit 1
    else
        echo "You are running as ROOT(sudo)"
        echo "Proceeding"
        if grep -q meer "$shitShell"; then
            echo "Already installed"
            echo "wanna Delete?"
        else
            echo "alias meer='go run $cobraMain'"
            echo "alias meer='go run $cobraMain'" >> $shitShell
            echo "Setting completed. Please restart zsh!"
            exit 0
        fi
        exit 1
    fi
else
    echo "You dont have files."
    echo "Please get files using 'go get github.com/jjwow73/MeerChat....?[beta]'"
    exit 1
fi
exit 1
