import csv

def get_data():
    input_list = []
    file = "data.csv"
    with open(file, "r") as csv_file:
        csv_reader = csv.reader(csv_file, delimiter = ',')
        for lines in csv_reader:
            input_list.append(int(lines[0]))
    return(input_list)

def find_three_addends(target):
    data = get_data()
    data.sort()
    length = len(data)
    for first in range (0, length-2):
        second = first + 1
        third = length-1
        while (second < third):
            if (data[first]+data[second]+data[third] == target):
                answer = [data[first], data[second], data[third]]
                print(answer)
                return answer
            elif (data[first] + data[second] + data[third] < target):
                second = second+1
            else:
                third = third - 1

def multiply_them_there_numbers_and_youll_get_an_answerydooo():
    addends = find_three_addends(2020)
    product = addends[0] * addends[1] * addends[2]
    print(product)

if __name__ == "__main__":
    multiply_them_there_numbers_and_youll_get_an_answerydooo()