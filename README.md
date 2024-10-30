# dotkeeper
dotkeeper is a tiny tool for managing dotfiles declaratively.

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
  - src: vim
    dst: ~/.vim
```
