package controller

import (
	"github.com/CoderI421/gframework/pkg/log"
	"go.opentelemetry.io/otel/trace"

	"github.com/gin-gonic/gin"
)

func (us *userServer) List(ctx *gin.Context) {

	span := trace.SpanFromContext(ctx.Request.Context())
	log.Info(span.SpanContext().TraceID().String())
	log.Info("GetUserList is called")
}
