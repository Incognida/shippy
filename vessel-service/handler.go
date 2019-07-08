package main

import (
	"context"
	vp "github.com/Incognida/shippy_protos/vessel"
)

type handler struct {
	repository
}

func (h *handler) FindAvailable(ctx context.Context, spec *vp.Specification, resp *vp.Response) error {
	vessel, err := h.repository.FindAvailable(spec)
	if err != nil {
		return nil
	}

	resp.Vessel = vessel
	return nil
}

func (h *handler) Create(ctx context.Context, vessel *vp.Vessel, resp *vp.Response) error {
	err := h.repository.Create(vessel)
	if err != nil {
		return err
	}

	resp.Vessel = vessel
	resp.Created = true
	return nil
}
