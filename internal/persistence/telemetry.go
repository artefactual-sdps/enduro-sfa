package persistence

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"github.com/artefactual-sdps/enduro/internal/datatypes"
	"github.com/artefactual-sdps/enduro/internal/telemetry"
)

type wrapper struct {
	wrapped Service
	tracer  trace.Tracer
}

var _ Service = (*wrapper)(nil)

// WithTelemetry enriches Service by adding instrumentation and context.
func WithTelemetry(wrapped Service, tracer trace.Tracer) *wrapper {
	return &wrapper{wrapped, tracer}
}

func updateError(err error, name string) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf("%s: %w", name, err)
}

func (w *wrapper) CreatePackage(ctx context.Context, p *datatypes.Package) error {
	ctx, span := w.tracer.Start(ctx, "CreatePackage")
	defer span.End()

	err := w.wrapped.CreatePackage(ctx, p)
	if err != nil {
		telemetry.RecordError(span, err)
		return updateError(err, "CreatePackage")
	}

	return nil
}

func (w *wrapper) UpdatePackage(ctx context.Context, id uint, updater PackageUpdater) (*datatypes.Package, error) {
	ctx, span := w.tracer.Start(ctx, "UpdatePackage")
	defer span.End()
	span.SetAttributes(attribute.Int("id", int(id)))

	r, err := w.wrapped.UpdatePackage(ctx, id, updater)
	if err != nil {
		telemetry.RecordError(span, err)
		return nil, updateError(err, "UpdatePackage")
	}

	return r, nil
}

func (w *wrapper) CreatePreservationTask(ctx context.Context, pt *datatypes.PreservationTask) error {
	ctx, span := w.tracer.Start(ctx, "CreatePreservationTask")
	defer span.End()

	err := w.wrapped.CreatePreservationTask(ctx, pt)
	if err != nil {
		telemetry.RecordError(span, err)
		return updateError(err, "CreatePreservationTask")
	}

	return nil
}

func (w *wrapper) UpdatePreservationTask(
	ctx context.Context,
	id uint,
	updater PresTaskUpdater,
) (*datatypes.PreservationTask, error) {
	ctx, span := w.tracer.Start(ctx, "UpdatePreservationTask")
	defer span.End()
	span.SetAttributes(attribute.Int("id", int(id)))

	r, err := w.wrapped.UpdatePreservationTask(ctx, id, updater)
	if err != nil {
		telemetry.RecordError(span, err)
		return nil, updateError(err, "UpdatePreservationTask")
	}

	return r, nil
}
