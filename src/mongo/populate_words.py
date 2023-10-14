from pymongo import MongoClient
import re

files = {"en5.txt":5,"en6.txt":6}

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

# Create regular expression pattern to filter out explicit and/or derogatory words
cuss_word_regex = []
cuss_words = []
with open("./word_files/remove_words.txt","r") as cuss_word_file:
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
words = []
for f in files.keys():
    with open("./word_files/"+f,"r") as word_file:
        while True:
            text = word_file.readline()
            if not text:
                break
            raw_word = text[:-1]
            if not remove_pattern.search(raw_word):
                new_word = {"word":raw_word,"language":"english","length":files[f]}
                words.append(new_word)
    word_file.close()

# Use this to convert python list to lower case text word list
# with open("en6.txt","w") as en6_file:
#     for word in en6:
#         new_word = {"word":word.lower(),"language":"english","length":6}
#         words.append(new_word)
#         en6_file.write(word.lower()+"\n")

# Insert all words into the database
word_collection.insert_many(words)