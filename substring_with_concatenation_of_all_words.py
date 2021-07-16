def find_substring(s, words):
    indices = []
    n = len(words[0])
    for i in range(len(s)):
        unused = words[:]
        j = i
        while s[j:j+n] in unused:
            unused.remove(s[j:j+n])
            j += n
            if unused == []:
                indices.append(i)
    return indices

print(find_substring("barfoothefoobarman", ["foo", "bar"]))
print(find_substring("wordgoodgoodgoodbestword", ["word","good","best","word"]))
print(find_substring("barfoofoobarthefoobarman", ["bar","foo","the"]))
