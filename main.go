package gotracert

import (
	ping "github.com/t0stbrot/go-ping"
)

type HopEntry struct {
	Address string `json:"address,omitempty"`
	RTT     string `json:"rtt,omitempty"`
	Error   string `json:"error,omitempty"`
}

type TracerouteResult struct {
	Target  string     `json:"target"`
	MaxHops int        `json:"maxhops"`
	Hops    []HopEntry `json:"hops"`
	Message string     `json:"message,omitempty"`
	Error   string     `json:"error,omitempty"`
}

func Traceroute4(target string, maxhops int, timeout int) (result TracerouteResult) {
	ttl := 1
	res := TracerouteResult{Target: target, MaxHops: maxhops}
	for ttl <= maxhops {
		cur := ping.Ping4(target, ttl, timeout)
		ttl++
		if cur.Message == "suceed" && cur.LastHop == target {
			resEntry := HopEntry{Address: cur.LastHop, RTT: cur.RTT}
			res.Hops = append(res.Hops, resEntry)
			break
		} else if cur.Message == "timeexceed" {
			resEntry := HopEntry{Address: cur.LastHop, RTT: cur.RTT}
			res.Hops = append(res.Hops, resEntry)
		} else if cur.Error != "" {
			resEntry := HopEntry{Address: cur.LastHop, RTT: cur.RTT, Error: cur.Error}
			res.Hops = append(res.Hops, resEntry)
		} else {
			resEntry := HopEntry{Address: cur.LastHop, RTT: cur.RTT, Error: cur.Error}
			res.Hops = append(res.Hops, resEntry)
		}
	}
	return res
}

func Traceroute6(target string, maxhops int, timeout int) (result TracerouteResult) {
	ttl := 1
	res := TracerouteResult{Target: target, MaxHops: maxhops}
	for ttl <= maxhops {
		cur := ping.Ping6(target, ttl, timeout)
		ttl++
		if cur.LastHop == target {
			resEntry := HopEntry{Address: cur.LastHop, RTT: cur.RTT}
			res.Hops = append(res.Hops, resEntry)
			break
		} else if cur.Message == "suceed" {
			resEntry := HopEntry{Address: cur.LastHop, RTT: cur.RTT}
			res.Hops = append(res.Hops, resEntry)
		} else if cur.Message == "timeexceed" {
			resEntry := HopEntry{Address: cur.LastHop, RTT: cur.RTT}
			res.Hops = append(res.Hops, resEntry)
		} else if cur.Error != "" {
			resEntry := HopEntry{Address: cur.LastHop, RTT: cur.RTT, Error: cur.Error}
			res.Hops = append(res.Hops, resEntry)
		}
	}
	return res
}
