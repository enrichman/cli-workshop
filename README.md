# Section 03

Now let's cleanup a bit the code, and move this into a GithubService.

Create a `service.go` file, and a struct that will help us calling the Github APIs.

```
type GithubService struct{}
```

This service will have a `GetUser` method that will return a `User` struct, or an error:

The user will be defined as

```go
type User struct {
	Username string
	Name     string
}
```

and the method GetUser

```go
func (s *GithubService) GetUser(username string) (*User, error) {
	// TODO
}
```

Let's copy paste all the code in the `main()` func, and put it in the `GetUser`. Then let's refactor it a bit.

1) Format the URL to get the username from the argument of the method

```go
url := fmt.Sprintf("https://api.github.com/users/%s", username)
resp, err := http.Get(url)
```

1) Handle the errors from the `log.Fatal`

```go
return nil, fmt.Errorf("getting user: %w", err)
```

2) returning a User in case of success

```go
return &User{
	Username: user["login"],
	Name:     user["name"],
}, nil
```

This will complain, because the type returned is `any` (or `interface{}`), while we need a `string`.

We can do a type assertion

```go
return &User{
	Username: user["login"].(string),
	Name:     user["name"].(string),
}, nil
```

or even better we could map this struct with JSON tags.

```go
type User struct {
	Username string `json:"login"`
	Name     string `json:"name"`
}
```

Doing so we can say to the `json.Unmarshal` to map the fields into the corresponding variables, instead of using a generic map:

```go
user := &User{}
err = json.Unmarshal(bodyBytes, user)
if err != nil {
	return nil, fmt.Errorf("unmarshalling the body: %w", err)
}

return user, nil
```

Now from the `main()` we can simply create a `GithubService` and call the `GetUser` func:

```go
func main() {
	githubService := &GithubService{}

	user, err := githubService.GetUser("enrichman")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("User '%v' (%v) found\n", user.Username, user.Name)
}
```

Now if we run `go run main.go` it will fail, because we have more than one file. So we need to do a `go run ./...`.

```
-> % go run ./...  
User 'enrichman' (Enrico Candino) found
```
