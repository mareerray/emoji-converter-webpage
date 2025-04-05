import json
from collections import Counter

# Load the JSON file
with open("../internal/emojis.json", "r") as file:  # Adjust path if needed
    emojis = json.load(file)

# Extract all names
names = [emoji["name"] for emoji in emojis]

# Count occurrences of each name
name_counts = Counter(names)

# Find duplicates (names with count > 1)
duplicates = {name: count for name, count in name_counts.items() if count > 1}

# Print duplicates
if duplicates:
    print("Duplicate names found:")
    for name, count in duplicates.items():
        print(f"{name}: {count} occurrences")
else:
    print("No duplicate names found.")
