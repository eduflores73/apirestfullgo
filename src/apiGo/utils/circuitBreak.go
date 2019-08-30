package utils

import (
	"time"
)

type State string

type CircuitBreaker struct {
	state 			State
	MaxRequestsErr   int32
	Timeout       	time.Duration
	Expiry		  	time.Time
	Generation 	  	uint64
	Interval		time.Duration
//	ReadyToTrip   func(counts Counts) bool

}

const (
	Open = "Open"
	HalfOpen = "HalfOpen"
	Close = "Close"
)

func (cb *CircuitBreaker) currentState(now time.Time) (State, uint64) {
	switch cb.state {
	case Close:
		if !cb.Expiry.IsZero() && cb.Expiry.Before(now) {
			cb.toNewGeneration(now)
		}
	case Open:
		if cb.Expiry.Before(now) {
			cb.state = Open
		}
	}
	return cb.state, cb.Generation
}

func (cb *CircuitBreaker) toNewGeneration(now time.Time) {
	cb.Generation++

	var zero time.Time
	switch cb.state {
	case Close:
		if cb.Interval == 0 {
			cb.Expiry = zero
		} else {
			cb.Expiry = now.Add(cb.Interval)
		}
	case Open:
		cb.Expiry = now.Add(cb.Timeout)
	default: // StateHalfOpen
		cb.Expiry = zero
	}
}

func (cb *CircuitBreaker) State() State {
	now := time.Now()
	state, _ := cb.currentState(now)
	return state
}

func ContadorError(c int) bool{

	if(c < 4){
		return false
	}
	return true
}
