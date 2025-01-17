// Copyright 2019-2022 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package gtpv1

import (
	"errors"
	"fmt"
)

var (
	// ErrUnexpectedType indicates that the type of incoming message is not expected.
	ErrUnexpectedType = errors.New("got unexpected type of message")

	// ErrInvalidConnection indicates that the connection type(C-Plane or U-Plane) is
	// not the expected one.
	ErrInvalidConnection = errors.New("got invalid connection type")

	// ErrConnNotOpened indicates that some operation is failed due to the status of
	// Conn is not valid.
	ErrConnNotOpened = errors.New("connection is not opened")
)

// ErrorIndicatedError indicates that Error Indication message is received on U-Plane Connection.
type ErrorIndicatedError struct {
	TEID uint32
	Peer string
}

func (e *ErrorIndicatedError) Error() string {
	return fmt.Sprintf("error received from %s, TEIDDataI: %#x", e.Peer, e.TEID)
}

// HandlerNotFoundError indicates that the handler func is not registered in *Conn
// for the incoming GTPv2 message. In usual cases this error should not be taken
// as fatal, as the other endpoint can make your program stop working just by
// sending unregistered message.
type HandlerNotFoundError struct {
	MsgType string
}

// Error returns violating message type to handle.
func (e *HandlerNotFoundError) Error() string {
	return fmt.Sprintf("no handlers found for incoming message: %s, ignoring", e.MsgType)
}
