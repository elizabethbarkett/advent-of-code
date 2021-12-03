def read_input(modifier):
    file = open("input.txt", mode = "r")
    lines = file.readlines()
    file.close()
    forest_length = len(lines)
    modifier = forest_length*modifier
    counter = 0
    forest = {}
    for line in lines:
        counter +=1
        line = line.strip('\n')
        forest[counter] = (line*modifier)
    return(forest)

def create_tobogganing_hillside(modifier):
    forest = read_input(modifier) 
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

def get_new_position(starting_position, y_modifier, x_modifier):
    y, x = starting_position
    new_position = ((y+y_modifier), (x+x_modifier))
    return new_position

def am_i_hitting_a_tree(position, mapped_forest)-> bool:
    y, x = position
    if mapped_forest[y][x] == '#':
        return True
    else:
        return False
    
def maybe_hit_some_trees_i_guess(x_modifier, y_modifier):
    mapped_forest = create_tobogganing_hillside(x_modifier)
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
            starting_position = get_new_position(starting_position, y_modifier, x_modifier)
    print("trees hit: ",oh_crap_a_tree, "\nopen hillside: ",all_clear)
    return oh_crap_a_tree
    

if __name__ == "__main__":
    first = maybe_hit_some_trees_i_guess(1, 1)
    second = maybe_hit_some_trees_i_guess(3, 1)
    third = maybe_hit_some_trees_i_guess(5, 1)
    fourth = maybe_hit_some_trees_i_guess(7, 1)
    fifth = maybe_hit_some_trees_i_guess(1, 2)
    print(first*second*third*fourth*fifth)