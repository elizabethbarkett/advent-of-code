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
    for passport in passports:
        try:
            assert passports[passport]['pid']
            assert passports[passport]['ecl']
            assert passports[passport]['hcl']
            assert passports[passport]['hgt']
            assert passports[passport]['eyr']
            assert passports[passport]['iyr']
            assert passports[passport]['byr']
            #print(passports[passport])
            valid_passport_count+=1
        except KeyError:
            print(passport, passports[passport])
            invalid_passport_count+=1
            continue
    print(valid_passport_count, invalid_passport_count)


if __name__ == "__main__":
    is_valid()