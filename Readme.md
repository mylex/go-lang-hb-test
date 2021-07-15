# HTTP Server To Get Data From Github

### Installation
- Clone the repository
`git clone git@github.com:mylex/go-lang-hb-test.git`
- Go to Repository Directory
`cd go-lang-hb-test`
- Export Github Token as Environment Variable
`export GITHUB_TOKEN={your_token}`
- Get necessary modules.
`go get`

#### Build CLI
`go build ./cmd/hbcli.go`

#### Run CLI
`./cmd/hbcli serve -port=8080`

#### CLI Commands:
    - Serve: For Start HTTP Server 
    - Port (optional): Server Port, Default: 8888

### Test Build
`go build main-test.go`



### Requirements:
Write a HTTP server that satisfies the specifications listed below:

- [x] Give a GitHub user name and repository name as a request parameter.
- [x] Return the number of stars in the repository and the list of followers of the user as a response.
- [x] Use go-github for the implementation.
- [x] Implement the test code.
- [x] Implement a command line tool to start the server.

#### Request parameter
```
{
  "user": "my_username",
  "repo": "my_tool"
}
```

#### Response
```
{
  "star_count": 123,
  "followers": [
    "bob",
    "clare",
    "Davis"
  ]
}
```

#### Other Requirements
Assume that the GitHub access token is stored in the environment variable GITHUB_TOKEN.

##### Bonus if you:
- implement the code in clean and organized package structure
- throttle the API call if it reaches the rate limit
- retry the API call properly if 500 error is responded
- set timeout on API call
- implement the code which works as fast as possible in multi-core environment
- allocate memory as little as possible
- implement the reproducible
- write simple and readable code