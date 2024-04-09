#!/bin/bash

# Initialize first outside the loop to manage comma placement
first=1

# Start of JSON array
echo "["  

while read -r line; do
    # Check if the line contains ping data
    if echo "$line" | grep -q "bytes from"; then
        if [ $first -eq 1 ]; then
            first=0
        else
            echo ","  # Print a comma before each JSON object except the first
        fi
        # Extract relevant information using grep and output directly in JSON format
        ip_address=$(echo "$line" | grep -oP '(?<=from ).*(?=: icmp_seq)')
        icmp_seq=$(echo "$line" | grep -oP '(?<=icmp_seq=)[0-9]+')
        ttl=$(echo "$line" | grep -oP '(?<=ttl=)[0-9]+')
        time=$(echo "$line" | grep -oP '(?<=time=)[0-9.]+')
        
        # Print the extracted information in JSON format to stdout
        echo -n "{\"IP Address\": \"$ip_address\", \"ICMP Sequence\": \"$icmp_seq\", \"TTL\": \"$ttl\", \"Time\": \"$time\"}"
    fi
done

# End of JSON array
echo -e "\n]"  
