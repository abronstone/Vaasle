# Word Files and Filtering

Run the word population program from this directory using the following command:

`python populate_words.py`

----------

Word file names are structured as so:

`en6.txt` --> English words with six letters
`sp5.txt` --> Spanish words with five letters

All words in the list must be separated by only a new line character, and must be lower case.

**IMPORTANT**: The file named *remove_words.txt* contains **explicit** and/or **derogatory** words that may offend users. Any words extracted from community created word lists using the `populate_words.py` program will be run through a regular expression filter to ensure that these words do not enter the database.