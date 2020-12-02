import re

def read_input():
    file = open("input.txt", mode = "r")
    lines = file.readlines()
    file.close()
    passwords = {}
    counter = 0
    for line in lines:
        line = line.strip("\n")
        newline = re.split(' |: |-', line)
        passwords[counter] = newline
        counter+=1
    return passwords

def find_positions(password_requriements):   
        first_position = password_requriements[0]
        second_position = password_requriements[1]
        available_positions = (int(first_position), int(second_position))
        return available_positions
        
def get_delimeter(password_requriements):
    delimeter = password_requriements[2]
    return delimeter

def character_positions(character, password):
    positions = []
    for i, letter in enumerate(password):
        if letter==character:
            positions.append(i+1)
    return positions

def analyze():
    password_requriements = read_input()
    correct_password_count = 0
    incorrect_password_count = 0
    for password in password_requriements:
        password_string = password_requriements[password][3]
        first_position, second_position = find_positions(password_requriements[password])
        character = get_delimeter(password_requriements[password])
        positions = character_positions(character, password_string)
        if first_position in positions and second_position not in positions:
            correct_password_count+=1
        elif first_position not in positions and second_position in positions:
            correct_password_count+=1
        else:
            incorrect_password_count+=1    
    print(correct_password_count)
        
        
if __name__ == "__main__":
    analyze()