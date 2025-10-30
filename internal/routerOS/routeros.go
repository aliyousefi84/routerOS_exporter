package routeros

import (
	"fmt"
	"strconv"

	"github.com/aliyousefi84/routerOS_exporter/internal/prometheus"
	"github.com/go-routeros/routeros/v3"
)



type MikSvc struct {
	conn *routeros.Client
	Prom *prometheus.Collect
}



func  Initialize() (*MikSvc , error ) {
	conn , err := routeros.Dial("192.168.100.88:8728" , "ali" , "Ali@1384")
	if err != nil {
		return nil , fmt.Errorf("error , problem to connect mikrotik routerOS")
	}
	fmt.Println("connection to routerOS successful")

	return &MikSvc{
		conn: conn,
	}, nil
}


// return float 64 for cpu gauge metric
func (m *MikSvc) GetCpu ()  {
	reply , err := m.conn.Run("/system/resource/cpu/print")
	if err != nil {
		fmt.Println("error , invalid command")
		return
	}

	Data := reply.Re[0].Map["load"]

	load , _:= strconv.ParseFloat(Data , 64)

	m.Prom.SetRouterCpuLoad(load)
	
}



func (m *MikSvc) GetFreeMem () {
	reply , err := m.conn.Run("/system/resource/print")

	if err != nil {
		fmt.Println("invalid command")
		return
	}

	for _ , re := range reply.Re {
		fmt.Printf("%v\n" , re.Map["free-memory"])
	}
}




func (m *MikSvc) GetFreeSpace () {
	reply , err := m.conn.Run("/system/resource/print")

	if err != nil {
		fmt.Println("invalid command")
		return
	}

	for _ , re := range reply.Re {
		fmt.Printf("%v\n" , re.Map["free-hdd-space"])
	
	}

}



