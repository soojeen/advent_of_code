package utils

import "io/ioutil"

func ReadInput(filename string) (string, error) {
  content, err := ioutil.ReadFile(filename)

  if err != nil {
    return "", err
  }

  return string(content), err
}
