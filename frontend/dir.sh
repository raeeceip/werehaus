#!/bin/bash

# Define the output file
output_file="project_structure_and_contents.txt"

# Write the directory tree to the output file
echo "Directory Structure:" > $output_file
tree -L 3 >> $output_file

# Add a separator
echo -e "\n----------------------------------------\n" >> $output_file

# Find and output the contents of the warehouse management file
echo "Contents of warehouse management file:" >> $output_file
find . -name "*warehouse-management*" -type f -exec cat {} \; >> $output_file

echo "Output has been written to $output_file"