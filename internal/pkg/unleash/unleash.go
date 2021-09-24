package unleash

import (
	"fmt"
	"net/http"

	"github.com/Unleash/unleash-client-go/v3"
)

type Unleash interface {
	Example_unleash()
	IsEnabled()
}

func Example_unleash() {
	unleash.Initialize(
		unleash.WithListener(&unleash.DebugListener{}),
		unleash.WithAppName("Example-go"),
		unleash.WithUrl("http://localhost:4243/api/"),
		unleash.WithCustomHeaders(http.Header{"Authorization": {"13dc681a3d730398b955e4a9b32392e7f4ababad0febd16caa1e94c37dab8366"}}),
	)
}

func IsEnabled() bool {
	if unleash.IsEnabled("Test") {
		fmt.Printf("Está activo el feature")
		return true
	}
	return false
}
