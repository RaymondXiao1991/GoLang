package main

// 计算总账单的账期,比如租12个月 ，押一付三，总账期就是4，如果租13个月，押一付三，总账期就是5，最后一个月单独处理
func CalcTotalBillTerm(payMonth int, lease int) (term int, month int) {
	term = lease / payMonth
	month = lease % payMonth
	if month > 0 {
		term += 1
	}
	return term, month
}
