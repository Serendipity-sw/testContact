package main

import (
	"fmt"
	"math"
	"strconv"
)

//实体
type NetWorkIp struct {
	Ip_1         int //传入的IP
	Ip_2         int
	Ip_3         int
	Ip_4         int
	Bits         int //掩码位
	NumOfAddress int //可用地址数
	Snm_1        int //子网掩码
	Snm_2        int
	Snm_3        int
	Snm_4        int
	BCast_1      int //广播
	BCast_2      int
	BCast_3      int
	BCast_4      int
	NWAdr_1      int //网络
	NWAdr_2      int
	NWAdr_3      int
	NWAdr_4      int
	FirstAdr_1   int //第一个可用
	FirstAdr_2   int
	FirstAdr_3   int
	FirstAdr_4   int
	LastAdr_1    int //最后一个可用
	LastAdr_2    int
	LastAdr_3    int
	LastAdr_4    int
}

func main() {
	nwip := &NetWorkIp{
		Ip_1: 192,
		Ip_2: 168,
		Ip_3: 0,
		Ip_4: 0,
		Bits: 32,
	}
	model, err := calNBFL(*nwip)
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Printf("IP：%s.%s.%s.%s \n",strconv.Itoa(model.Ip_1),strconv.Itoa(model.Ip_2),strconv.Itoa(model.Ip_3),strconv.Itoa(model.Ip_4))
	fmt.Printf("掩码位：%s \n",strconv.Itoa(model.Bits))
	fmt.Printf("可用地址：%s \n",strconv.Itoa(model.NumOfAddress))
	fmt.Printf("子网掩码：%s.%s.%s.%s \n",strconv.Itoa(model.Snm_1),strconv.Itoa(model.Snm_2),strconv.Itoa(model.Snm_3),strconv.Itoa(model.Snm_4))
	fmt.Printf("网络：%s.%s.%s.%s \n",strconv.Itoa(model.NWAdr_1),strconv.Itoa(model.NWAdr_2),strconv.Itoa(model.NWAdr_3),strconv.Itoa(model.NWAdr_4))
	fmt.Printf("第一个可用：%s.%s.%s.%s \n",strconv.Itoa(model.FirstAdr_1),strconv.Itoa(model.FirstAdr_2),strconv.Itoa(model.FirstAdr_3),strconv.Itoa(model.FirstAdr_4))
	fmt.Printf("最后可用：%s.%s.%s.%s \n",strconv.Itoa(model.LastAdr_1),strconv.Itoa(model.LastAdr_2),strconv.Itoa(model.LastAdr_3),strconv.Itoa(model.LastAdr_4))
	fmt.Printf("广播：%s.%s.%s.%s \n",strconv.Itoa(model.BCast_1),strconv.Itoa(model.BCast_2),strconv.Itoa(model.BCast_3),strconv.Itoa(model.BCast_4))
}

