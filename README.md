# golang-learning
Experiments with GoLang
## Install (Brew)

```sh
# Install golang
brew update && brew install golang

# Setup Workspace
mkdir -p $HOME/go/{bin,src,pkg}

# Setup environment
vim ~/.zshrc

# export GOPATH=$HOME/go
# export GOROOT="$(brew --prefix golang)/libexec"
# export PATH="$PATH:${GOPATH}/bin:${GOROOT}/bin"

source ~/.zshrc
```