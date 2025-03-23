# docker-compose-remote

Deploy to remote host using docker compose --host feature

```
	docker --host "ssh://root@lts" compose up --build --force-recreate
```

This is a toy project that deploys both tiny service and Redis to remote host.
Then exposes 9003 port to few API calls.
