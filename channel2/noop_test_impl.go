package channel2

import (
	"crypto/x509"
	"github.com/openziti/foundation/identity/identity"
	"time"
)

type NoopTestChannel struct {
}

func (ch *NoopTestChannel) Id() *identity.TokenId {
	panic("implement Id()")
}

func (ch *NoopTestChannel) LogicalName() string {
	panic("implement LogicalName()")
}

func (ch *NoopTestChannel) ConnectionId() string {
	panic("implement ConnectionId()")
}

func (ch *NoopTestChannel) Certificates() []*x509.Certificate {
	panic("implement Certificates()")
}

func (ch *NoopTestChannel) Label() string {
	return "testchannel"
}

func (ch *NoopTestChannel) SetLogicalName(logicalName string) {
	panic("implement SetLogicalName")
}

func (ch *NoopTestChannel) Bind(h BindHandler) error {
	panic("implement Bind")
}

func (ch *NoopTestChannel) AddPeekHandler(h PeekHandler) {
	panic("implement AddPeekHandler")
}

func (ch *NoopTestChannel) AddTransformHandler(h TransformHandler) {
	panic("implement AddTransformHandler")
}

func (ch *NoopTestChannel) AddReceiveHandler(h ReceiveHandler) {
	panic("implement AddReceiveHandler")
}

func (ch *NoopTestChannel) AddErrorHandler(h ErrorHandler) {
	panic("implement me")
}

func (ch *NoopTestChannel) AddCloseHandler(h CloseHandler) {
	panic("implement AddErrorHandler")
}

func (ch *NoopTestChannel) SetUserData(data interface{}) {
	panic("implement SetUserData")
}

func (ch *NoopTestChannel) GetUserData() interface{} {
	panic("implement GetUserData")
}

func (ch *NoopTestChannel) Send(m *Message) error {
	return nil
}

func (ch *NoopTestChannel) SendWithPriority(m *Message, p Priority) error {
	return nil
}

func (ch *NoopTestChannel) SendAndSync(m *Message) (chan error, error) {
	return ch.SendAndSyncWithPriority(m, Standard)
}

func (ch *NoopTestChannel) SendAndSyncWithPriority(m *Message, p Priority) (chan error, error) {
	result := make(chan error, 1)
	result <- nil
	return result, nil
}

func (ch *NoopTestChannel) SendWithTimeout(m *Message, timeout time.Duration) error {
	return nil
}

func (ch *NoopTestChannel) SendPrioritizedWithTimeout(m *Message, p Priority, timeout time.Duration) error {
	return nil
}

func (ch *NoopTestChannel) SendAndWaitWithTimeout(m *Message, timeout time.Duration) (*Message, error) {
	panic("implement SendAndWaitWithTimeout")
}

func (ch *NoopTestChannel) SendPrioritizedAndWaitWithTimeout(m *Message, p Priority, timeout time.Duration) (*Message, error) {
	panic("implement SendPrioritizedAndWaitWithTimeout")
}

func (ch *NoopTestChannel) SendAndWait(m *Message) (chan *Message, error) {
	panic("implement SendAndWait")
}

func (ch *NoopTestChannel) SendAndWaitWithPriority(m *Message, p Priority) (chan *Message, error) {
	panic("implement SendAndWaitWithPriority")
}

func (ch *NoopTestChannel) SendForReply(msg TypedMessage, timeout time.Duration) (*Message, error) {
	panic("implement SendForReply")
}

func (ch *NoopTestChannel) SendForReplyAndDecode(msg TypedMessage, timeout time.Duration, result TypedMessage) error {
	return nil
}

func (ch *NoopTestChannel) Close() error {
	panic("implement Close")
}

func (ch *NoopTestChannel) IsClosed() bool {
	panic("implement IsClosed")
}

func (ch *NoopTestChannel) Underlay() Underlay {
	panic("implement Underlay")
}
