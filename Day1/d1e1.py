import csv

def get_data():
    input_list = []
    file = "data.csv"
    with open(file, "r") as csv_file:
        csv_reader = csv.reader(csv_file, delimiter = ',')
        for lines in csv_reader:
            input_list.append(int(lines[0]))
    return(input_list)

def get_addends():
    target = 2020
    input_data = get_data()
    print(input_data)
    for number in input_data:
        #number = int(number)
        print(number)
        missing_addend = target - number
        if missing_addend in input_data:
            print(f"found! {number} + {missing_addend} equals 2020")
            answer = (number, missing_addend)
            return answer

def multiply_addends():
    number, addend = get_addends()
    print(number * addend)
    

if __name__ == "__main__":
    multiply_addends()