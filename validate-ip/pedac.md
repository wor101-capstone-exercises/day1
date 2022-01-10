***PROBLEM***
Given a string queryIP, return "IPv4" if IP is a valid IPv4 address, "IPv6" if IP is a valid IPv6 address or "Neither" if IP is not a correct IP of any type.

A valid IPv4 address is an IP in the form "x1.x2.x3.x4" where 0 <= xi <= 255 and xi cannot contain leading zeros. For example, "192.168.1.1" and "192.168.1.0" are valid IPv4 addresses but "192.168.01.1", while "192.168.1.00" and "192.168@1.1" are invalid IPv4 addresses.

A valid IPv6 address is an IP in the form "x1:x2:x3:x4:x5:x6:x7:x8" where:

1 <= xi.length <= 4
xi is a hexadecimal string which may contain digits, lower-case English letter ('a' to 'f') and upper-case English letters ('A' to 'F').
Leading zeros are allowed in xi.
For example, "2001:0db8:85a3:0000:0000:8a2e:0370:7334" and "2001:db8:85a3:0:0:8A2E:0370:7334" are valid IPv6 addresses, while "2001:0db8:85a3::8A2E:037j:7334" and "02001:0db8:85a3:0000:0000:8a2e:0370:7334" are invalid IPv6 addresses.

RULES:
- Must return "IPv4", "IPv6", or "Neither"
- IPV4
  - must have 3 '.'
  - each 'number' must be between 0 and 255 (inclusive)
  - can't have leading 0's
- IPV6
  - must have 7 ':'
  - each 'number' is a hexadecimal string
    - can contain letters 'a-f' and 'A-F'
    - must have length of 1 - 4

***EXAMPLES***
Example 1:

Input: queryIP = "172.16.254.1"
Output: "IPv4"
Explanation: This is a valid IPv4 address, return "IPv4".
Example 2:

Input: queryIP = "2001:0db8:85a3:0:0:8A2E:0370:7334"
Output: "IPv6"
Explanation: This is a valid IPv6 address, return "IPv6".
Example 3:

Input: queryIP = "256.256.256.256"
Output: "Neither"
Explanation: This is neither a IPv4 address nor a IPv6 address.

***DATA***

***ALGORITHM***
- check if IPv4 (return boolean)
  - split string by '.'
  - if length is NOT 4 then not IPv4 (return false)
  - loop over slice
    - confirm if each element is IPv4 subaddress (return boolean)
      - if string length is greater than 1 AND starts with 0 return false
      - convert string to integer
      - if integer is between 0 and 255 (inclusive) return true
    - if any element returns false, then return false
    - else, return true

- if not IPv4, check if IPv6
  - split string by ':'
  - if length is NOT 8 then not IPv6 (return false)
  - loop over slice
    - confirm if each element is IPv6 subaddress (return boolean)
      - if length of string is NOT 1 -4 then return false
      - use regex to confirm is string contains only digits, a-f, A-F
    - if false ever returned by subaddress, then return false
- return "neither"