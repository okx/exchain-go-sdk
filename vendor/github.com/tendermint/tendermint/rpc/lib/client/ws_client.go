package rpcclient

import (
	"net"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/ok-chain/ok-gosdk/common/libs/go-metrics"

	"github.com/tendermint/go-amino"
	cmn "github.com/tendermint/tendermint/libs/common"
	types "github.com/tendermint/tendermint/rpc/lib/types"
)


// WSClient is a WebSocket client. The methods of WSClient are safe for use by
// multiple goroutines.
type WSClient struct {
	cmn.BaseService

	conn *websocket.Conn
	cdc  *amino.Codec

	Address  string // IP:PORT or /path/to/socket
	Endpoint string // /websocket/url/endpoint
	Dialer   func(string, string) (net.Conn, error)

	// Time between sending a ping and receiving a pong. See
	// https://godoc.org/github.com/rcrowley/go-metrics#Timer.
	PingPongLatencyTimer metrics.Timer

	// Single user facing channel to read RPCResponses from, closed only when the client is being stopped.
	ResponsesCh chan types.RPCResponse

	// Callback, which will be called each time after successful reconnect.
	onReconnect func()

	// internal channels
	send            chan types.RPCRequest // user requests
	backlog         chan types.RPCRequest // stores a single user request received during a conn failure
	reconnectAfter  chan error            // reconnect requests
	readRoutineQuit chan struct{}         // a way for readRoutine to close writeRoutine

	wg sync.WaitGroup

	mtx            sync.RWMutex
	sentLastPingAt time.Time
	reconnecting   bool

	// Maximum reconnect attempts (0 or greater; default: 25).
	maxReconnectAttempts int

	// Time allowed to write a message to the server. 0 means block until operation succeeds.
	writeWait time.Duration

	// Time allowed to read the next message from the server. 0 means block until operation succeeds.
	readWait time.Duration

	// Send pings to server with this period. Must be less than readWait. If 0, no pings will be sent.
	pingPeriod time.Duration

	// Support both ws and wss protocols
	protocol string
}

