package routeros

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/aliyousefi84/routerOS_exporter/internal/prometheus"
	"github.com/go-routeros/routeros/v3"
)

type MikSvc struct {
	conn *routeros.Client
	Prom *prometheus.Collect
	wg   sync.WaitGroup
}

func Initialize(addr, user, pass string) (*MikSvc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	conn, err := routeros.DialContext(ctx, addr, user, pass)
	defer cancel()
	if err != nil {
		return nil, fmt.Errorf("error , problem to connect mikrotik routerOS")
	}
	fmt.Println("connection to routerOS successful")

	return &MikSvc{
		conn: conn,
	}, nil
}

func (m *MikSvc) GetCpu(ctx context.Context) {
	reply, err := m.conn.RunContext(ctx, "/system/resource/cpu/print")

	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			fmt.Println("problem to get metrics from cpu")
			return
		}
	}
	Data := reply.Re[0].Map["load"]

	load, err := strconv.ParseFloat(Data, 64)

	if err != nil {
		fmt.Println("can't parse this data")
	}

	fmt.Printf("your cpu load is %f\n", load)

	m.Prom.SetRouterCpuLoad(load)

}

func (m *MikSvc) GetFreeMem(ctx context.Context) {
	reply, err := m.conn.RunContext(ctx, "/system/resource/print")

	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			fmt.Println("problem to get metrics from memory")
			return
		}
	}

	Data := reply.Re[0].Map["free-memory"]

	freemem, err := strconv.ParseFloat(Data, 64)

	if err != nil {
		fmt.Println("can't parse this data")
		return
	}
	fmt.Printf("your free memory is :%f\n", freemem)
	m.Prom.SetRouterFreeMem(freemem)

}

func (m *MikSvc) GetFreeSpace(ctx context.Context) {

	reply, err := m.conn.RunContext(ctx, "/system/resource/print")

	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			fmt.Println("problem to get metrics from hard space")
			return
		}
	}

	Data := reply.Re[0].Map["free-hdd-space"]

	free_space, err := strconv.ParseFloat(Data, 64)

	if err != nil {
		fmt.Println("can't parse this data")
		return
	}
	fmt.Printf("your free space is :%f\n", free_space)
	m.Prom.SetHardFreeSpace(free_space)
}

func (m *MikSvc) InetTrafikIn(ctx context.Context) {
	reply, err := m.conn.RunContext(ctx, "/interface/print")

	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			fmt.Println("problem to get metrics from interfaces")
			return
		}
	}
	for _, re := range reply.Re {
		name := re.Map["name"]
		Rx := re.Map["rx-byte"]
		Rxfloat, _ := strconv.ParseFloat(Rx, 64)
		m.wg.Add(1)
		go func() {
			defer m.wg.Done()
			m.Prom.GetTrafikIn(name, "rx-byte", Rxfloat)
		}()
		m.wg.Wait()
	}
	fmt.Println("recive input metrics from interfaces")
}

func (m *MikSvc) InetTrafikOut(ctx context.Context) {
	reply, err := m.conn.RunContext(ctx, "/interface/print")

	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			fmt.Println("problem to get metrics from interfaces")
			return
		}
	}
	for _, re := range reply.Re {
		name := re.Map["name"]
		Tx := re.Map["tx-byte"]
		Txfloat, _ := strconv.ParseFloat(Tx, 64)
		m.wg.Add(1)
		go func() {
			defer m.wg.Done()
			m.Prom.GetTrafikOut(name, "tx-byte", Txfloat)
		}()
		m.wg.Wait()
	}
	fmt.Println("recive transfer metrics from interfaces")
}

func (m *MikSvc) UserTrafik(ctx context.Context) {
	reply, err := m.conn.RunContext(ctx, "/ip/firewall/mangle/print")
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			fmt.Println("problem to get metrics from mangle roule")
			return
		}
	}

	for _, re := range reply.Re {
		name := re.Map["src-address"]
		packet := re.Map["bytes"]
		data, _ := strconv.ParseFloat(packet, 64)
		m.wg.Add(1)
		go func() {
			defer m.wg.Done()
			m.Prom.GetUserTrafik(name, data)
		}()
		m.wg.Wait()
	}
	fmt.Println("get user trafik successfuly")
}
