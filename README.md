# dotkeeper
dotkeeper is a tiny tool for managing dotfiles declaratively.

## Table of Contents
- [Getting Started](#getting-started)
- [Reference](#reference)
- [License](#license)

## Getting Started
### STEP1. Install dotkeeper
First, install dotkeeper.
```bash
```

### STEP2. Create configurationa file
Move to your dotfiles repository and create `dotkeeper.yaml`.

```bash
$ cd ~/dotfiles
$ touch dotkeeper.yaml
```

```yaml
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

## Reference
### Configuration
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

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
