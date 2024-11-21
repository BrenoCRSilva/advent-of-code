import sys

def main():
    line = open(sys.argv[1]).read().strip().split("\n")
    add_dotted_line(line)
    matrix = [list(x) for x in line]
    add_dotted_column(matrix)
    row = len(matrix)
    column = len(matrix[0])
 
def add_dotted_line(line):
    dot_line = "." * len(line)
    line.insert(0, dot_line)
    line.insert(len(line), dot_line) 

def add_dotted_column(matrix):
    for i in range(len(matrix)):
        matrix[i].insert(0, ".")
        matrix[i].insert(len(matrix[i]), ".")
    
class Box:
    def __init__(self, x, y, matrix):
        self.matrix = matrix
        self.center = matrix[x][y]
        self.south = matrix[x][y+1]
        self.north = matrix[x][y-1] 
        self.east = matrix[x+1][y]
        self.west = matrix[x-1][y]
        self.neast = matrix[x+1][y+1]
        self.nwest = matrix[x-1][y+1]
        self.seast = matrix[x+1][y-1]
        self.swest = matrix[x-1][y-1]
        self.coordinates = [self.center, self.south, self.north, self.east, self.west, self.neast, self.nwest, self.seast, self.swest]
        self.top = [self.nwest, self.north, self.neast]
        self.bottom = [self.swest, self.south, self.seast]
        
def p1(matrix, row, column):
    boxes = []
    number_index = []
    for i in range(1, row):
        for j in range(1, column):
            box = Box(i, j, matrix)
            if not box.center.isalpha():
                boxes.append(box)
    for box in boxes:
        for coordinate in box.coordinates:
            if coordinate.isdigit():
                number_index.append(coordinate)
    for i,j in number_index:
        num = ''
        id = ''
        a = 0
        parts = {}
        while a <= 2 or not matrix[i][j].isdigit():
            num += matrix[i+a][j]
            id += i+a
            a += 1
        while a <= 2 or not matrix[i][j].isdigit():
            num += matrix[i-a][j]
            id += i-a
            a += 1
        if int(id) not in parts.keys():
            parts[int(id)] = int(num)
main()

