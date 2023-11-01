package cli

import "context"

type CLI struct {
}

func New() *CLI {
	return &CLI{}
}

func (cli *CLI) DoSomething(ctx context.Context, arg1 string, arg2 int) error {

}
