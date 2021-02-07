# 環境構築

## WSL2
[ガイド](https://docs.microsoft.com/ja-jp/windows/wsl/install-win10)に従い、WSL2をインストールする。  
Linux distributionは**Ubuntu 20.04 LTS**を使用する。

## Ubuntu 20.04
1. ```wsl -d Ubuntu-20.04```コマンドを実行しUbuntuのshellを起動する。  
1. ユーザの登録が求められるので登録を行う。
1.  パッケージを最新化するため、以下の手順を実施する。
    - ```sudo apt update```
    - ```sudo apt dist-upgrade```
    - ```sudo apt autoremove```
1. Gitのブランチ名をプロンプトに表示させるため、以下の手順を実施する。
    - ```curl -o ~/.git-prompt.sh https://raw.githubusercontent.com/git/git/master/contrib/completion/git-prompt.sh```
    - ```echo 'source ~/.git-prompt.sh' >> ~/.bashrc```
    - **~/.bashrc**に元々存在するPS1をコメントアウトし、以下を設定する。  
    PS1='${debian_chroot:+($debian_chroot)}\[\e[1;32m\]\u@\h\[\e[0m\]:\[\e[1;33m\]\w\[\e[1;36m\]$(__git_ps1)\[\e[0m\]$ '
1. Explorerから **\\\\wsl$\\Ubuntu-20.04**にアクセスし、 任意の場所にProjectを移動させる。  
※必ずUbuntuの中にProjectを移動させること。


## Docker
[ガイド](https://docs.docker.com/docker-for-windows/wsl/)に従いDockerをインストールし、WSL2バックエンドで動作するよう設定する。

## VS Code
1. [VS Code](https://code.visualstudio.com/download)をインストールする。  
1. VS Codeに以下の拡張機能をインストールする。
    - [Docker](https://marketplace.visualstudio.com/items?itemName=ms-azuretools.vscode-docker)
    - [Remote - WSL](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-wsl)
    - [Remote - Containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)
1. Ubuntuのシェル上で以下のコマンドを実行する。
    - ```code {project directory path}```
1. Ubuntu上で起動したVS Codeでコマンドパレットより以下を実行。
    - ```Remote-Containers: Reopen in Container```

# その他

## Build
1. ```ctrl + Shift + B ```またはコマンドパレットから```Tasks: Run Build Task```を実行する。

## Lint
1. ```golangci-lint run ./... --config=.golangci.yml --fast``` を実行する。

## Doc
1. ```godoc -http=:6060``` を実行する。
1. 以下のURLにアクセスする。  
http://localhost:6060/pkg/monkey/
