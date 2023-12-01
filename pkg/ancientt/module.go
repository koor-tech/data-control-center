package ancientt

import "go.uber.org/fx"

var Module = fx.Module("ancientt",
	fx.Provide(
		New,
	),
)
