package main
import (
  "net/http"
  "bufio"
  "os"
  "path/filepath"
  "time"
  "fmt"
  "strings"

  "github.com/jasonlvhit/gocron"

  "./config"
)

func check (e error) {
  if e != nil {
    panic(e)
  }
}

func DoShowRecording(streamURL string, duration string, savePath string) {
  // Crude workaround until gocron is fixed
  if !strings.Contains(savePath, time.Now().Weekday().String()) {
    return
  }

  os.MkdirAll(savePath, 0777)

  fname := time.Now().Format("Jan-2-2006") + filepath.Ext(streamURL)
  f, err := os.Create(filepath.Join(savePath, fname))
  check(err)
  defer f.Close()

  resp, err := http.Get(streamURL)
  check(err)

  startTime := time.Now().UnixNano()
  durationObject, e := time.ParseDuration(duration)
  check(e)
  endTime := startTime + durationObject.Nanoseconds()

  fmt.Printf("Beginning recording, lasting for %v. Saving to %v\n",
    duration, savePath)
  reader := bufio.NewReader(resp.Body)
  for time.Now().UnixNano() < endTime {
    line, err := reader.ReadBytes('\n')
    check(err)

    f.Write(line)
  }
  fmt.Printf("Ending recording of %v\n", savePath)
  f.Sync()
}

func ScheduleShowRecordings(cfg cfgparser.Config) {
  mon := gocron.NewScheduler()
  for _,show := range cfg.Shows.Monday {
    showDir := filepath.Join(cfg.BaseSavePath, "Monday", show.StartTime)
    fmt.Println("Scheduling show on Monday at", show.StartTime)
    mon.Every(1).Monday().At(show.StartTime).Do(DoShowRecording, cfg.StreamURL,
      show.Duration, showDir)
  }

  tue := gocron.NewScheduler()
  for _,show := range cfg.Shows.Tuesday {
    showDir := filepath.Join(cfg.BaseSavePath, "Tuesday", show.StartTime)
    fmt.Println("Scheduling show on Tuesday at", show.StartTime)
    tue.Every(1).Tuesday().At(show.StartTime).Do(DoShowRecording, cfg.StreamURL,
      show.Duration, showDir)
  }

  wed := gocron.NewScheduler()
  for _,show := range cfg.Shows.Wednesday {
    showDir := filepath.Join(cfg.BaseSavePath, "Wednesday", show.StartTime)
    fmt.Println("Scheduling show on Wednesday at", show.StartTime)
    wed.Every(1).Wednesday().At(show.StartTime).Do(DoShowRecording, cfg.StreamURL,
      show.Duration, showDir)
  }

  thu := gocron.NewScheduler()
  for _,show := range cfg.Shows.Thursday {
    showDir := filepath.Join(cfg.BaseSavePath, "Thursday", show.StartTime)
    fmt.Println("Scheduling show on Thursday at", show.StartTime)
    thu.Every(1).Thursday().At(show.StartTime).Do(DoShowRecording, cfg.StreamURL,
      show.Duration, showDir)
  }

  fri := gocron.NewScheduler()
  for _,show := range cfg.Shows.Friday {
    showDir := filepath.Join(cfg.BaseSavePath, "Friday", show.StartTime)
    fmt.Println("Scheduling show on Friday at", show.StartTime)
    fri.Every(1).Friday().At(show.StartTime).Do(DoShowRecording, cfg.StreamURL,
      show.Duration, showDir)
  }

  sat := gocron.NewScheduler()
  for _,show := range cfg.Shows.Saturday {
    showDir := filepath.Join(cfg.BaseSavePath, "Saturday", show.StartTime)
    fmt.Println("Scheduling show on Saturday at", show.StartTime)
    sat.Every(1).Saturday().At(show.StartTime).Do(DoShowRecording, cfg.StreamURL,
      show.Duration, showDir)
  }

  sun := gocron.NewScheduler()
  for _,show := range cfg.Shows.Sunday {
    showDir := filepath.Join(cfg.BaseSavePath, "Sunday", show.StartTime)
    fmt.Println("Scheduling show on Sunday at", show.StartTime)
    sun.Every(1).Sunday().At(show.StartTime).Do(DoShowRecording, cfg.StreamURL,
      show.Duration, showDir)
  }

  go mon.Start()
  go tue.Start()
  go wed.Start()
  go thu.Start()
  go fri.Start()
  go sat.Start()
  go sun.Start()
}

func main() {
  cfg := cfgparser.GetConfig()
  fmt.Println("Hypatia's live!")
  ScheduleShowRecordings(cfg)
  fmt.Println(`Successfully scheduled recordings. You should now minimize this window. Be careful not to close it.`)
  select {}
}

