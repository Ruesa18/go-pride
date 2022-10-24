# go-pride

## What is "go-pride"?
The go-pride project introduces a new command for your terminals. It's basically a go program that allows you to show your pride flag! 🏳️‍🌈😃

## Installation
You can use `go install pride.go` to build and install this program directly into your Go working directory. In order to be able to run commands installed like this, you will have to add the Go working directory to your path variable.
After the installation you can use the command `pride` from inside your favorite terminal.

## Configuration
After executing the pride command at least once, you will be able to find a config file in your home directory under: `.config/go-pride/config.json`.
You may use this config.json file to create your own custom flag.
The `FlagColors` attribute describes what colors will be used on each line of the flag. For more information on what colors are available, have a look into [this repository](https://github.com/gookit/color).

![image](https://github.com/gookit/color/raw/master/_examples/images/color-256.png)

You may use some of the example configurations. You can find them in the `examples/` directory of this project.

## Contributing
I would like to implement the go code by myself. But what you may contribute are the configurations for different flags (can be found under `examples/`). Please create a json-file with the name of your sexuality and choose the same name as the branch name for your feature.

Of course you may also create new issues with requests for features. Please use the `wishlist` label for your requests.

## Credits
I have been inspired by [this repository](https://github.com/ExperiBass/cli-pride-flags) of [ExperiBass](https://github.com/ExperiBass).
It's a good solution, but I wanted to be able to scale it and also use it as my terminal welcome screen (bashrc). 
Also I wanted to learn Go and this was a fitting first project 😉
