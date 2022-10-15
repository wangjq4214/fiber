package client

import (
	"net"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/require"
	"github.com/valyala/fasthttp/fasthttputil"
)

func createHelperServer(t *testing.T) (*fiber.App, func(addr string) (net.Conn, error), func()) {
	t.Helper()

	ln := fasthttputil.NewInmemoryListener()
	app := fiber.New()

	return app, func(addr string) (net.Conn, error) {
			return ln.Dial()
		}, func() {
			require.Nil(t, app.Listener(ln, fiber.ListenConfig{DisableStartupMessage: true}))
		}
}

func testAgent(t *testing.T, handler fiber.Handler, wrapAgent func(agent *Request), excepted string, count ...int) {
	t.Parallel()

	app, ln, start := createHelperServer(t)
	app.Get("/", handler)
	go start()

	c := 1
	if len(count) > 0 {
		c = count[0]
	}

	for i := 0; i < c; i++ {
		req := AcquireRequest().SetDial(ln)
		wrapAgent(req)

		resp, err := req.Get("http://example.com")

		require.NoError(t, err)
		require.Equal(t, fiber.StatusOK, resp.StatusCode())
		require.Equal(t, excepted, resp.String())
		resp.Close()
	}
}

func testAgentFail(t *testing.T, handler fiber.Handler, wrapAgent func(agent *Request), excepted error, count ...int) {
	t.Parallel()

	app, ln, start := createHelperServer(t)
	app.Get("/", handler)
	go start()

	c := 1
	if len(count) > 0 {
		c = count[0]
	}

	for i := 0; i < c; i++ {
		req := AcquireRequest().SetDial(ln)
		wrapAgent(req)

		_, err := req.Get("http://example.com")

		require.Equal(t, excepted.Error(), err.Error())
	}
}