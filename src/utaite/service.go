package utaite

import "fmt"

type IntIDRequest struct {
	ID int64 `json:"id,string" swagger:",in=query"`
}

type StringIDRequest struct {
	ID string `json:"id" swagger:",in=query"`
}

type ListOpts struct {
	Offset int `json:"offset" swagger:",in=query"`
	Limit  int `json:"limit" swagger:",in=query"`
}

type HttpError struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

func (err *HttpError) Error() string {
	return fmt.Sprintf("status %d: %s", err.Code, err.Text)
}

func (err *HttpError) StatusCode() int {
	return err.Code
}

func (err *HttpError) ErrorMessage() interface{} {
	return err
}
