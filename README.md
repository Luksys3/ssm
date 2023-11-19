# SSM - Simple SSH Manager

SSM is a command-line interface tool written in Go, designed for simple management of SSH connections. It simplifies the process of connecting and copying SSH keys to remote servers.

## Installation

```
git clone git@github.com:Luksys3/ssm.git
cd ssm
go install
```

## Commands

#### Connect to a Server (Default Command)

- **Command**: `ssm`
- **Usage**: Displays a list of configured servers and allows you to select one for connection.

#### Connect to a Server

- **Command**: `ssm connect [name] [environment]`
- **Usage**: Initiates an SSH session to a specified server.
- **Arguments**:
  - `name`: Name of the server (required).
  - `environment`: The environment to connect to (optional).

#### Edit Server Configuration

- **Command**: `ssm gedit`
- **Usage**: Opens the servers configuration file in Gedit for editing.

#### Copy SSH Public Key

- **Command**: `ssm copy-id`
- **Usage**: Copies the user's SSH public key (`~/.ssh/id_ed25519.pub`) to a selected server.

## Servers File Structure

Servers configuration are stored in a file located at `~/.ssm/servers`. The format of the file is as follows:

```
name user@ip [environment]
...
```

Each line represents a server configuration with the following format:

- `name`: A unique identifier for the server.
- `user@host`: The username and IP address of the server.
- `environment` (optional): An optional tag to specify the environment (e.g., prod).

Can be edited with `ssm gedit` command.

## License

MIT
