from pymongo import MongoClient
import re

user_confirmation = input("This program requires the file \"remove_words.txt\" in the same directory, do you have that? (y/n)")

if user_confirmation != "y":
    print("Please include a \'remove_words.txt\" file")
    exit()

files = ['en5.txt','en6.txt','sp5.txt']

words = []

# Create regular expression pattern to filter out explicit and/or derogatory words
cuss_word_regex = []
cuss_words = []
with open("remove_words.txt","r") as cuss_word_file:
    while True:
        word = cuss_word_file.readline()
        if not word:
            break
        word = word[:-1].strip()
        cuss_words.append(word)
        cur_regex={'word': {'$regex': word}}
        cuss_word_regex.append(cur_regex)
    cuss_word_file.close()
remove_pattern = re.compile('|'.join(re.escape(phrase) for phrase in cuss_words))

# Add all words to a list
for f in files:
    with open(f,"r") as word_file:
        while True:
            text = word_file.readline()
            if not text:
                break
            raw_word = text[:-1]
            if not remove_pattern.search(raw_word):
                words.append(raw_word)
        word_file.close()
    with open(f,"w") as word_file:
        for word in words:
            word_file.write(str(word)+"\n")
        words.clear()
        word_file.close()



# Use this to convert python list to lower case text word list
# with open("en6.txt","w") as en6_file:
#     for word in en6:
#         new_word = {"word":word.lower(),"language":"english","length":6}
#         words.append(new_word)
#         en6_file.write(word.lower()+"\n")

# Insert all words into the database