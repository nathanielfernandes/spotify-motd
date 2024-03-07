## spotify-motd

a fake minecraft server that shows my current spotify song as the motd (message of the day)

![screenshot_1](./screenshots/ss1.png)
![screenshot_2](./screenshots/ss2.png)
![screenshot_3](./screenshots/ss3.png)

## how it works

this server listens for my spotify activity changes through server sent events from https://github.com/nathanielfernandes/watcher

it then uses https://github.com/nathanielfernandes/motd to handle the tcp connection and send the motd to the client

## ip

`mc.ncp.nathanferns.xyz`
