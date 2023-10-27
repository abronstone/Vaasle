import re

file_prefixes = ["en6"]

pattern = r".*(?<!s)$"

for prefix in file_prefixes:
    words = []
    with open(prefix+".txt","r") as word_file:
        while True:
            text = word_file.readline()
            if not text:
                break
            raw_word = text[:-1]
            if re.search(pattern, raw_word):
                words.append(raw_word)
        word_file.close()
    with open(prefix+"_solutions.txt","w") as word_file:
        for word in words:
            word_file.write(word+"\n")
        word_file.close()