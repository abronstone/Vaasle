from pymongo import MongoClient
import re

user_confirmation = input("Have you run filter.py on all text files to be inserted? (y/n)")

if user_confirmation != "y":
    print("Please run the filter first")
    exit()

file_wordlengths = {"en5.txt":5,"en6.txt":6, "sp5.txt":5}
file_languages = {"en5.txt":"english","en6.txt":"english","sp5.txt":"spanish"}


# Returns a Collection callable from Mongo Client
def get_database():
    uri = "mongodb+srv://vaas_admin:adv1software2design3@vaasdatabase.sarpr4r.mongodb.net"
    client = MongoClient(uri)
    return client["VaasDatabase"]

# Get the database
db = get_database()

# Create a collection called 'test_collection'
word_collection = db["words"]

# Clean all documents within the collection
word_collection.delete_many({})

words = []

'''
    NOTE: Use the code below only if the "clean_words
'''

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
for f in file_wordlengths.keys():
    with open(f,"r") as word_file:
        while True:
            text = word_file.readline()
            if not text:
                break
            raw_word = text[:-1]
            if not remove_pattern.search(raw_word):
                new_word = {"word":raw_word,"language":file_languages[f],"length":file_wordlengths[f]}
                words.append(new_word)
            else:
                print("CUSS WORD FOUND, please run the filter.py before you run this code!")
                exit()
    word_file.close()


# Use this to convert python list to lower case text word list
# with open("en6.txt","w") as en6_file:
#     for word in en6:
#         new_word = {"word":word.lower(),"language":"english","length":6}
#         words.append(new_word)
#         en6_file.write(word.lower()+"\n")

# Insert all words into the database
word_collection.insert_many(words)