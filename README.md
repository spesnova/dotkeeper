# dotkeeper
dotkeeper is a tiny tool for managing dotfiles declaratively.

## Installation

## How to use
First, move to the directory containing your dotfiles.
```bash
$ cd dotfiles
```

Then, run the following command to create symlinks and initialize git submodules.
```bash
$ dotkeeper apply
```

## Example Configuration
```yaml
symlinks:
  - src: bash/bashrc
    dst: ~/.bashrc
  - src: vim/vimrc
    dst: ~/.vimrc
  - src: vim/vim
    dst: ~/.vim
```

## References
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
