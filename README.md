## Study JWT

```
curl -XGET localhost:8080/login
```
This command can supply JWT.
```sample response
{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNjE4MjM3MTk5LCJpYXQiOiIyMDIxLTA0LTEyVDIyOjE5OjU5LjY1NTE4MSswOTowMCIsIm5hbWUiOiJ0ZXN0LXVzZXIiLCJzdWIiOiJ0ZXN0LXVzZSIsInlvdGVpIjoiZ28gdG8gaG9nZWhvZ2UifQ.TzCOS7fBxiTQZiWp4I3V1c2YDxIsM5Tr_uvqGjDCS_c"}
```

Use this, you can
```
curl -XGET localhost:8080/restricted -H 'Authorization: Bearer {token}'
```

Then, you can get claims in jwt.