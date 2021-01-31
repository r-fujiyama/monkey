FROM golang:1.15.7-alpine3.12
WORKDIR /home
RUN apk update && \
    # Tools
    apk add --no-cache curl && \
    apk add --no-cache vim && \
    apk add --no-cache git && \
    apk add --no-cache gcc && \
    apk add --no-cache musl-dev && \
    # Bash
    apk add --no-cache bash && \
    apk add --no-cache bash-completion && \
    sed -e 's;/bin/ash$;/bin/bash;g' -i /etc/passwd && \
    curl -o ~/.git-completion.bash https://raw.githubusercontent.com/git/git/master/contrib/completion/git-completion.bash && \
    echo 'source ~/.git-completion.bash' >> ~/.bashrc && \
    curl -o ~/.git-prompt.sh https://raw.githubusercontent.com/git/git/master/contrib/completion/git-prompt.sh && \
    echo 'source ~/.git-prompt.sh' >> ~/.bashrc && \
    echo "PS1='\[\e[1;32m\]\u@\h\[\e[0m\]:\[\e[1;33m\]\w\[\e[1;36m\]"'$(__git_ps1)\[\e[0m\]$ '"'" >> ~/.bashrc && \
    echo 'if [ -n "$BASH_VERSION" ]; then' >> ~/.profile && \
    echo '  if [ -f "$HOME/.bashrc" ]; then' >> ~/.profile && \
    echo '    . "$HOME/.bashrc"' >> ~/.profile && \
    echo '  fi' >> ~/.profile && \
    echo 'fi' >> ~/.profile && \
    # Git Config
    git config --global core.editor vim && \
    git config --global core.commentChar ';' && \
    # Golang
    GO111MODULE=on go get golang.org/x/tools/gopls@latest && \
    go get -u github.com/go-delve/delve/cmd/dlv && \
    go get -u github.com/uudashr/gopkgs/cmd/gopkgs && \
    go get -u github.com/ramya-rao-a/go-outline && \
    go get -u github.com/haya14busa/goplay && \
    go get -u github.com/fatih/gomodifytags && \
    go get -u github.com/josharian/impl && \
    go get -u github.com/cweill/gotests/... && \
    go get -u golang.org/x/tools/cmd/godoc  && \
    # golangci-lint
    curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.35.2
