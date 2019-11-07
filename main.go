package main

import (
  restservice "SpaceApp/conf"
)

func main() {
  restservice.Init_restfulAPI_service("HTTP")
  //restservice.Init_restfulAPI_service("HTTPS")
}
