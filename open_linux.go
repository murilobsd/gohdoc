package main

import (
	"context"
	"os/exec"
)

import (
	"context"
	"os/exec"
)

func openBrowserCmd(ctx context.Context, url string) *exec.Cmd {
	return exec.CommandContext(ctx, "xdg-open", url) // linux
}
