/*
 * hi, let's test
 *
 * =======================
 *
 * date:20170525
 * author:xiaosean 
 *
 */

// variables
const Fn = 1
var n int  = 8
var FNminus1 int = 1
var temp int
func void test_func1(){

}
func void test_func2(a int){
	
}
func test_func3(){
	
}
func test_func4(a int){
	
}
// compute the nth Fibonacci number
func void main( ) {
	var a int = 5
	var b int 
	var c int = 6
	var d real 
	var d_flag bool = false

	if(a > b)
		temp = Fn
	else
		temp = Fn

	if(a > b)
		temp = Fn
	if(a > b)
		temp = Fn
	else
		temp = Fn
	for (n == 1) 
		temp = Fn
	for (n == 1) {
		temp = Fn
	}
	for (n = 2; n == 1) 
		temp = Fn
	for (n = 2; n == 1) {
		temp = Fn
	}
	//next part will type match error
	// for (n = 2; n = 1) {
	// 	temp = Fn
	// }

	for (n == 2; n = 1) {
		temp = Fn
	}
	for (n == 2; n = 1) 
		temp = Fn
	for (n = 1; n > 2; n = n - 1) {
		{}
		temp = Fn
		temp = Fn + FNminus1
		FNminus1 = temp
		for (n = 1; n > 2; n = n - 1) {
			{}
			temp = Fn
			temp = Fn + FNminus1
			FNminus1 = temp
			
		}		
	}
	for (n > 2; n = n - 1) {
		temp = Fn
		temp = Fn + FNminus1
		FNminus1 = temp
	}
	if(a > 5)
		temp = Fn
	if(!true)
		temp = Fn
	if(!(a == b))
		a = b
	if(d_flag | !(false))
		a = b
	if(d_flag | !(d_flag))
		a = b
	if(!(a > b) | !(a == b))
		a = b
	if(!(a > b) | (c > b) & (a > b) | !(a == b))
		temp = Fn
	d_flag = a > b
	/* print result */
	print "Result of computation: "
	println n
	print 7 * 10
	println 2 - 7 * 10
	read a
	go test_func1()
	go test_func2(5)
	go test_func3()
	go test_func4(a)
	
	//next will type error
	// a = test_func2(5)
	
	
	return
	return a

	//next line will be abort  bcz a is integer, d is real
	// if(!(a > b) | (c > b) & (a > d) | !(a == b))
	// 	temp = Fn

}
