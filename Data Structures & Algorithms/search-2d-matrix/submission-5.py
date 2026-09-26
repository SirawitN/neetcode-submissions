import math

class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        numsRow, numsCol = len(matrix), len(matrix[0])

        rowLeft, rowRight = 0, numsRow-1
        while rowLeft <= rowRight:
            targetRow = math.ceil((rowLeft+rowRight)/2)

            if matrix[targetRow][0] <= target and matrix[targetRow][numsCol-1] >= target:
                break
            

            if matrix[targetRow][0] > target:
                rowRight = targetRow-1
            else:
                rowLeft = targetRow+1
            
        
        colLeft, colRight = 0, numsCol-1
        while colLeft <= colRight:
            targetCol = math.ceil((colLeft+colRight)/2)
            if matrix[targetRow][targetCol]==target:
                return True

            if matrix[targetRow][targetCol] > target:
                colRight = targetCol-1
            else:
                colLeft = targetCol+1

        return False