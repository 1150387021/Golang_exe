package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

//二分查找

func bin_searchstruct(arr []uidx, data int) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := (left + right) / 2
		if data > arr[mid].id {
			left = mid + 1
		} else if data < arr[mid].id {
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

//快速排序，内存占用较大

func QuickSortstruct(arr []uidx) []uidx {
	length := len(arr)
	if length <= 1 {
		return arr //一个元素的数组，直接返回
	} else {
		splitdata := arr[0] //以第一个为基准
		low := make([]uidx, 0, 0)
		high := make([]uidx, 0, 0)
		mid := make([]uidx, 0, 0)
		mid = append(mid, splitdata) //加入第一个相等的
		for i := 1; i < length; i++ {
			if arr[i].id < splitdata.id {
				low = append(low, arr[i])
			} else if arr[i].id > splitdata.id {
				high = append(high, arr[i])
			} else {
				mid = append(mid, arr[i])
			}
		}
		low, high = QuickSortstruct(low), QuickSortstruct(high) //切割递归处理
		myarr := append(append(low, mid...), high...)
		return myarr
	}
}

//进阶版快速排序

func SortForMergeu(arr []uidx, left int, right int) {
	for i := left; i <= right; i++ {
		temp := arr[i].id //数据备份
		var j int
		for j = i; j > left && arr[j-1].id > temp; j-- { //定位
			arr[j].id = arr[j-1].id //数据往后移动
		}
		arr[j].id = temp //插入
	}
}

//数据交换

func swapu(arr []uidx, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

//递归快速排序

func QuickSortXu(arr []uidx, left int, right int) {
	if right-left < 3 { //数组剩下3个数，直接插入排序（超大数据默认15个）
		SortForMergeu(arr, left, right)

	} else {
		//随机找一个数字，放在第一个位置
		swapu(arr, left, rand.Int()%(right-left+1)+left)
		vdata := arr[left].id //坐标数组，比我小，左边，比我大右边
		lt := left            //arr[left+1, lt]<vata
		gt := right + 1       //arr[gt ... r]>vata
		i := left + 1         //arr[lt+1,...i] == vdata
		for i < gt {
			if arr[i].id < vdata {
				swapu(arr, i, lt+1) //移动到小于到地方
				lt++                //前进循环
				i++
			} else if arr[i].id > vdata {
				swapu(arr, i, gt-1) //移动到大于到地方
				gt--
			} else {
				i++
			}
		}
		swapu(arr, left, lt)         //交换头部位置
		QuickSortXu(arr, left, lt-1) //递归处理小于那一段
		QuickSortXu(arr, gt, right)  //递归处理大于那一段

	}
}

func QuicksortPlusu(arr []uidx) {
	QuickSortXu(arr, 0, len(arr)-1)
}

type uidx struct {
	list   int
	id     int
	like   int
	school string
	is_985 string
}

const N1 = 637

func main() {
	allstrs := make([]uidx, N1, N1) //初始化数组

	//文件读取
	path := "/Users/st/space/数据/data.csv"
	fi, err := os.Open(path)
	if err != nil {
		fmt.Println("文件读取失败", err)
		return
	}
	defer fi.Close() //延迟关闭文件
	i := 0
	br := bufio.NewReader(fi)
	for {
		line, _, end := br.ReadLine()
		if end == io.EOF {
			break
		}
		linestr := string(line)                     //读取，转化为字符串
		lines := strings.Split(linestr, ",")        //字符串切割
		allstrs[i].list, _ = strconv.Atoi(lines[0]) //字符串转整数
		allstrs[i].id, _ = strconv.Atoi(lines[1])
		allstrs[i].like, _ = strconv.Atoi(lines[2])
		allstrs[i].school = lines[3]
		allstrs[i].is_985 = lines[4]

		i++
	}
	//fmt.Println(i)
	fmt.Println("数据载入内存")
	//time.Sleep(time.Second * 1)

	//开始排序
	startsorttime := time.Now()
	fmt.Println("开始排序")
	allstrs = QuickSortstruct(allstrs) //使用快速排序
	//QuicksortPlusu(allstrs) //加强版快速排序
	fmt.Println("结束排序一共用了", time.Since(startsorttime))
	//fmt.Println(allstrs)

	for {
		fmt.Println("请输入要查询的用户ID：")
		var userid int
		fmt.Scanf("%d", &userid)

		starttime := time.Now()
		index := bin_searchstruct(allstrs, userid)

		if index == -1 {
			fmt.Println("数据查找不到")
		} else {
			fmt.Println("数据找到", index, allstrs[index])
		}
		fmt.Println("一共用了", time.Since(starttime))

	}
}

/*
54	19489141	0	华中科技大学同济医学院	TRUE
55	19498675	1	信阳师范学院	FALSE
56	19498774	1	吉林建筑大学	FALSE
57	19501262	0	湖南工艺美术职业学院	FALSE
58	19521230	0	哈尔滨师范学院	FALSE
59	19525934	1	西北师范大学	FALSE
60	19529213	1	广西大学行健文理学院	FALSE
61	19557614	1	弗吉尼亚大学	FALSE
62	19560802	1	新疆财经大学	FALSE
63	19573166	1	辽宁师范大学	FALSE
64	19574818	1	燕京理工学院	FALSE
65	19581637	0	北京交通大学	FALSE
66	19585262	1	北京第二外国语学院	FALSE

129	19655898	1	西安铁道职业学校	FALSE
130	19657508	1	黑龙江大学	FALSE
131	19657574	0	内蒙古大学	FALSE
132	19658942	0	重庆外语外事学院	FALSE
133	19660551	0	南开大学	TRUE
134	19660784	1	中国政法大学	FALSE
135	19661849	0	佳木斯大学	FALSE
136	19663304	1	临沂大学	FALSE
137	19663571	1	印度尼西亚大学	FALSE
138	19665489	1	河南财经政法大学	FALSE
139	19669162	1	北京师范大学	TRUE
*/
