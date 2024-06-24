package task

import "encoding/json"

type Status int

const (
  StatusIncomplete Status = 0
  StatusComplete   Status = 1
)

func (s *Status) UnmarshalJSON(input []byte) error {
  var status int
  if err := json.Unmarshal(input, &status); err != nil {
    return err
  }

  switch Status(status) {
  case StatusComplete:
    *s = Status(status)
  default:
    *s = StatusIncomplete
  }
  return nil
}