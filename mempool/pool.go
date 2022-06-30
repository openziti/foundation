/*
	Copyright NetFoundry Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	https://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package mempool

type Pool interface {
	AcquireBuffer() *DefaultPooledBuffer
}

func NewPool(poolSize int, bufSize int) Pool {
	pool := &pool{
		bufChan: make(chan *DefaultPooledBuffer, poolSize),
	}

	pool.newBuf = func() *DefaultPooledBuffer {
		buffer := make([]byte, bufSize)
		pooled := &DefaultPooledBuffer{
			Buf: buffer,
		}
		pooled.release = func() {
			pooled.Buf = buffer
			select {
			case pool.bufChan <- pooled:
			default: // if pool is full, drop the buffer
			}
		}
		return pooled
	}

	return pool
}

type pool struct {
	bufChan chan *DefaultPooledBuffer
	newBuf  func() *DefaultPooledBuffer
}

func (p *pool) AcquireBuffer() *DefaultPooledBuffer {
	select {
	case buf := <-p.bufChan:
		return buf
	default:
		return p.newBuf()
	}
}

type PooledBuffer interface {
	GetPayload() []byte
	Release()
}

type DefaultPooledBuffer struct {
	Buf     []byte
	release func()
}

func (buffer *DefaultPooledBuffer) GetPayload() []byte {
	return buffer.Buf
}

func (buffer *DefaultPooledBuffer) Release() {
	buffer.release()
}

func NewStrictPool(poolSize int, bufSize int) Pool {
	pool := &strictPool{
		bufChan: make(chan *DefaultPooledBuffer, poolSize),
	}
	for i := 0; i < poolSize; i++ {
		buffer := make([]byte, bufSize)
		pooled := &DefaultPooledBuffer{
			Buf: buffer,
		}
		pooled.release = func() {
			pooled.Buf = buffer
			pool.bufChan <- pooled
		}
		pool.bufChan <- pooled
	}

	return pool
}

type strictPool struct {
	bufChan chan *DefaultPooledBuffer
}

func (p *strictPool) AcquireBuffer() *DefaultPooledBuffer {
	return <-p.bufChan
}
