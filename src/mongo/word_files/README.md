# Word Files and Filtering

Run the word population program from this directory using the following command:

`python populate_words.py`

----------

Word file names are structured as so:

`en6.txt` --> English words with six letters
`sp5.txt` --> Spanish words with five letters

All words in the list must be separated by only a new line character, and must be lower case.

**IMPORTANT**: Before running the populate words script, run `python filter.py` and make sure you have a file called *remove_words.txt* that contains **explicit** and/or **derogatory** words that may offend users. Any words extracted from community created word lists using the `filter.py` program will be run through a regular expression filter to ensure that these words do not enter the database.