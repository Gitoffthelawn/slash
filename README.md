# Slash

<img align="right" src="./resources/logo.png" height="64px" alt="logo">

`Slash` is a bookmarking and link shortening service that enables easy saving and sharing of links. It allows you to store, categorize, and share links with custom short URLs. You can search, filter, and access your saved links from any device. It also supports team sharing of link libraries for easy collaboration.

Try it out on <a href="https://slash.stevenlgtm.com">Live Demo</a>.

<p>
  <a href="https://discord.gg/QZqUuUAhDV"><img alt="Discord" src="https://img.shields.io/badge/discord-chat-5865f2?logo=discord&logoColor=f5f5f5" /></a>
  <a href="https://hub.docker.com/r/yourselfhosted/slash"><img alt="Docker pull" src="https://img.shields.io/docker/pulls/yourselfhosted/slash.svg" /></a>
  <a href="https://github.com/boojack/slash/stargazers"><img alt="GitHub stars" src="https://img.shields.io/github/stars/boojack/slash?logo=github" /></a>
</p>

## Features

- Create customizable `/s/` short links for any URL.
- Share short links privately or with others.
- View analytics on short link traffic and sources.
- Open source self-hosted solution.

## Deploy with Docker in seconds

```bash
docker run -d --name slash -p 5231:5231 -v ~/.slash/:/var/opt/slash yourselfhosted/slash:latest
```

Learn more in [Self-hosting Slash with Docker](https://github.com/boojack/slash/blob/main/docs/install.md).
