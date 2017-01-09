# In The Name Of God
# ========================================
# [] File Name : level-generator.py
#
# [] Creation Date : 04-01-2017
#
# [] Created By : Parham Alvani (parham.alvani@gmail.com)
# =======================================
import requests
import re


if __name__ == '__main__':
    levels = 10
    res = requests.get('http://loripsum.net/api/%d/short/plaintext' % levels)
    index = 0
    for line in res.text.splitlines():
        if line != '':
            index += 1
            f = open('level-%d.txt' % index, 'w')
            output = ''
            for word in line.split(' '):
                if re.fullmatch(r'[a-z]+', word):
                    output += word + ' '
            f.write(output)
