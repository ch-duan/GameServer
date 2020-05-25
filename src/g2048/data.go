package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	m := &Model{}
	m.Init()
}

type Model struct {
	data  [4][4]int
	socre int
	end   bool
}

//作操作
func (t *Model) Left() {
	for k := 4; k > 0; k-- {
		for x := 1; x < 4; x++ {
			for y := 0; y < 4; y++ {
				if t.data[y][x-1] == t.data[y][x] || t.data[y][x-1] == 0 {
					t.data[y][x-1] += t.data[y][x]
					t.data[y][x] = 0
				}

			}
		}
	}
	t.Random()
	fmt.Println("←")
}

//右操作
func (t *Model) Right() {
	for k := 3; k > 0; k-- {
		for x := 2; x >= 0; x-- {
			for y := 0; y < 4; y++ {
				if t.data[y][x+1] == t.data[y][x] || t.data[y][x+1] == 0 {
					t.data[y][x+1] += t.data[y][x]
					t.data[y][x] = 0
				}
			}
		}
	}

	t.Random()
	fmt.Println("→")
}

//上操作
func (t *Model) Up() {
	for k := 3; k >= 0; k-- {
		for y := 1; y < 4; y++ {
			for x := 0; x < 4; x++ {
				if t.data[y-1][x] == t.data[y][x] || t.data[y-1][x] == 0 {
					t.data[y-1][x] += t.data[y][x]
					t.data[y][x] = 0
				}
			}
		}
	}
	t.Random()
	fmt.Println("↑")
}

//下操作
func (t *Model) Down() {
	for k := 3; k >= 0; k-- {
		for y := 2; y >= 0; y-- {
			for x := 0; x < 4; x++ {
				if t.data[y+1][x] == t.data[y][x] || t.data[y+1][x] == 0 {
					t.data[y+1][x] += t.data[y][x]
					t.data[y][x] = 0
				}
			}
		}
	}

	t.Random()
	fmt.Println("↓")
}

//在地图空白处随机产生一个2或4
func (t *Model) Random() {
	nilPosition := make([]int, 16)
	p := 0
	// 将地图中的空点放在一个数组中
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if t.data[i][j] == 0 {
				nilPosition[p] = i*4 + j
				p++
			}
		}
	}
	// 此时p为 地图上所有空点的总数量,且数组中存储了地图上所有空点的坐标
	tmp := rand.Intn(p)
	score := (rand.Int()%2 + 1) * 2
	t.data[tmp/4][tmp%4] = score
	t.socre += score
	fmt.Println("随机数")
}

// 将地图数据展示出来
func (t *Model) View() {
	fmt.Printf("\n\n\n\n\n\n\n\n")

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if t.data[i][j] == 0 {
				fmt.Printf("\t")
			} else {
				fmt.Printf("%d\t", t.data[i][j])
			}

		}
		fmt.Println()
	}
	fmt.Println("--------------------------------")
	fmt.Println(t.socre)
}

// 设置游戏进入控制流
func (t *Model) Controller() {

	var name string
	for !t.end {
		fmt.Println("等待输入")
		_, _ = fmt.Scan(&name)
		fmt.Println("输入", name, name == "a")
		switch name {
		case "4":
			t.Left()
		case "6":
			t.Right()
		case "8":
			t.Up()
		case "2":
			t.Down()
		default:
			fmt.Println("无法识别")
		}
		t.View()
	}

}

// 将地图数据展示出来
func (t *Model) Init() {
	// 设置游戏状态 为未结束
	t.end = false
	//设置随机数种子
	rand.Seed(time.Now().Unix())
	t.Random()
	t.Random()
	t.View()
	t.Controller()
}
