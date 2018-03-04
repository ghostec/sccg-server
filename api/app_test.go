package api_test

import (
	"net/url"

	"github.com/ghostec/sccg-server/api"
	"github.com/gorilla/websocket"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AppHandler", func() {
	var app *api.App
	var c *websocket.Conn
	var err error

	var write = func(m string) {
		err = c.WriteMessage(websocket.TextMessage, []byte(m))
		Expect(err).NotTo(HaveOccurred())
	}

	var read = func(m string) {
		_, message, err := c.ReadMessage()
		Expect(err).NotTo(HaveOccurred())
		Expect(message).To(Equal([]byte(m)))
	}

	BeforeEach(func() {
		addr := "localhost:8080"
		app, err = api.NewApp(addr)
		Expect(err).NotTo(HaveOccurred())
		go app.ListenAndServe()

		u := url.URL{Scheme: "ws", Host: addr, Path: "/"}
		c, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		err = c.WriteMessage(
			websocket.CloseMessage,
			websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
		)
		Expect(err).NotTo(HaveOccurred())
		c.Close()
		app.Stop()
	})

	It("Example", func() {
		write("test message")
		read("test message")
	})
})
