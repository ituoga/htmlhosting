# Basic setup to start project in go


requirements: 
golang, pnpm

```
go install github.com/go-task/task/v3/cmd/task@latest
task deps
task tools
./rename.sh github.com/yourusername/yourrepo

```

and then run `task -w`



Inspiration from https://github.com/zangster300/gonads-starter
