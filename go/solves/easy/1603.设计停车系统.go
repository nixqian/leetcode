package easy

/*
 * @lc app=leetcode.cn id=1603 lang=golang
 *
 * [1603] 设计停车系统
 */

// @lc code=start
type ParkingSystem struct {
	cnt []int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	cnt := []int{0, big, medium, small}
	return ParkingSystem{cnt: cnt}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.cnt[carType] <= 0 {
		return false
	}
	this.cnt[carType]--
	return true
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
// @lc code=end
