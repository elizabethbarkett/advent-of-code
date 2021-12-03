def read_input():
    file = open("input.txt", mode = "r")
    lines = file.readlines()
    file.close()
    forest_length = len(lines)
    modifier = forest_length*3
    counter = 0
    forest = {}
    for line in lines:
        counter +=1
        line = line.strip('\n')
        forest[counter] = (line*modifier)
    return(forest)

def create_tobogganing_hillside():
    forest = read_input() 
    toboggan_map = {}
    counter = 0
    mapped_forest = {}
    for hillside in forest:
        counter+=1
        hillside_section = []
        for j, geographic_status in enumerate(forest[hillside]):
            hillside_section.append(geographic_status)
        mapped_forest[counter] = hillside_section   
    return mapped_forest

def get_new_position(starting_position):
    y, x = starting_position
    new_position = (y+1, x+3)
    return new_position

def am_i_hitting_a_tree(position, mapped_forest)-> bool:
    y, x = position
    if mapped_forest[y][x] == '#':
        return True
    else:
        return False
    
def maybe_hit_some_trees_i_guess():
    mapped_forest = create_tobogganing_hillside()
    starting_position = (1, 0)
    forest_length = len(mapped_forest)
    oh_crap_a_tree = 0
    all_clear = 0
    for hillside_section in mapped_forest:
        while starting_position[0] <= forest_length:
            very_important_question = am_i_hitting_a_tree(starting_position, mapped_forest)
            if very_important_question:
                oh_crap_a_tree+=1
            else:
                all_clear+=1
            starting_position = get_new_position(starting_position)
    print("trees hit: ",oh_crap_a_tree, "\nopen hillside: ",all_clear)
    

if __name__ == "__main__":
    maybe_hit_some_trees_i_guess()