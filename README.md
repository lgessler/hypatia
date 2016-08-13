<img src="http://i.imgur.com/6ogRder.jpg" width=250></img>
# Hypatia

Hypatia listens to HTTP streaming response bodies on (1) a day of the week (2)
at a certain time (3) for as long you tell her to.

Or, Hypatia lets you record your radio shows!

# Usage

```yaml
StreamURL: "http://wtju.net:8000/wtju-opus-256.ogg"
BaseSavePath: "."
Shows:
  Monday:
    - StartTime: "08:00"
      Duration: "2h"
    - StartTime: "06:00"
      Duration: "1m1h"
    # ...
  Saturday:
    - StartTime: "12:10"
      Duration: "1m"
```

Edit `config.yaml` to your liking. In this sample config, we've told Hypatia to
record WXTJ's internet stream, save them in the directory we're running Hypatia
from, and all the times when we'd like her to listen.

# TODO

* Compile and distribute binaries
* Automatic MP3 conversion
* More forgiving time format
