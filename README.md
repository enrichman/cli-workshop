# Section 04

Now let's create our cli.

Import again Cobra with `go get github.com/spf13/cobra`, and create our first root cobra.Command in a `cli.go` file:

```go
func NewRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "stargazer",
		Short: "Stargazer helps you starring Go repositories",
		Long: `
A very simple cli made during a workshop
that helps you searching and starring Go repositories.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
		},
	}
}
```

We can move the logic from the `main` into the `Run` func of the root command, and initialize the cli in the main:

```go
func main() {
	rootCmd := NewRootCmd()
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
```

Let's try to run a `go build` now, and see how it looks:

```
go build -o stargazer .
```

```
./stargazer -h

A very simple cli made during a workshop
that helps you searching and starring Go repositories.

Usage:
  stargazer [flags]

Flags:
  -h, --help   help for stargazer
```

and

```
-> % ./stargazer   
User 'enrichman' (Enrico Candino) found
```

Now let's get the argument at runtime, and run again:

```go
githubService.GetUser(args[0])
```

```
-> % ./stargazer            
panic: runtime error: index out of range [0] with length 0
```

to avoid these errors we can use the `Args` field, and one of the `cobra.PositionalArgs`:

```go
Args: cobra.ExactArgs(1)
```

and now the same command will exit in a nicer way

```
-> % go build -o stargazer . && ./stargazer 
Error: accepts 1 arg(s), received 0
Usage:
  stargazer [flags]

Flags:
  -h, --help   help for stargazer

accepts 1 arg(s), received 0
```

and the `./stargazer enrichman` will work fine

```
-> % ./stargazer enrichman       
User 'enrichman' (Enrico Candino) found
```
