package calculateP

type Numbers struct {
  First int `json:"a"`
  Second int `json:"b"`}
type NumbersRec struct {
  First int `json:"!a"`
  Second int `json:"!b"`}
type Errorstruct struct {
  Error string `json:"error"`}
