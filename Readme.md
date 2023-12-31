# Volume CLI

Volume CLI is an open-source command-line interface (CLI) tool for controlling your computer's volume. It is built using Go and the Cobra library, and it interfaces with the ALSA library to control the volume.

## Features

Volume CLI supports the following commands:

- `volumecli up` - Increase volume by two percent
- `volumecli down` - Decrease volume by two percent
- `volumecli set <volume>` - Set the volume to a specific value
- `volumecli mute` - Mute the volume
- `volumecli unmute` - Unmute the volume
- `volumecli toggle` - Toggle mute/unmute

It also supports optional flags for specifying the mixer and element:

- `volumecli -mixer "default" -element "Master" <command>`

## Installation

Download the latest release from github and extract the binary to your preferred location.
Link to latest release: https://github.com/Fingann/volumecli/releases/latest

## Usage 

To use Volume CLI, you can run the binary directly:

```bash
# Increase volume by two percent
./volumecli up

# Decrease volume by two percent
./volumecli down

# Set volume to 50 percent
./volumecli set 50

# Mute the volume
./volumecli mute

# Unmute the volume
./volumecli unmute

# Toggle mute/unmute
./volumecli toggle

# Set volume to 50 percent with specific mixer and element
./volumecli -mixer "default" -element "Master" set 50
```

# Building with Docker
This project includes a Dockerfile and a Makefile that can be used to build the project inside a Docker container. This can be useful if you want to build the project in a consistent environment, or if you don't want to install the necessary dependencies on your local machine.

To build the project with Docker, you will need to have Docker installed on your machine. You can download Docker from the official website.

Once you have Docker installed, you can use the Makefile to build the project:
```bash
# Build the Docker image and compile the project
make run

```
This will create a Docker image with all the necessary dependencies, compile the project inside a Docker container, and then copy the resulting binary to your local machine. The binary will be located in the ./bin directory.

Please note that the Docker build process uses the Golang Docker image as a base, which includes the Go compiler and other necessary tools for building Go projects.

## Cleaning up
After you're done, you can use the Makefile to clean up the Docker image and the compiled binary:

```bash
# Remove the Docker image and the compiled binary
make clean
```
This will remove the Docker image and the compiled binary from your local machine. If you want to build the project again, you will need to run make run again.

## Troubleshooting
If you encounter any issues while building the project with Docker, please check the following:

 - Make sure you have Docker installed and running on your machine.
 - Make sure you have permission to run Docker commands. On some systems, you might need to use sudo to run Docker commands.
 - Make sure you have enough disk space to build the Docker image and the project.
 - If you're getting errors related to ALSA or other dependencies, make sure the Dockerfile includes all the necessary dependencies for the project.

# Contributing
If you would like to contribute, you can fork this repository, make your changes, and then submit a pull request.

# License

Volume CLI is licensed under a MIT license. See the [LICENSE](LICENSE) file for more information.