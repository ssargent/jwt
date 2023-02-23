# jwt
JWT CLI Utility

Basically this is a really simple tool that will decode a jwt that you give it.  

Let's look at a JWT

```json
{
    "iss": "Online JWT Builder",
    "iat": 1677168418,
    "exp": 1708704418,
    "aud": "www.example.com",
    "sub": "jrocket@example.com",
    "GivenName": "Johnny",
    "Surname": "Rocket",
    "Email": "jrocket@example.com",
    "Role": [
        "Manager",
        "Project Administrator"
    ]
}
```

That when signed might look like this

```
eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJPbmxpbmUgSldUIEJ1aWxkZXIiLCJpYXQiOjE2NzcxNjg0MTgsImV4cCI6MTcwODcwNDQxOCwiYXVkIjoid3d3LmV4YW1wbGUuY29tIiwic3ViIjoianJvY2tldEBleGFtcGxlLmNvbSIsIkdpdmVuTmFtZSI6IkpvaG5ueSIsIlN1cm5hbWUiOiJSb2NrZXQiLCJFbWFpbCI6Impyb2NrZXRAZXhhbXBsZS5jb20iLCJSb2xlIjpbIk1hbmFnZXIiLCJQcm9qZWN0IEFkbWluaXN0cmF0b3IiXX0.VDYoU7n41uyabfQAA0lK8vMfZc6Y1MjNCfWqsttu6bQ
```

To use this tool, all you have to do is copy the jwt into your clipboard (macos: cmd + c,  win|linux ctrl + c) then run the tool via the command line.  The tool when run will copy the text out of your clipboard, decode it and print it to the screen.   Its written in ~30 lines of Go 

```shell
➜  ~ jwt
{
    "Email": "jrocket@example.com",
    "GivenName": "Johnny",
    "Role": [
        "Manager",
        "Project Administrator"
    ],
    "Surname": "Rocket",
    "aud": "www.example.com",
    "exp": 1708704418,
    "iat": 1677168418,
    "iss": "Online JWT Builder",
    "sub": "jrocket@example.com"
}
➜  ~
```

If you want to use this:

```shell
go get github.com:ssargent/jwt
go install

-- or --
go install github.com/ssargent/jwt@latest
```

> You may get linker (ld) warnings on macos.  This is because the library that accesses the clipboard uses cgo.  These are safe to ignore.  
