# Development choice
# I have choose to avoid the hidden files and directories
#!/usr/bin/python3
import json
import os
import re
import math

# count the alphanum characters in the string received as input
import sys


def count_alphanum(data):
    counter = 0
    for ch in data:
        if ch.isalnum():
            counter += 1
    return counter

# I have used this definition of word https://en.wikipedia.org/wiki/Word


def length_words(word):
    v_response = []
    for word in re.findall("[a-zA-Z]+", word):
        v_response.append(len(word))
    return v_response


def calculate_standard_deviation(array, avg, n):
    sol = 0.0
    for item in array:
        sol += (item - avg)**2
    return math.sqrt(sol/n)

# the core of the implantation is in this folder


def get_stats_folder(root_dir):
    array_stats = []
    for dir_name, subdir_list, file_list in os.walk(root_dir):
        # check if the directory is hidden
        if dir_name[0] == '.':
            continue

        # initializations for each folder read
        v_count_alphanum, v_length_word = [], []
        sum_bytes = 0
        # calculations for each file

        for file in file_list:
            # check if the directory is hidden
            if file[0] == '.':
                continue

            # open and read from each file
            dir_name += "/" if dir_name[-1]!='/'else ""
            f = open(dir_name+file, encoding="latin-1", mode="r")
            payload = f.read()
            v_count_alphanum.append(count_alphanum(payload))
            v_length_word += length_words(payload)
            sum_bytes += len(payload)

        # refining the results got from the read
        n_files = len(file_list)
        l_v_count_alphanum = len(v_count_alphanum)
        l_v_count_length_word = len(v_length_word)

        if l_v_count_alphanum > 0:
            avg_alphanum = (sum(v_count_alphanum)/l_v_count_alphanum)
            std_dev_alphanum = calculate_standard_deviation(v_count_alphanum, avg_alphanum, n_files)
        else:
            avg_alphanum = std_dev_alphanum = 0
        if l_v_count_length_word>0:
            avg_length_word = (sum(v_length_word)/l_v_count_length_word)
            std_dev_legth_word = calculate_standard_deviation(v_length_word, avg_length_word, n_files)
        else:
            avg_length_word = std_dev_legth_word = 0

        # preparing the Json record and appending it
        tmpdict = {dir_name: {}}
        tmpdict[dir_name]["nFiles"] = n_files
        tmpdict[dir_name]["avgAlphanum"] = avg_alphanum
        tmpdict[dir_name]["stdDevAlphanum"] = std_dev_alphanum
        tmpdict[dir_name]["avgLengthWord"] = avg_length_word
        tmpdict[dir_name]["stdDevLengthWord"] = std_dev_legth_word
        tmpdict[dir_name]["sumBites"] = sum_bytes
        array_stats.append(tmpdict)

    return json.dumps(array_stats)


if __name__ == "__main__":
    try:
        root_directory = sys.argv[1]
    except IndexError:
        root_directory = '.'
    print(get_stats_folder(root_directory))
