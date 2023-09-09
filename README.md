# Section 01

Welcome! This is the first section, and we are going through some basic steps to set you up and running.

You should already have installed Go. If not, please follow the official guide [here](https://go.dev/doc/install).

And now let's see if everything works!

## Let's start!

Create a `main.go` file that will print the classic Hello World!

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
}
```

and run `go run main.go`. You should see:

```
> % go run main.go
Hello world!
```

Awesome!

## Go modules

Modules are how Go manages dependencies.

If you want to know more about them you can have a read at the official documentation [here](https://go.dev/ref/mod).

For now it's enough to know that to set them you you will need to run just a `go mod init <mypackage>`:

```
go mod init github.com/enrichman/cli-workshop
```

Now you should habe a `go.mod` file in your folder.

Some more useful commands that you should know about modules and dependencies are
- `go mod tidy` to cleanup your dependencies
- `go get <package>` to get/download a dependency

For example we can try to download Cobra running

```
go get github.com/spf13/cobra
```

This will download `spf13/cobra`, its dependencies, and it will change the `go.mod` accordingly.
It will also create a `go.sum` file with the checksums of these libraries.

We can now run a `go mod tidy` that will cleanup the `go.mod`, because in our project we are not using Cobra yet.
We can delete the empty `go.sum` as well.