func calNBFL(nwip NetWorkIp) (NetWorkIp, error) {
	reset_rest_from4(nwip)
	if nwip.Ip_1 > 255 || nwip.Ip_1 < 0 || nwip.Ip_2 > 255 || nwip.Ip_2 < 0 || nwip.Ip_3 > 255 || nwip.Ip_3 < 0 || nwip.Ip_4 > 255 || nwip.Ip_4 < 0 {
		nwip.NumOfAddress = 0
		return nwip, nil
	}
	nwip, err := calcNWmask(nwip)
	if err != nil {
		return nwip, err
	}
	if nwip.Bits < 0 || nwip.Bits > 32 {
		nwip.NumOfAddress = 0
		return nwip, nil
	}
	if nwip.Bits == 31 {
		nwip.NumOfAddress = 2
		nwip.FirstAdr_1 = nwip.Ip_1 & nwip.Snm_1
		nwip.FirstAdr_2 = nwip.Ip_2 & nwip.Snm_2
		nwip.FirstAdr_3 = nwip.Ip_3 & nwip.Snm_3
		nwip.FirstAdr_4 = nwip.Ip_4 & nwip.Snm_4

		nwip.LastAdr_1 = nwip.Ip_1 | (^nwip.Snm_1 & 0xff)
		nwip.LastAdr_2 = nwip.Ip_2 | (^nwip.Snm_2 & 0xff)
		nwip.LastAdr_3 = nwip.Ip_3 | (^nwip.Snm_3 & 0xff)
		nwip.LastAdr_4 = nwip.Ip_4 | (^nwip.Snm_4 & 0xff)
		return nwip, nil
	}
	if nwip.Bits == 32 {
		nwip.NumOfAddress = 1
		nwip.FirstAdr_1 = nwip.Ip_1
		nwip.FirstAdr_2 = nwip.Ip_2
		nwip.FirstAdr_3 = nwip.Ip_3
		nwip.FirstAdr_4 = nwip.Ip_4
		return nwip, nil
	}
	nwip.NumOfAddress =int(math.Pow(2, float64(32 - nwip.Bits))) - 2

	nwip.BCast_1 = nwip.Ip_1 | (^nwip.Snm_1 & 0xff)
	nwip.BCast_2 = nwip.Ip_2 | (^nwip.Snm_2 & 0xff)
	nwip.BCast_3 = nwip.Ip_3 | (^nwip.Snm_3 & 0xff)
	nwip.BCast_4 = nwip.Ip_4 | (^nwip.Snm_4 & 0xff)

	nwip.NWAdr_1 = nwip.Ip_1 & nwip.Snm_1
	nwip.NWAdr_2 = nwip.Ip_2 & nwip.Snm_2
	nwip.NWAdr_3 = nwip.Ip_3 & nwip.Snm_3
	nwip.NWAdr_4 = nwip.Ip_4 & nwip.Snm_4

	nwip.FirstAdr_1 = nwip.NWAdr_1
	nwip.FirstAdr_2 = nwip.NWAdr_2
	nwip.FirstAdr_3 = nwip.NWAdr_3
	nwip.FirstAdr_4 = nwip.NWAdr_4 + 1

	nwip.LastAdr_1 = nwip.BCast_1
	nwip.LastAdr_2 = nwip.BCast_2
	nwip.LastAdr_3 = nwip.BCast_3
	nwip.LastAdr_4 = nwip.BCast_4 - 1
	return nwip, nil
}

//计算子网掩码
func calcNWmask(nwip NetWorkIp) (NetWorkIp, error) {
	tmpvar :=nwip.Bits
	if tmpvar > 32 || tmpvar < 0 {
		nwip.Snm_1 = 0
		nwip.Snm_2 = 0
		nwip.Snm_3 = 0
		nwip.Snm_4 = 0
		return nwip, nil
	}
	nwip.Snm_1 = 0
	nwip.Snm_2 = 0
	nwip.Snm_3 = 0
	nwip.Snm_4 = 0
	if tmpvar >= 8 {
		nwip.Snm_1 = 255
		tmpvar -= 8
	} else {
		nwip.Snm_1 = h_fillbitsfromleft(tmpvar)
		return nwip, nil
	}
	if tmpvar >= 8 {
		nwip.Snm_2 = 255
		tmpvar -= 8
	} else {
		nwip.Snm_2 = h_fillbitsfromleft(tmpvar)
		return nwip, nil
	}
	if tmpvar >= 8 {
		nwip.Snm_3 = 255
		tmpvar -= 8
	} else {
		nwip.Snm_3 = h_fillbitsfromleft(tmpvar)
		return nwip, nil
	}
	nwip.Snm_4 = h_fillbitsfromleft(tmpvar)
	return nwip, nil
}

func h_fillbitsfromleft(num int) int {
	if num >= 8 {
		return 255
	}
	bitpat := 0xff00
	for {
		if num > 0 {
			bitpat = bitpat >> 1
			num--
		}else{
			break
		}
	}
	return (bitpat & 0xff)
}

func reset_rest_from4(nwip NetWorkIp) {
	nwip.BCast_1 = 0
	nwip.BCast_2 = 0
	nwip.BCast_3 = 0
	nwip.BCast_4 = 0

	nwip.NWAdr_1 = 0
	nwip.NWAdr_2 = 0
	nwip.NWAdr_3 = 0
	nwip.NWAdr_4 = 0

	nwip.FirstAdr_1 = 0
	nwip.FirstAdr_2 = 0
	nwip.FirstAdr_3 = 0
	nwip.FirstAdr_4 = 0

	nwip.LastAdr_1 = 0
	nwip.LastAdr_2 = 0
	nwip.LastAdr_3 = 0
	nwip.LastAdr_4 = 0

	nwip.Snm_1 = 0
	nwip.Snm_2 = 0
	nwip.Snm_3 = 0
	nwip.Snm_4 = 0

	nwip.NumOfAddress = 0
}
