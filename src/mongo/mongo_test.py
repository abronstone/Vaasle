from pymongo import MongoClient

# Returns a Collection callable from Mongo Client
def get_database():
    #client = MongoClient('localhost', port=27017, username="vaas1", password="pass")
    client = MongoClient('localhost', port=27017)

    return client['vaasdb']

# Get the database
db = get_database()

# Create a collection called 'test_collection'
test_collection = db["test_collection"]

# Clean all documents within the collection
test_collection.delete_many({})

# Add a test element to the collection
item1 = {"id":1,"element":"element one"}
item2 = {"id":2,"element":"element two"}
test_collection.insert_many([item1,item2])

# Aggregate pipeline that returns all elements
pipeline1 = []
retrieval = test_collection.aggregate(pipeline1)
print("\nALL ELEMENTS:")
for r in retrieval:
    print(r)

# Aggregate pipeline that returns element with id 1
pipeline2 = [
    {
        "$match": {
            "id": 1
        }
    }
]
retrieval = test_collection.aggregate(pipeline2)
print("\nELEMENT WITH ID 1:")
for r in retrieval:
    print(r)