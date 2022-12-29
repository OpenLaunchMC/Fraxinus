package errors

import (
	jsoniter "github.com/json-iterator/go"
)

type YggdrasilErrorResponse struct {
	StatusCode	 int
	Error  		 string
	ErrorMessage string
	Cause 		 string
}

func (r YggdrasilErrorResponse) JSON() []byte {
	str , err := jsoniter.Marshal(r)
	if err != nil {
		panic(err)
	}
	return str
}

func (r YggdrasilErrorResponse) JSONString() string {
	str , err := jsoniter.MarshalToString(r)
	if err != nil {
		panic(err)
	}
	return str
}

func (r YggdrasilErrorResponse) String() string {
	return r.ErrorMessage
}