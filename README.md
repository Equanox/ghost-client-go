# ghost-client-go
Inclomplete api client for https://github.com/TryGhost/Ghost.

# Getting Started

* Start a ghost instance using docker https://hub.docker.com/_/ghost
* Create a key for the content api


```
u, _ := url.Parse("http://localhost:3001/ghost/api/v3")

ghostClient := ghostclient.New(
    ghostclient.WithBaseURL(u),
)

posts, _ := ghostClient.ContentPosts(
	ghostclient.WithKey("62ebdfd5992678506b1e7ddbc3"),
	ghostclient.WithInclude("tags", "authors"),
)

litter.Dump(posts) // Print posts to stdout
```
[Example](./example/main.go)