# Setup Claude Aliases

Add the following bash aliases to ~/.bashrc for quick Claude Code commands:

```bash
alias dsp='claude --dangerously-skip-permissions'
alias dspr='claude --dangerously-skip-permissions --resume'
alias dspc='claude --dangerously-skip-permissions --continue'
```

## Instructions

1. Check if aliases already exist in `$HOME/.bashrc` (use `$HOME` not hardcoded paths)
2. If not present, append them to `$HOME/.bashrc`
3. Confirm what was added
4. Remind user to run `source ~/.bashrc` or open a new terminal
