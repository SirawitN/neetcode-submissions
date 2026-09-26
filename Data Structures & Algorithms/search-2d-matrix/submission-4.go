func searchMatrix(matrix [][]int, target int) bool {
    numRows := len(matrix)
    numCols := len(matrix[0])
    rowLeft, rowRight := 0, numRows-1
    var targetRow, targetCol int

    for rowLeft <= rowRight {
        targetRow = (rowLeft + rowRight) / 2
        if matrix[targetRow][0] <= target && matrix[targetRow][numCols-1] >= target {
            break
        }

        if matrix[targetRow][0] > target {
            rowRight = targetRow-1
        } else {
            rowLeft = targetRow+1
        }
    }

    colLeft, colRight := 0, numCols-1
    for colLeft <= colRight {
        targetCol = (colLeft+colRight) / 2
        if matrix[targetRow][targetCol]==target {
            return true
        }

        if matrix[targetRow][targetCol]>target {
            colRight = targetCol-1
        } else {
            colLeft = targetCol+1
        }
    }

    return false
}
