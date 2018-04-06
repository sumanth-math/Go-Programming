# This program checks if the given strings are Anagrams or not
# Anagram - is a word, phrase, or name formed by rearranging the letters of another, such as cinema, formed from iceman.

# Method which takes two strings as input parameters and determines if they are anagrams or not
def isAnagram(string1, string2):
    # holds the state of whether the 
    anagramCheck = True
    # Dictionary to hold the number of occurences of each of the characters in the input strings
    charDictionary = {}
    
    # Remove any spaces in the input strings
    firstString = string1.replace(' ', '')
    secondString = string2.replace(' ', '')

    # If the length of both strings are not equal then they are not anagrams
    if len(firstString) == len(secondString):
        
        # Loop though each of the characters in the string
        for characterCount in range(len(firstString)):
            # Get character from the first string
            firstStringCharcter = firstString[characterCount]
            # Get character from the second string
            secondStringCharcter = secondString[characterCount]

            # If character in the first string is already in the dictionary and its count is > than 0 then decrease the count
            if firstStringCharcter in charDictionary and charDictionary[firstStringCharcter] > 0:
                charDictionary[firstStringCharcter] -= 1
            # If character in the first string  already in the dictionary and its count is > than 1 then decrease the count by 1
            else:
                charDictionary[firstStringCharcter] = 1
            
            # If character in the first string is already in the dictionary and its count is > than 0 then decrease the count
            if secondStringCharcter in charDictionary and charDictionary[secondStringCharcter] > 0:
                charDictionary[secondStringCharcter] -= 1
            # If character in the first string  already in the dictionary and its count is > than 1 then decrease the count by 1
            else:
                charDictionary[secondStringCharcter] = 1

        # Loop through all the chracters in the dictionary and check if any character count is not 0
        # If there is atleast one character with count greater than 0 then the string are not anagrams
        for value in charDictionary.values():
            if value != 0:
                anagramCheck = False
                break
    else:
        anagramCheck = False

    # Return back stating whether the given strings are anagrams or not
    return anagramCheck

# Get the input strings
firstString = input('Enter the first string: ')
secondString = input('Enter the second string: ')
print(firstString)
print(secondString)

# Call the method to check if the given input strings are anagrams are not
if isAnagram(firstString, secondString) == True:
    print('The two given strings are ANAGRAMS!')
else:
    print('The two given strings are NOT ANAGRAMS!')