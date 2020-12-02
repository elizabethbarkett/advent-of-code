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

def get_range(password_requriements):   
        lower_bound = password_requriements[0]
        upper_bound = password_requriements[1]
        delimeter_range = (int(lower_bound), int(upper_bound))
        return delimeter_range
        
def get_delimeter(password_requriements):
    delimeter = password_requriements[2]
    return delimeter

def count_character(character, password):
    count = 0
    for i in password:
        if i==character:
            count+=1
    return count

def analyze():
    password_requriements = read_input()
    correct_password_count = 0
    for password in password_requriements:
        password_string = password_requriements[password][3]
        lower_bound, upper_bound = get_range(password_requriements[password])
        character = get_delimeter(password_requriements[password])
        count = count_character(character, password_string)
        if lower_bound<= count <= upper_bound:
            correct_password_count+=1
    print(correct_password_count)
        



if __name__ == "__main__":
    analyze()