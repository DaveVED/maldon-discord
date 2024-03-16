<div align="center">

# Maldon Discord Bot

##### Versatile tool to enhance the discord experience

[![GO](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://www.go.dev)
[![Discord](https://img.shields.io/badge/Discord-5865F2?style=for-the-badge&logo=discord&logoColor=white)](https://www.mongodb.com/)

<img alt="Maldon Discord" height="280" src="/assets/maldon-discord-icon.png" />

</div>

## ⇁ Table of Contents
* [Getting Started](#-Getting-Started)
* [Commands](#-Commands)
* [Concepts](#-Concepts)
* [Contribution](#-Contribution)

## ⇁ Getting Started

## ⇁ Commands

The following commands are available to use:

| Command | Inputs  | Description |
| :------- | :-----   | :----------- |
| `!ping` | `na`    | Sends `pong` to the channel that recieved the message. This is simple health check. |
| `!whoami ` | `na` | Sends custom message to the server with user details |

We also have a few passive commands implmented:

| type | Description |
| :---- | :----------- |
| `WelcomeNewMember` | Sends a welcome message to a `welcome` channel when new members join. |

## ⇁ Concepts

### Adding in a Command

Take our `internal/commands` directory structure as an example:

```bash
├── internal
│   ├── commands
│   │   ├── handler.go
│   │   ├── pingpong.go
│   │   ├── xyz.go
```

From this you'll notice that we have two key things going on we have 
`handler.go` and then our command files. When adding a new command you can 
leverage the same structure.

- `handler.go` - This file is the main entry point for the bot. It's where we 
  handle the incoming messages and then route them to the correct command.

- command files - These are the files that contain the actual command logic. 
  They are named after the command they are handling. For example, `pingpong.go` 
  contains the logic for the `!ping` command.

If the `handler.go` does not have support for what you are trying to do, you
will also need to update the `/internal/maldon/maldon.go` file to add the new
hanlder. For example, `dg.AddHandler(commands.XZY)` but you should only do this
*if and only if* there is not already a hanlder you can reuse.

### Hosting your own instance

You can host you own instance if you want to use the bot in your own server. We store the image on `DockerHub` so you can get started with some easy steps.

```bash
todo
```

If you have made custom modifications to the bot, you can build your own image and host it on your own server.

```bash
docker build -t maldon-discord:latest -f build/Dockerfile .
```

## ⇁ Contribution

We're excited to announce that this project is now officially open source, which 
means we're not just allowing access to the source code but are actively 
inviting contributions from the community. If you're interested in 
contributing, here's how to get started:

1. Before making any changes, please open an issue to discuss your proposed contribution.
1. Create your branch `git checkout -b feature/fooBar`
1. Commit your changes `git commit -am 'feat: some message'`
1. Push to the branch `git push origin feature/fooBar`
1. Create a new Pull Request
