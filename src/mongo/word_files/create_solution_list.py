import re

file_prefixes = ["en6"]
pattern = r".*(?<!s)$"  # Pattern to match words not ending with 's'

for prefix in file_prefixes:
    words = []

    # Read words from the file and filter out those ending with 's'
    with open(prefix + ".txt", "r") as word_file:
        for text in word_file:
            if re.search(pattern, text.strip()):
                words.append(text.strip())

    # Write the filtered words to a new file
    with open(prefix + "_solutions.txt", "w") as solution_file:
        for word in words:
            solution_file.write(word + "\n")
