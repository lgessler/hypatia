
# Hypatia
<img src="http://i.imgur.com/6ogRder.jpg" width=250></img>

Maybe sometimes you feel like your only listener is your mom. But don't worry,
Hypatia's always listening.

# Usage (detailed)

## Windows
1. Download Hypatia
   [here](https://github.com/lgessler/hypatia/raw/master/dist/windows/hypatia.exe)
2. Put Hypatia in a good place. `C:\Users\<yourname>\Documents\Radio Recordings\` 
   should be good.
3. Right click inside the folder, mouse over "New", and select "Text Document". Name it anything. 
4. Double click it to open it in an editor.
5. Copy and paste the following:

   ```yaml
   StreamURL: "http://wtju.net:8000/wtjx-128.mp3"
   BaseSavePath: "."
   Shows:
     Monday:
       - StartTime: "09:00"
         Duration: "1h30m"
   ```
5. **Being careful to preserve the spacing**, change the file to reflect when
   your show is. E.g., if your show was on Sundays at 4:00PM, your file would
   look like this:

   ```yaml
   StreamURL: "http://wtju.net:8000/wtjx-128.mp3"
   BaseSavePath: "."
   Shows:
     Sunday:
       - StartTime: "16:00"
         Duration: "2h"
   ```
6. Save the file (CTRL+S) and exit your editor.
7. Right click the file, click "Rename", hit CTRL+A, and type in "config.yaml"
8. Click "yes" when the warning pops up. Don't worry, we know what we're doing.
9. Hit your Windows key and type in `cmd` to open a command prompt.
10. Type in `cd "<path>"`, where `<path>` is where you put `hypatia.exe`. E.g., `cd "Documents\Radio Recordings"`
11. Type in `.\hypatia` to start Hypatia. 
12. Minimize the window, if you want, but don't close it. It needs to stay open.
13. When the time is right, Hypatia will start recording your show. It'll appear in a folder in the directory, something like `Saturday\16:00\Aug 13 2016.mp3`.

## Apple
1. Download Hypatia 
   [here](https://github.com/lgessler/hypatia/raw/master/dist/osx/hypatia) 
2. Move Hypatia somewhere safe, like under your `Documents` folder in a new folder
   called `Radio Recordings`.
3. Hit your F4 button and type in `TextEdit`. Copy the following and paste it
   in:

   ```yaml
   StreamURL: "http://wtju.net:8000/wtjx-128.mp3"
   BaseSavePath: "."
   Shows:
     Monday:
       - StartTime: "09:00"
         Duration: "1h30m"
   ```
4. **Being careful to preserve the spacing**, change the file to reflect when
   your show is. E.g., if your show was on Sundays at 4:00PM, your file would
   look like this:

   ```yaml
   StreamURL: "http://wtju.net:8000/wtjx-128.mp3"
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
11. Hit the minimize button on the top left, if you like, but do not close the window. It needs to stay open.
12. When the time is right, Hypatia will start recording your show. It'll appear in a folder in the directory, something like `Saturday\16:00\Aug 13 2016.mp3`.

# Usage

1. Download the appropriate binary from
   [here](https://github.com/lgessler/hypatia/tree/master/dist).
2. Put the binary wherever you want the recordings to be.
3. Edit `config.yaml` to your liking. In this sample config, we've told Hypatia to
   record WXTJ's internet stream, save them in the directory we're running Hypatia
   from, and all the times when we'd like her to listen.

   If you prefer MP3, then set StreamURL to `...wtjx-128.mp3` rather than the
   provided. Note, however, that `128` is really low quality. If you want a
   higher quality MP3, you'd better listen to the `ogg` stream and transcode.

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
4. *(optional)* If you are on a UNIX-like system and you have `ffmpeg` installed, Hypatia will automatically try to transcode the recording. This is for two reasons: one, for WXTJ, the higher-quality stream is OGG 256 (as opposed to MP3 128) but MP3 is a much more convenient format; two, reading directly from an HTTP stream gets the audio itself just fine but doesn't result in correct metadata on the file.

# TODO

* More forgiving time format
* Look into other solutions to fixing metadata (like length) that don't require transcoding

# License

GPL
