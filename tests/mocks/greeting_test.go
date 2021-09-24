package main

import (
	"github.com/golang/mock/gomock"
	"testing"
)

func TestGreeting_ExecuteGreeting(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	mock := NewMockMockeableInterface(ctrl)

	// definition of behaviour
	mock.EXPECT().GetGreeting(gomock.Eq("test")).Return("Hello")

	subject := NewGreeting(mock)
	subject.ExecuteGreeting("test")

	// definition of behaviour
	mock.EXPECT().GetGreeting(gomock.Any()).Return("hello")

	subject.ExecuteGreeting("test2")
}