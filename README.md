# automated-discord-trickcord-treat

Automation of the recuperation of Discord's Halloween bot.

This project was make for fun, to learn more about the Discord API.

## How to use

Create a .env file with props:

`DISCORD_USER_TOKEN` Your user account's token

`DISCORD_CHANNEL_ID_ONE` ID of a channel

`DISCORD_CHANNEL_ID_TWO` ID of another channel

**By default the script supports two channels, but you can add more by adding more DISCORD_CHANNEL_ID_X in the .env file and adding an `os.Getenv("DISCORD_CHANNEL_ID_X")` in the append function line 20**