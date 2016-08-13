package main
import (
  "net/http"
  "bufio"
  "os"
  "path/filepath"
  "time"
  "fmt"

  "github.com/jasonlvhit/gocron"

  "./config"
)

func check (e error) {
  if e != nil {
    panic(e)
  }
}

func ScheduleShow(streamURL string, duration string, savePath string) {
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

  fmt.Printf("Starting at %s, ending at %s", startTime, endTime)

  reader := bufio.NewReader(resp.Body)
  for time.Now().UnixNano() < endTime {
    line, err := reader.ReadBytes('\n')
    check(err)

    f.Write(line)
  }
  f.Sync()
}

func ScheduleShowRecordings(cfg cfgparser.Config) {
  for _,show := range cfg.Shows.Monday {
    showDir := filepath.Join(cfg.BaseSavePath, "Monday", show.StartTime)
    gocron.Every(1).Monday().At(show.StartTime).Do(ScheduleShow, cfg.StreamURL,
      show.Duration, showDir)
  }

  for _,show := range cfg.Shows.Tuesday {
    showDir := filepath.Join(cfg.BaseSavePath, "Tuesday", show.StartTime)
    gocron.Every(1).Tuesday().At(show.StartTime).Do(ScheduleShow, cfg.StreamURL,
      show.Duration, showDir)
  }

  for _,show := range cfg.Shows.Wednesday {
    showDir := filepath.Join(cfg.BaseSavePath, "Wednesday", show.StartTime)
    gocron.Every(1).Wednesday().At(show.StartTime).Do(ScheduleShow, cfg.StreamURL,
      show.Duration, showDir)
  }

  for _,show := range cfg.Shows.Thursday {
    showDir := filepath.Join(cfg.BaseSavePath, "Thursday", show.StartTime)
    gocron.Every(1).Thursday().At(show.StartTime).Do(ScheduleShow, cfg.StreamURL,
      show.Duration, showDir)
  }

  for _,show := range cfg.Shows.Friday {
    showDir := filepath.Join(cfg.BaseSavePath, "Friday", show.StartTime)
    gocron.Every(1).Friday().At(show.StartTime).Do(ScheduleShow, cfg.StreamURL,
      show.Duration, showDir)
  }

  for _,show := range cfg.Shows.Saturday {
    showDir := filepath.Join(cfg.BaseSavePath, "Saturday", show.StartTime)
    gocron.Every(1).Saturday().At(show.StartTime).Do(ScheduleShow, cfg.StreamURL,
      show.Duration, showDir)
  }

  for _,show := range cfg.Shows.Sunday {
    showDir := filepath.Join(cfg.BaseSavePath, "Sunday", show.StartTime)
    gocron.Every(1).Sunday().At(show.StartTime).Do(ScheduleShow, cfg.StreamURL,
      show.Duration, showDir)
  }
}

func main() {
  cfg := cfgparser.GetConfig()
  ScheduleShowRecordings(cfg)
  <-gocron.Start()
}

