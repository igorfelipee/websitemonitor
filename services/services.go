package services

import (
  "fmt"
)

func MonitorarSites(){
  site := "http://www.meriti.rj.gov.br"
  res, _ := http.Get(site)

  fmt.Println("Statuscode: ", res.StatusCode)
}
