# Volume CLI

Volume CLI is an open-source command-line interface (CLI) tool for controlling your computer's volume. It is built using Go and the Cobra library, and it interfaces with the ALSA library to control the volume.

## Features

Volume CLI supports the following commands:

- `vcli up` - Increase volume by two percent
- `vcli down` - Decrease volume by two percent
- `vcli set <volume>` - Set the volume to a specific value
- `vcli mute` - Mute the volume
- `vcli unmute` - Unmute the volume
- `vcli toggle` - Toggle mute/unmute

It also supports optional flags for specifying the mixer and element:

- `vcli -mixer "default" -element "Master" <command>`

## Installation

To install Volume CLI, you need to have Go installed on your system. You can download Go from the [official website](https://golang.org/dl/).

Once you have Go installed, you can clone this repository and build the project:

```bash
git clone https://github.com/<your-username>/volumecli.git
cd volumecli
go build
```
This will create a binary named \`***vcli***\` in the current directory.

## Usage 

To use Volume CLI, you can run the binary directly:

```bash
# Increase volume by two percent
./vcli up

# Decrease volume by two percent
./vcli down

# Set volume to 50 percent
./vcli set 50

# Mute the volume
./vcli mute

# Unmute the volume
./vcli unmute

# Toggle mute/unmute
./vcli toggle

# Set volume to 50 percent with specific mixer and element
./vcli -mixer "default" -element "Master" set 50
```

# Contributing
If you would like to contribute, you can fork this repository, make your changes, and then submit a pull request.

# License

Volume CLI is licensed under a MIT license. See the [LICENSE](LICENSE) file for more information.