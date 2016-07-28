package main

import (
	"net/http"

	"code.cloudfoundry.org/lager"

	"github.com/tedsuo/rata"
)

func New(
	logger lager.Logger,
) http.Handler {
	startPushersHandler := NewStartPushersHandler(logger)
	pusherUpdatesHandler := NewPusherUpdatesHandler(logger)

	actions := rata.Handlers{
		StartPushersRoute: route(startPushersHandler.StartPushers),
		PostUpdateRoute:   route(pusherUpdatesHandler.PostUpdate),
	}

	handler, err := rata.NewRouter(Routes, actions)
	if err != nil {
		panic("unable to create router: " + err.Error())
	}
	return handler
}

func route(f http.HandlerFunc) http.Handler {
	return f
}
