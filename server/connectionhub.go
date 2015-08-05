package entrapped

import (
	"strconv"
)

// connectionhub to manage all connections
type connectionhub struct {
	// registered connections
	troopers map[*trooper]bool
	// matches active
	matches map[*trooper]*trooper
	// register requests
	enter chan *trooper
	// unregister requests
	dead chan *trooper
	// errors
	errors chan string
	// message from requests
	message chan *message
}

var ch = &connectionhub{
	troopers: make(map[*trooper]bool),
	matches:  make(map[*trooper]*trooper),
	enter:    make(chan *trooper),
	dead:     make(chan *trooper),
	message:  make(chan *message),
	errors:   make(chan string),
}

func (ch *connectionhub) run() {
	readyState := "data:ready:" +
		"[size=" + strconv.Itoa(size) + "]:" +
		"[life=" + strconv.Itoa(lifes) + "]:"
	resultState := "data:result:"
	enemyState := "data:enemy:"

	for {
		select {
		case t := <-ch.enter:
			ch.troopers[t] = false
			for key, val := range ch.troopers {
				if !val && key != t {
					ch.matches[t] = key
					ch.matches[key] = t
					p1Tag := readyState + "[name=" + t.nickname + "]"
					key.data <- []byte(p1Tag)
					p2Tag := readyState + "[name=" + key.nickname + "]"
					t.data <- []byte(p2Tag)
					break
				}
			}
		case t := <-ch.dead:
			if _, ok := ch.troopers[t]; ok {
				delete(ch.troopers, t)
			}
			if val, ok := ch.matches[t]; ok {
				ch.troopers[val] = true
				ch.matches[val] = nil
				delete(ch.matches, t)
				close(t.data)
			}
		case m := <-ch.message:
			msg, msgErr := reqParser(m.msg)
			if len(msgErr) != 0 {
				m.t.data <- []byte(msgErr)
			} else {
				switch msg.command {
				case cmdOpen:
					id, idErr := strconv.Atoi(msg.params["idx"])
					if idErr != nil {
						m.t.data <- []byte("error:invalid idx")
					} else {
						ele, life, err := m.t.trap.open(id)
						if len(err) != 0 {
							m.t.data <- []byte(err)
						} else {
							if val, ok := ch.matches[m.t]; ok {
								m.t.data <- []byte(resultState +
									"[idx=" + msg.params["idx"] + "]:[type=" + strconv.Itoa(ele) +
									"]:[life=" + strconv.Itoa(life) + "]")
								if val != nil {
									val.data <- []byte(enemyState +
										"[idx=" + msg.params["idx"] + "]:[type=" + strconv.Itoa(ele) +
										"]:[life=" + strconv.Itoa(life) + "]")
								} else {
									m.t.data <- []byte("data:result:[status:won]")
								}
							} else {
								m.t.data <- []byte("error:no oponent")
							}
						}
					}
				default:
					m.t.data <- []byte("error:unknown command")
				}
			}
		}
	}
}

func (ch *connectionhub) add(ut *trooper) {
	ch.enter <- ut

	ut.data <- []byte("data:registered:[size=" + strconv.Itoa(size) + "]:[life=" + strconv.Itoa(lifes) + "]")

	go ut.writePump()
	ut.readPump()
}

func (ch *connectionhub) checkPlayer(id string) int{
	t := ch.troopers
	for key, _ := range t {
    	if id == key.nickname {
    		return 0
    	}
	}
	return 1
}