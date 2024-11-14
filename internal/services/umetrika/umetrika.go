package umetrika

import "log/slog"

type UMetrika struct {
	log      *slog.Logger
	provider Provider
}

type Provider interface{}

func New(
	log *slog.Logger,
	provider Provider,
) *UMetrika {
	return &UMetrika{
		log:      log,
		provider: provider,
	}
}
