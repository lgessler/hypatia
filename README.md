<img src="http://i.imgur.com/6ogRder.jpg" width=250></img>
# Hypatia
Hypatia lets you record your radio shows *and* helps you convince yourself you
have listeners.

# Usage (Help how do I computer)

## Windows
1. Download Hypatia
   [here](https://github.com/lgessler/hypatia/raw/master/dist/windows/hypatia.exe)
2. Put Hypatia in a safe place. `C:\Users\<yourname>\Documents\Radio Recordings\` 
   should be good.
3. (Haven't finished this yet!)

## Apple
1. Download Hypatia 
   [here](https://github.com/lgessler/hypatia/raw/master/dist/osx/hypatia) 
2. Move Hypatia somewhere safe, like under your `Documents` folder in a new folder
   called `Radio Recordings`.
3. Hit your F4 button and type in `TextEdit`. Copy the following and paste it
   in:

   ```yaml
   StreamURL: "http://wtju.net:8000/wtjx-opus-256.ogg"
   BaseSavePath: "."
   Shows:
     Monday:
       - StartTime: "08:00"
         Duration: "2h"
   ```
4. **Being careful to preserve the spacing**, change the file to reflect when
   your show is. E.g., if your show was on Sundays at 4:00PM, your file would
   look like this:

   ```yaml
   StreamURL: "http://wtju.net:8000/wtjx-opus-256.ogg"
   BaseSavePath: "."
   Shows:
     Sunday:
       - StartTime: "16:00"
         Duration: "2h"
   ```
5. Hit `COMMAND+SHIFT+T` if you have a bunch of stuff on the top part of the
   screen.
6. Navigate to where you put Hypatia, and save the file (`COMMAND+S`) as 
   **config.yaml**
7. Hit F4 again and type `terminal`.
8. Type in `cd "~/Documents/Radio Recordings"`, or `cd "~/<Wherever you put
   Hypatia>"`
9. Type `chmod +x hypatia`
10. Type `./hypatia`. Now she's listening!
11. Hit the minimize button on the top left, if you like.

# Usage (I know what a CSV file is)

1. Download the appropriate binary from
   [here](https://github.com/lgessler/hypatia/tree/master/dist).
2. Put the binary wherever you want the recordings to be.
3. Edit `config.yaml` to your liking. In this sample config, we've told Hypatia to
   record WXTJ's internet stream, save them in the directory we're running Hypatia
   from, and all the times when we'd like her to listen.

   ```yaml
   StreamURL: "http://wtju.net:8000/wtjx-opus-256.ogg"
   BaseSavePath: "."
   Shows:
     Monday:
       - StartTime: "08:00" # don't forget the zero!
         Duration: "2h"
       - StartTime: "16:00" # use military time
         Duration: "1m1h"
       # ...
     Saturday:
       - StartTime: "12:10"
         Duration: "1m"
   ```

# The nerdy explanation

Hypatia listens to HTTP streaming response bodies on (1) a day of the week (2)
at a certain time (3) for as long you tell her to. 

# TODO

* Fix OGG metadata
* Automatic MP3 conversion
* More forgiving time format
