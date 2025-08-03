def draw_hexagon(width, height):
    lines = []
    
    # Top part of hexagon
    for i in range(height):
        spaces_before = height - 1 - i
        spaces_between = width + 2 * i
        line = ' ' * spaces_before + '/' + ' ' * spaces_between + '\\'
        lines.append(line)
    
    # Middle part (top edge)
    line = '/' + '_' * width + '\\'
    lines.append(line)
    
    # Bottom part of hexagon
    for i in range(height):
        spaces_before = i
        spaces_between = width + 2 * (height - 1 - i)
        line = ' ' * spaces_before + '\\' + '_' * spaces_between + '/'
        lines.append(line)
    
    return lines

def main():
    t = int(input())
    
    all_hexagons = []
    
    for _ in range(t):
        width, height = map(int, input().split())
        hexagon_lines = draw_hexagon(width, height)
        all_hexagons.append(hexagon_lines)
    
    # Print all hexagons with one empty line between them
    for i, hexagon in enumerate(all_hexagons):
        if i > 0:
            print()  # Empty line between hexagons
        for line in hexagon:
            print(line)

if __name__ == "__main__":
    main()