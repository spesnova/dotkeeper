# dotkeeper
dotkeeper is a tiny tool for managing dotfiles declaratively.

## Table of Contents
- [Getting Started](#getting-started)
- [Reference](#reference)
- [Contributing](#contributing)
- [License](#license)

## Getting Started
### STEP1. Install dotkeeper
Get the binary from the [Releases](https://github.com/spesnova/dotkeeper/releases) page.

```bash
# Linux AMD64
$ curl -L https://github.com/spesnova/dotkeeper/releases/download/v0.1.0/dotkeeper-linux-amd64 -o dotkeeper

# Linux ARM64
$ curl -L https://github.com/spesnova/dotkeeper/releases/download/v0.1.0/dotkeeper-linux-arm64 -o dotkeeper

# MacOS AMD64
$ curl -L https://github.com/spesnova/dotkeeper/releases/download/v0.1.0/dotkeeper-darwin-amd64 -o dotkeeper

# MacOS ARM64
$ curl -L https://github.com/spesnova/dotkeeper/releases/download/v0.1.0/dotkeeper-darwin-arm64 -o dotkeeper
```

Then, extract the binary and move it to `/usr/local/bin`.

```bash
$ tar -xzf dotkeeper.tar.gz
$ sudo mv dotkeeper /usr/local/bin/
$ chmod +x /usr/local/bin/dotkeeper
```

Check if the binary is installed correctly.

```bash
$ dotkeeper version
v0.1.0
```

### STEP2. Create configurationa file
Move to your dotfiles repository and create `dotkeeper.yaml`.

```bash
$ cd ~/dotfiles
$ touch dotkeeper.yaml
```

```yaml
version: v0

symlinks:
  - src: bash/bashrc
    dst: ~/.bashrc
  - src: vim/vimrc
    dst: ~/.vimrc
  - src: vim/vim
    dst: ~/.vim

git_submodules:
  - path: vim/vim/pack/plugins/start/lightline
    url: https://github.com/itchyny/lightline.vim.git
  - path: vim/vim/bundle/nerdtree
    url: https://github.com/preservim/nerdtree.git

homebrew:
  formulae:
    - git
    - fish
  casks:
    - google-chrome

mas:
  app_ids:
    - 497799835 # Xcode

apt_packages:
  - git
  - fish
```

### STEP3. Apply configuration
Then, run the following command to apply the configuration.
```bash
$ dotkeeper apply
```

It will:
- create symlinks for dotfiles
- initialize git submodules
- install apt packages if you are on Ubuntu/Debian
- install Homebrew packages if you are on macOS
- install Mac App Store applications if you are on macOS

## Reference
### Configuration (v0)
#### Version
The schema version of the configuration file. 

```yaml
version: v0
```

This version is used to validate the configuration file against the CLI version. This version must be a major version of the CLI version.

#### Symlinks
A single file can be symlinked to another location.
```yaml
symlinks:
  - src: bash/bashrc
    dst: ~/.bashrc
```

A directory can be symlinked to another location.
```yaml
symlinks:
  - src: vim/vim
    dst: ~/.vim
```

#### Git Submodules
```yaml
symlinks:
  - src: vim/vim
    dst: ~/.vim

git_submodules:
  - path: vim/vim/pack/plugins/start/lightline
    url: https://github.com/itchyny/lightline.vim.git
  - path: vim/vim/bundle/nerdtree
    url: https://github.com/preservim/nerdtree.git
```

#### Apt Packages
```yaml
apt_packages:
  - git
  - fish
```

#### Homebrew Packages
```yaml
homebrew:
  formulae:
    - git
    - fish
  casks:
    - google-chrome
```

#### Mac App Store Applications
```yaml
mas:
  app_ids:
    - 497799835 # Xcode
```

### Commands
#### `dotkeeper apply`
Apply the configuration.
```bash
$ dotkeeper apply
```

You can specify the path to the configuration file.
```bash
$ dotkeeper apply -c ~/path/to/dotkeeper.yaml
```

#### `dotkeeper version`
Show the version of the CLI.
```bash
$ dotkeeper version
v0.1.0
```

## Contributing
See [CONTRIBUTING.md](CONTRIBUTING.md).

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
