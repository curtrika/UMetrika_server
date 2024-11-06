package admin_panel

import "log/slog"

type AdminPanel struct {
	log      *slog.Logger
	provider Provider
}

type Provider interface {
	//TODO: сюда воообще все методы напишем
}

func New(
	log *slog.Logger,
	provider Provider,
) *AdminPanel {
	return &AdminPanel{
		log:      log,
		provider: provider,
	}
}
