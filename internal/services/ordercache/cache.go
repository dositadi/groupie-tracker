package ordercache

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"
)

type Ticket string
type TicketAmount float64
type Fee float64
type Vat int

const (
	GENERAL  Ticket = Ticket("General Admission")
	VIP      Ticket = Ticket("VIP")
	RESERVED Ticket = Ticket("Reserved")
)

const (
	GENERAL_AMT  TicketAmount = TicketAmount(49)
	VIP_AMT      TicketAmount = TicketAmount(120)
	RESERVED_AMT TicketAmount = TicketAmount(85)

	BOOKING_FEE Fee = Fee(2.50)

	VAT Vat = Vat(19)
)

var ttl = 15 * time.Minute

type Booking struct {
	TicketType string
	Quantity   int
	createdAt  time.Time
	expiryTime time.Time
}

type Cache struct {
	mu    sync.RWMutex
	store map[string]Booking
}

var global = &Cache{
	store: make(map[string]Booking),
}

func key(userId, location string, artistId int) string {
	location = strings.ReplaceAll(location, " ", "-")
	return fmt.Sprintf("%s-%v-%s", userId, artistId, location)
}

func Init(ctx context.Context) {
	go global.evictLoop(ctx)
}

func (c *Cache) evictLoop(ctx context.Context) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			now := time.Now()
			c.mu.Lock()
			for key := range c.store {
				if now.After(c.store[key].expiryTime) {
					delete(c.store, key)
				}
			}
			c.mu.Unlock()
		}
	}
}

func Set(userId, location string, artistId int, tickettype string) {
	key := key(userId, location, artistId)

	global.mu.Lock()
	global.store[key] = Booking{
		TicketType: tickettype,
		Quantity:   1,
		createdAt:  time.Now(),
		expiryTime: time.Now().Add(ttl),
	}
	global.mu.Unlock()
}

func Get(userId, location string, artistId int) (Booking, bool) {
	key := key(userId, location, artistId)

	global.mu.RLock()
	booking, ok := global.store[key]
	defer global.mu.RUnlock()
	return booking, ok
}

func Increment(userId, location string, artistId int) bool {
	key := key(userId, location, artistId)

	global.mu.Lock()
	defer global.mu.Unlock()
	booking, ok := global.store[key]
	if !ok || booking.Quantity <= 1 {
		return false
	}
	booking.Quantity++
	booking.expiryTime.Add(ttl)
	global.store[key] = booking
	return ok
}

func Decrement(userId, location string, artistId int) bool {
	key := key(userId, location, artistId)

	global.mu.Lock()
	defer global.mu.Unlock()
	booking, ok := global.store[key]
	if !ok || booking.Quantity > 8 {
		return false
	}
	booking.Quantity--
	booking.expiryTime.Add(ttl)
	global.store[key] = booking
	return ok
}

func Delete(userId, location string, artistId int) {
	key := key(userId, location, artistId)

	global.mu.Lock()
	delete(global.store, key)
	global.mu.Unlock()
}
