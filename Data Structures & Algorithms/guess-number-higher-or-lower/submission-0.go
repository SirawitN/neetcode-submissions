/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

 func binGuess(n int, start, stop int) int {
	if start > stop {
		return -1
	}

	mid := start + (stop-start)/2
	guessResult := guess(mid)
	if guessResult==0 {
		return mid
	}

	if guessResult==-1 {
		return binGuess(n, start, mid-1)
	}
	return binGuess(n, mid+1, stop)
 }

func guessNumber(n int) int {
	return binGuess(n, 0, n)
}
