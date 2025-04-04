# Dahua loader
This application is used to download videos from Dahua cameras.
The videos are then converted using ffmpeg to the *.mp4 format with stream copy enabled, which doesn't put load on the CPU.
The converted videos are uploaded to the converted folder.
You can define your own settings for the application using environment variables or by editing the .env file.

# Web control
Using the Web control panel, you can add or remove cameras.
To add a new user, use the command:
./dahua-loader-console add-admin

# Compilation
Compilation is done by running ./build.sh.
As a result, two files are generated: dahua-loader-console and dahua-loader-server.
