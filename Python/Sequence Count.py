# This program finds the count of each sequence in a given list
# For example, if the given list is [1,1,1,1,4,4,1,1,1,1,1,1,5,6,2,2,2,2] then the output should be
# Count of Sequence Number, Sequence Number, i.e Four 1's, Two 4's, Six 1's, one 5, one 6, four 2's [4,1, 2,4, 6,1, 1,5, 1,6, 4,2]

# Method which returns the sequence count list for a given input list
def runLengthEncode(unCompressedDataList):
    # Sequence Count List
    sequenceCountList = []
    # Holds the current sequence number
    currentSequenceNumber = None
    # Holds the previous sequence number - first item from the input list
    previousSequenceNumber = unCompressedDataList[0]
    # Holds the count of each number sequence
    sequenceNumberCount = 0

    # Loop though the input list
    for uncompressedDataItem in unCompressedDataList:
        # set the indexed list item as current
        currentSequenceNumber = uncompressedDataItem
        # Current sequence number not the same as previous then 
        # Add the count and sequence number to the Sequence Count List and
        # reset sequence number count to 1
        if currentSequenceNumber != previousSequenceNumber:
            sequenceCountList.append(sequenceNumberCount)
            sequenceCountList.append(previousSequenceNumber)
            sequenceNumberCount = 1
            previousSequenceNumber = currentSequenceNumber
        # Current sequence number same as previous then increment sequence number count 
        else:
            sequenceNumberCount += 1

    # Add the last list item count and sequence number to the Sequence Count List
    sequenceCountList.append(sequenceNumberCount)
    sequenceCountList.append(previousSequenceNumber)

    # Return the sequence count list
    return sequenceCountList

# input List
unCompressedDataList = [1,1,1,1,4,4,1,1,1,1,1,1,5,6,2,2,2,2]

print("Uncompressed Data List")
for uncompressedDataItem in unCompressedDataList:
         print(uncompressedDataItem)

# Method call to get the sequence count
sequenceCountList = runLengthEncode(unCompressedDataList)

# Print the Sequence count list
print("Compressed Data List")
for sequenceCountItem in sequenceCountList:
         print(sequenceCountItem)