import re

def read_passports():
    file = open("input.txt", mode = "r")
    lines = file.read()
    file.close()
    line = lines.split("\n\n")
    passport_counter = 0
    passport_dictionary = {}
    for passport in line:
        passport = passport.replace("\n", " ")
        passport_counter+=1
        passport = passport.split(" ")
        passport_fields = {}
        for field in passport:
            new_field = field.split(":")
            key = new_field[0]
            value = new_field[1]
            passport_fields[key] = value
        passport_dictionary[passport_counter] = passport_fields
    return passport_dictionary

def is_valid():
    passports = read_passports()
    valid_passport_count = 0
    invalid_passport_count = 0
    eye_colors = ['amb', 'blu', 'brn', 'gry', 'grn', 'hzl', 'oth']
    hair_color = re.compile("^#((\d|[a-f]){6})")
    height = re.compile('((5[9])|(6[0-9])|(7[0-6])in)|((1[5-8][0-9])|(1[9][0-3])cm)')
    expiration_year = re.compile('(202[0-9])|(2030)')
    birth_year = re.compile('(19[2-9][0-9])|(200[0-2])')
    issue_year = re.compile('(201[0-9])|(202[0])')
    for passport in passports:
        try:
            assert len(passports[passport]['pid']) == 9
            int(passports[passport]['pid'])
            assert passports[passport]['ecl'] in eye_colors
            assert re.match(hair_color, passports[passport]['hcl'])
            assert re.match(height, passports[passport]['hgt'])
            assert re.match(expiration_year, passports[passport]['eyr'])
            assert re.match(issue_year, passports[passport]['iyr'])
            assert re.match(birth_year, passports[passport]['byr'])
            valid_passport_count+=1
        except:
            invalid_passport_count+=1
            continue

    print(valid_passport_count, invalid_passport_count)


if __name__ == "__main__":
    is_valid()