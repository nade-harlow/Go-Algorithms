package main
func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	// Write your code here
	if v1>v2{
		if ((x1-x2)%(v2-v1))==0{
			return "YES"
		}
	}

	return "NO"
}